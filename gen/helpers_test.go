package gen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPointer(t *testing.T) {
	require.True(t, isPointer("*model.Profile"))
	require.False(t, isPointer("model.Profile"))
}
