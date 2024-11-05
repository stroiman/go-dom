package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	tokens   []Token
	source   string
	pos      int
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	for !lex.at_eof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())
			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("Lexer::Error -> unrecognized token near %s", lex.remainder))
		}
	}
	lex.push(NewToken(EOF, "EOF"))
	return lex.tokens
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advance(len(value))
		lex.push(NewToken(kind, value))
	}
}

func stringHandler(l *lexer, regex *regexp.Regexp) {
	match := regex.FindString(l.remainder())
	l.push(NewToken(IDENTIFIER, match))
	l.advance(len(match))
}

func (l *lexer) advance(count int) {
	l.pos += count
}

func (l *lexer) push(token Token) {
	l.tokens = append(l.tokens, token)
}

func (l *lexer) at() byte {
	return l.source[l.pos]
}

func (l *lexer) at_eof() bool {
	return l.pos >= len(l.source)
}

func (l *lexer) remainder() string {
	return l.source[l.pos:]
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		tokens: []Token{},
		patterns: []regexPattern{
			{regexp.MustCompile("</"), defaultHandler(TAG_CLOSE_START, "</")},
			{regexp.MustCompile("<"), defaultHandler(TAG_START, "<")},
			{regexp.MustCompile("/>"), defaultHandler(TAG_CLOSE_END, "/>")},
			{regexp.MustCompile(">"), defaultHandler(TAG_END, ">")},
			{regexp.MustCompile(`\w+`), stringHandler},
		},
	}
}
