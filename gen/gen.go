package gen

import (
	"bytes"
	"io"
	"path/filepath"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

// Option includes golang options
type Option struct {
	TemplateFile string
	TemplateName string
	Identifier   string
	Dir          string
	Output       string
}

// Process takes a list of options and generate codes based on those options.
func Process(opts *Option) error {
	p := &pipeline{
		opts: opts,
	}

	return execute(p)
}

type pipeline struct {
	opts   *Option
	pkgs   []*packages.Package
	data   *Data
	buf    []byte
	writer io.Writer
}

func execute(p *pipeline) error {
	steps := []func(*pipeline) error{
		parseSource,
		parseTemplateData,
		generateFromTemplate,
		formatOutput,
		writeToFile,
	}

	for _, step := range steps {
		if err := step(p); err != nil {
			return err
		}
	}

	return nil
}

func parseSource(p *pipeline) error {
	parseMode := packages.NeedName |
		packages.NeedFiles |
		packages.NeedImports |
		packages.NeedDeps |
		packages.NeedCompiledGoFiles |
		packages.NeedTypes |
		packages.NeedSyntax |
		packages.NeedTypesInfo

	pkgs, err := packages.Load(
		&packages.Config{
			Mode: parseMode,
		},
		p.opts.Dir,
	)

	if err != nil {
		return err
	}

	p.pkgs = pkgs
	return nil
}

func parseTemplateData(p *pipeline) error {
	parser := &interfaceParser{
		Name: p.opts.Identifier,
	}

	d, err := parser.Parse(p.pkgs)
	if err != nil {
		return err
	}

	p.data = d
	return nil
}

func generateFromTemplate(p *pipeline) error {
	var g Generator
	if p.opts.TemplateName > "" {
		g = &inMemoryTemplate{
			Name: p.opts.TemplateName,
		}
	} else {
		g = &fileTemplate{
			filePath: p.opts.TemplateFile,
		}
	}

	buf := &bytes.Buffer{}
	if err := g.Generate(buf, p.data); err != nil {
		return err
	}

	p.buf = buf.Bytes()
	return nil
}

func formatOutput(p *pipeline) error {
	sourcePath := p.opts.Output
	if p.opts.Output == "" {
		sourcePath = filepath.Join(p.opts.Dir, "z_auto_generated.go")
	}

	out, err := imports.Process(sourcePath, p.buf, nil)
	if err != nil {
		return err
	}

	p.buf = out
	return nil
}
