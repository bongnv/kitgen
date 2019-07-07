package gen

import (
	"io"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/bongnv/kitgen/templates"
)

type Generator interface {
	Generate(writer io.Writer, data *Data) error
}

type fileTemplate struct {
	filePath string
}

// Generate implements generator
func (t *fileTemplate) Generate(writer io.Writer, data *Data) error {
	templateContent, err := ioutil.ReadFile(t.filePath)
	if err != nil {
		return err
	}

	codeTmpl, err := template.New("code").
		Funcs(template.FuncMap{
			"emptyValue": emptyValue,
			"isPointer":  isPointer,
		}).
		Parse(string(templateContent))
	if err != nil {
		return err
	}

	return codeTmpl.Execute(writer, data)
}

type inMemoryTemplate struct {
	Name string
}

// Generate implements generator
func (t *inMemoryTemplate) Generate(writer io.Writer, data *Data) error {
	templateContent, err := templates.Asset(t.Name)
	if err != nil {
		return err
	}

	codeTmpl, err := template.New("code").
		Funcs(template.FuncMap{
			"emptyValue": emptyValue,
			"isPointer":  isPointer,
			"isNumber":   isNumber,
		}).
		Parse(string(templateContent))
	if err != nil {
		return err
	}

	return codeTmpl.Execute(writer, data)
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
