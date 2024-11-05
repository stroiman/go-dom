package parser

import (
	"github.com/stroiman/go-dom/interfaces"
	"github.com/stroiman/go-dom/lexer"
)

type bindingPower int

const (
	default_bp bindingPower = iota
)

type elementHandler func(p *parser) interfaces.Element

type elementLookup map[lexer.TokenKind]elementHandler
type bpLookup map[lexer.TokenKind]bindingPower

var bp_lu = bpLookup{}
var element_lo = elementLookup{}
