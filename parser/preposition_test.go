package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerbObjectUnknownPreposition1(t *testing.T) {
	cmd := "put book UNKNOWN bag"
	ast, err := ParseCommand(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)
	assert.Equal(t, "book", ast.Verb.DirectObj.Name)
	assert.Equal(t, "bag", ast.Verb.DirectObj.IndirectObj.Name)
	assert.Empty(t, "", ast.Verb.DirectObj.IndirectObj.PrePosition)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestVerbObjectUnknownPreposition2(t *testing.T) {
	cmd := "put book UNKNOWN PRE bag"
	ast, err := ParseCommand(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)
	assert.Equal(t, "book", ast.Verb.DirectObj.Name)
	assert.Equal(t, "bag", ast.Verb.DirectObj.IndirectObj.Name)
	assert.Empty(t, "", ast.Verb.DirectObj.IndirectObj.PrePosition)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestPrePositions1(t *testing.T) {
	// in/inside/into
	cmd := "verb obj1 in obj2"
	ast, err := ParseCommand(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 inside obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 into obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestPrePositions2(t *testing.T) {
	// at/to
	cmd := "verb obj1 at obj2"
	ast, err := ParseCommand(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 to obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestPrePositions3(t *testing.T) {
	// with/using
	cmd := "verb obj1 with obj2"
	ast, err := ParseCommand(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 using obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestPrePositions4(t *testing.T) {
	// on top of/on/onto/upon
	cmd := "verb obj1 on top of obj2"
	ast, err := ParseCommand(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 on obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 onto obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 upon obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestPrePositions5(t *testing.T) {
	// out of/from inside/from
	cmd := "verb obj1 out of obj2"
	ast, err := ParseCommand(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 from inside obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 from obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestPrePositions6(t *testing.T) {
	// under/underneath/beneath/in front of/behind/beside
	cmd := "verb obj1 under obj2"
	ast, err := ParseCommand(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 underneath obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 beneath obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 in front of obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 behind obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 beside obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}

func TestPrePositions7(t *testing.T) {
	// off/off of/over/is/as/for/about
	cmd := "verb obj1 off obj2"
	ast, err := ParseCommand(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 off of obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 over obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 through obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 is obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 as obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 for obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())

	cmd = "verb obj1 about obj2"
	ast, err = ParseCommand(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, ast)
	assert.NotNil(t, ast.Verb.DirectObj)
	assert.NotNil(t, ast.Verb.DirectObj.IndirectObj)

	fmt.Printf("'%s' -> %s\n", cmd, ast.ToString())
}
