package gen

import (
	"errors"
	"go/ast"

	"golang.org/x/tools/go/packages"
)

// Data includes data for templates.
type Data struct {
	TypeName    string
	Package     string
	PackageName string
	Methods     []*Method
	Imports     []*Import
}

// Method presents a method.
type Method struct {
	Name    string
	Params  []*Field
	Results []*Field
}

// Field presents a field.
type Field struct {
	Name string
	Type string
}

type Import struct {
	Path string
	Name string
}

type DataParser interface {
	Parse([]*packages.Package) (*Data, error)
}

type interfaceParser struct {
	Name string
}

func (p *interfaceParser) Parse(pkgs []*packages.Package) (*Data, error) {
	for _, pkg := range pkgs {
		for _, f := range pkg.Syntax {
			for _, decl := range f.Decls {
				if decl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range decl.Specs {
						spec, ok := spec.(*ast.TypeSpec)
						if !ok {
							continue
						}

						sType, ok := spec.Type.(*ast.InterfaceType)
						if !ok {
							continue
						}

						if spec.Name.Name != p.Name {
							continue
						}

						d := &Data{
							TypeName:    p.Name,
							Package:     pkg.PkgPath,
							PackageName: f.Name.Name,
							Imports:     extractImports(f),
							Methods:     extractMethodsFromInterfaces(sType),
						}

						return d, nil
					}
				}
			}
		}
	}

	return nil, errors.New("no interface found")
}
