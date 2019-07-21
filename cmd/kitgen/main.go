package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bongnv/kitgen/gen"
)

func exitWithErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	opts := parseOptionsFromCmd()
	if err := gen.Process(opts); err != nil {
		exitWithErr(err)
	}
}

func parseOptionsFromCmd() *gen.Option {
	templateFile := flag.String("t", "", "Path to template file")
	identifier := flag.String("i", "", "Identifier of type to parse")
	dir := flag.String("d", ".", "Path to source directory")
	out := flag.String("o", "", "Write to file")

	flag.Parse()

	opts := &gen.Option{
		TemplateFile: *templateFile,
		Identifier:   *identifier,
		Dir:          *dir,
		Output:       *out,
	}

	validateOptions(opts)
	return opts
}

func validateOptions(opts *gen.Option) {
	if opts.TemplateFile == "" || opts.Identifier == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
