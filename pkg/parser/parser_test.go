package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildParser(t *testing.T) {
	require.NotNil(t, parser)
}

func TestEmptyCommand(t *testing.T) {
	ast, err := ParseCommand("")
	assert.Error(t, err)
	assert.Nil(t, ast)
}

func TestSimpleCommand(t *testing.T) {
	cmd := "look"
	ast, err := ParseCommand(cmd)

	require.NoError(t, err)
	assert.NotNil(t, ast)
	assert.Nil(t, ast.Verb.DirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestVerbObj(t *testing.T) {
	cmd := "take watch"
	ast, err := ParseCommand(cmd)

	require.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestVerbObjAddress(t *testing.T) {
	cmd := "take #0123456789abcdefffff"
	ast, err := ParseCommand(cmd)

	require.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestVerbObjComplexName(t *testing.T) {
	cmd := "follow \"The White Rabbit\""
	ast, err := ParseCommand(cmd)

	require.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.Nil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestVerbObjectIndirectObject1(t *testing.T) {
	cmd := "put book in bag"
	ast, err := ParseCommand(cmd)

	require.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestVerbObjIndirectOb2(t *testing.T) {
	cmd := "put \"The White Rabbit\" into bag"
	ast, err := ParseCommand(cmd)

	require.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestVerbObjIndirectObj3(t *testing.T) {
	cmd := "put \"The White Rabbit\" into \"small bag\""
	ast, err := ParseCommand(cmd)

	require.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}
