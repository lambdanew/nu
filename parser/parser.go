package parser

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

// regex patterns for various rules
// see https://github.com/google/re2/wiki/Syntax

const (
	IdentDef      = `[a-zA-Z]\w*`
	NumberDef     = `[-+]?(\d*\.)?\d+`
	AddressDef    = `#[a-fA-F0-9]{20}`
	StringDef     = `"(\\"|[^"])*"`
	PunctDef      = `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`
	WhiteSpaceDef = `[ \t\n\r]+`
)

type (
	Command struct {
		Verb *Verb `@@`
	}

	Verb struct {
		Name      string        `@Ident`
		DirectObj *DirectObject `@@*`
	}

	DirectObject struct {
		Name        string          `(@Address | @Ident | @String)`
		IndirectObj *IndirectObject `@@*`
	}

	IndirectObject struct {
		PrePosition string `@(
			(("in" "front" "of") | "in" | "inside" | "into") | 
			("at" | "to") | ("with" | "using") | 
			(("on" "top" "of") | "on" | "onto" | "upon") |
			(("out" "of") | ("from" "inside") | "from") |
			("under" | "underneath" | "beneath" | "behind" | "beside") |
			(("off" "of") | "off" | "through" | "over" | "is" | "as" | "for" | "about")
			)?`

		Name string `(@Address | @Ident | @String)`
	}
)

var (
	basicLexer = lexer.MustSimple([]lexer.Rule{
		{"Ident", IdentDef, nil},
		{"Number", NumberDef, nil},
		{"String", StringDef, nil},
		{"Address", AddressDef, nil},
		{"Punct", PunctDef, nil},
		{"Whitespace", WhiteSpaceDef, nil},
	})

	parser = participle.MustBuild(&Command{},
		participle.Lexer(basicLexer),
		//participle.Elide("Comment", "Whitespace"),
		participle.Unquote("String"),
		participle.Elide("Whitespace"),
		participle.UseLookahead(2),
	)
)

func ParseCommand(cmd string) (*Command, error) {
	ast := &Command{}
	if err := parser.ParseString("", cmd, ast); err != nil {
		return nil, err
	}
	return ast, nil
}

func (cmd *Command) Dump() {
	fmt.Printf("%s\n", cmd.ToString())
}

func (cmd *Command) ToString() string {
	obj := "<EMPTY>"
	if cmd.Verb.DirectObj != nil {
		obj = cmd.Verb.DirectObj.ToString()
	}
	return fmt.Sprintf("verb: '%s' direct object: %s", cmd.Verb.Name, obj)
}

func (obj *DirectObject) ToString() string {
	ind := "<EMPTY>"
	if obj.IndirectObj != nil {
		ind = obj.IndirectObj.ToString()
	}
	return fmt.Sprintf("'%s' indirect object: %s", obj.Name, ind)
}

func (obj *IndirectObject) ToString() string {
	return fmt.Sprintf("'%s' preposition: '%s'", obj.Name, obj.PrePosition)
}
