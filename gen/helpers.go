package gen

import (
	"fmt"
	"go/ast"
	"io"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
)

func extractImports(f *ast.File) []*Import {
	var res []*Import
	for _, impt := range f.Imports {
		name := ""
		if impt.Name != nil {
			name = impt.Name.Name
		}

		res = append(res, &Import{
			Path: strings.Trim(impt.Path.Value, `"`),
			Name: name,
		})
	}

	return res
}

func extractMethodsFromInterfaces(i *ast.InterfaceType) []*Method {
	var methods []*Method
	for _, method := range i.Methods.List {
		switch fnType := method.Type.(type) {
		case *ast.FuncType:
			params, results := extractParamsAndResults(fnType)
			fn := &Method{
				Name:    method.Names[0].Name,
				Params:  params,
				Results: results,
			}

			methods = append(methods, fn)
		}
	}

	return methods
}

func extractParamsAndResults(fnDesl *ast.FuncType) ([]*Field, []*Field) {
	params := extractFieldsFromAst(fnDesl.Params.List)
	results := extractFieldsFromAst(fnDesl.Results.List)

	return params, results
}

func extractFieldsFromAst(items []*ast.Field) []*Field {
	output := []*Field{}

	for _, item := range items {
		typeStr := getTypeString(item.Type)
		name := ""

		//  nil if anonymous field
		if len(item.Names) > 0 {
			name = item.Names[0].Name
		}

		funcField := &Field{
			Type: typeStr,
			Name: name,
		}

		output = append(output, funcField)
	}

	return output
}

func getTypeString(expr ast.Expr) string {
	var result string

	switch etype := expr.(type) {
	case *ast.ArrayType:
		result = fmt.Sprintf("[]%s", getTypeString(etype.Elt))
	case *ast.MapType:
		result = fmt.Sprintf("map[%s]%s", etype.Key, etype.Value)

	case *ast.SelectorExpr:
		result = fmt.Sprintf("%s.%s", etype.X, etype.Sel)

	case *ast.StarExpr:
		result = fmt.Sprintf("*%s", getTypeString(etype.X))

	default:
		result = fmt.Sprintf("%s", etype)
	}
	return result
}

func parsePath(path string) ([]*packages.Package, error) {
	return packages.Load(
		&packages.Config{
			Mode: packages.NeedName |
				packages.NeedFiles |
				packages.NeedImports |
				packages.NeedDeps |
				packages.NeedCompiledGoFiles |
				packages.NeedTypes |
				packages.NeedSyntax |
				packages.NeedTypesInfo,
		},
		path,
	)
}

func emptyValue(typeName string) string {
	if isPointer(typeName) || isArray(typeName) {
		return "nil"
	}

	if isNumber(typeName) {
		return "0"
	}

	if typeName == "string" {
		return `""`
	}

	return typeName + "{}"
}

func isPointer(typeName string) bool {
	return strings.HasPrefix(typeName, "*")
}

func isArray(typeName string) bool {
	return strings.HasPrefix(typeName, "[]")
}

func isNumber(typeName string) bool {
	return typeName == "int" ||
		typeName == "int32" ||
		typeName == "int64" ||
		typeName == "uint32" ||
		typeName == "uint64" ||
		typeName == "float64" ||
		typeName == "float32"
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"emptyValue": emptyValue,
		"isPointer":  isPointer,
		"isNumber":   isNumber,
	}
}

func executeTemplate(tmplContext string, writer io.Writer, data interface{}) error {
	codeTmpl, err := template.New("default").
		Funcs(getFuncMap()).
		Parse(tmplContext)

	if err != nil {
		return err
	}

	return codeTmpl.Execute(writer, data)
}
