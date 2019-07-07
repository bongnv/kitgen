package gen

import (
	"io"
	"io/ioutil"

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

	return executeTemplate(string(templateContent), writer, data)
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

	return executeTemplate(string(templateContent), writer, data)
}
