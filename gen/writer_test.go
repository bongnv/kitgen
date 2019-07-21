package gen

import (
	"bytes"
	"os"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/require"
)

func Test_getWriter(t *testing.T) {
	t.Run("non-nil-writer", func(t *testing.T) {
		p := &pipeline{
			writer: &bytes.Buffer{},
		}

		w, err := getWriter(p)
		require.NoError(t, err)
		require.IsType(t, w, &bytes.Buffer{})
	})

	t.Run("stdout", func(t *testing.T) {
		p := &pipeline{
			opts: &Option{},
		}

		w, err := getWriter(p)
		require.NoError(t, err)
		require.Equal(t, os.Stdout, w)
	})
}

func Test_writeToFile(t *testing.T) {
	p := &pipeline{
		opts: &Option{
			Output: randomdata.Letters(10),
			Dir:    ".",
		},
		buf: []byte("rand data"),
	}
	err := writeToFile(p)

	defer func() {
		t.Logf("Removing file with result: %v.", os.Remove(p.opts.Output))
	}()
	require.NoError(t, err)
}
