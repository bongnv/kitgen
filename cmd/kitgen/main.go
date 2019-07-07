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
	templateFile := flag.String("t", "", "Path to template file")
	identifier := flag.String("i", "", "Identifier of type to parse")
	dir := flag.String("d", ".", "Path to source directory")
	flag.Parse()

	if *templateFile == "" || *identifier == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := gen.Process(&gen.Options{
		TemplateFile: *templateFile,
		Identifier:   *identifier,
		Dir:          *dir,
	})

	if err != nil {
		exitWithErr(err)
	}
}
