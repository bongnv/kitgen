package gen

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPointer(t *testing.T) {
	require.True(t, isPointer("*model.Profile"))
	require.False(t, isPointer("model.Profile"))
}

func Test_executeTemplate(t *testing.T) {
	t.Run("without-func", func(t *testing.T) {
		buf := &bytes.Buffer{}
		err := executeTemplate("{{ . }}", buf, "data")
		require.NoError(t, err)
		require.Equal(t, "data", buf.String())
	})

	t.Run("with-func", func(t *testing.T) {
		buf := &bytes.Buffer{}
		err := executeTemplate("{{ . | isPointer }}", buf, "*data")
		require.NoError(t, err)
		require.Equal(t, "true", buf.String())
	})
}
