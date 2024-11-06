package lexer

import "fmt"

type TokenKind int

const (
	TAG_OPEN_BEGIN = iota
	TAG_CLOSE_BEGIN
	TAG_START       // Start of tag: <
	TAG_CLOSE_START // Start of close tag </
	TAG_END         // >
	TAG_CLOSE_END   // self-closing tag />
	IDENTIFIER
)

type Token struct {
	Kind TokenKind
	Data string
}

func NewToken(kind TokenKind, data string) Token { return Token{kind, data} }

var names = map[TokenKind]string{
	TAG_OPEN_BEGIN:  "<",
	TAG_CLOSE_BEGIN: "</",
	TAG_START:       "<",
	TAG_CLOSE_START: "</",
	TAG_END:         ">",
	TAG_CLOSE_END:   "/>",
	IDENTIFIER:      "ident",
}

func (k TokenKind) String() string {
	if val, ok := names[k]; ok {
		return val
	}
	return "?"
}

func (t Token) String() string {
	if t.Data != "" {
		return fmt.Sprintf("%s (%s)", t.Kind, t.Data)
	} else {
		return t.Kind.String()
	}
}
