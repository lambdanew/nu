package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRuntime(t *testing.T) {
	rt, err := New()

	require.NoError(t, err)
	require.NotNil(t, rt)
	assert.NotNil(t, rt.root)
}
