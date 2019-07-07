package gen

import (
	"bytes"
	"io"
	"log"
	"os"

	"golang.org/x/tools/imports"
)

type Options struct {
	TemplateFile string
	Identifier   string
	Dir          string
}

func Process(opts *Options) error {
	c := &Command{
		DataParser: &interfaceParser{
			Name: opts.Identifier,
		},
		Generator: &fileTemplate{
			filePath: opts.TemplateFile,
		},
		Writer:   os.Stdout,
		FilePath: opts.Dir + "noop.go",
		CodePath: opts.Dir,
	}

	return c.Process()
}

type Command struct {
	DataParser DataParser
	Generator  Generator
	Writer     io.Writer
	FilePath   string
	CodePath   string
}

func (c *Command) Process() error {
	pkgs, err := parsePath(c.CodePath)
	if err != nil {
		return err
	}

	d, err := c.DataParser.Parse(pkgs)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}

	err = c.Generator.Generate(buf, d)
	if err != nil {
		return err
	}

	log.Println(buf.String())

	out, err := imports.Process(c.FilePath, buf.Bytes(), nil)
	if err != nil {
		return err
	}

	_, err = c.Writer.Write(out)
	return err
}
