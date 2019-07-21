package gen

import (
	"io"
	"log"
	"os"
)

func getWriter(p *pipeline) (io.Writer, error) {
	if p.writer != nil {
		return p.writer, nil
	}

	if p.opts.Output == "" {
		return os.Stdout, nil
	}

	f, err := os.Create(p.opts.Output)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func writeToFile(p *pipeline) error {
	w, err := getWriter(p)
	if err != nil {
		return err
	}

	if closeable, ok := w.(io.Closer); ok {
		defer func() {
			err := closeable.Close()
			if err != nil {
				log.Printf("Failed to close file with error %v.\n", err)
			}
		}()
	}

	if _, err := w.Write(p.buf); err != nil {
		return err
	}

	return nil
}
