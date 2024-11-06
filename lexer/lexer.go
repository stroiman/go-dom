package lexer

import (
	"fmt"
	"io"
	"regexp"
)

type regexHandler func(lex *lexer, m match)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	tokens   []Token
	source   *forkableReader
	pos      int
}

func Tokenize(source io.Reader) []Token {
	lex := createLexer(source)

	for !lex.at_eof() {
		matched := false
		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindReaderSubmatchIndex(lex.remainder())
			if loc != nil {
				if loc[0] != 0 {
					panic("Error in code; match should start at beginning of string")
				}
				pattern.handler(lex, match{lex.source, loc})
				matched = true
				break
			}
		}

		if !matched {
			panic(
				fmt.Sprintf(
					"Lexer::Error -> unrecognized token near %s",
					string(lex.remainderString()),
				),
			)
		}
	}
	lex.push(NewToken(EOF, ""))
	return lex.tokens
}

type match struct {
	b       *forkableReader
	matches []int
}

func (m match) getMatch() string {
	if len(m.matches) < 4 {
		return ""
	}
	return m.b.subString(m.matches[2], m.matches[3])
}

func (m match) getLength() int {
	return m.matches[1]
}

func (l *lexer) advance(count int) {
	l.pos += count
	l.source.pos += count
}

func (l *lexer) push(token Token) {
	l.tokens = append(l.tokens, token)
}

func (l *lexer) at_eof() bool {
	return l.source.eof && l.pos >= len(l.source.cache)
}

func (l *lexer) remainderString() string {
	return string(l.source.cache[l.pos:])
}

func (l *lexer) remainder() io.RuneReader {
	return l.source.forkRuneReader()
}

func createPattern(s string, kind TokenKind) regexPattern {
	return regexPattern{
		// Match the beginning of the string. When searching for a token, we are
		// always looking for the next token.
		regexp.MustCompile("^" + s),
		func(lex *lexer, m match) {
			lex.push(NewToken(kind, m.getMatch()))
			lex.advance(m.getLength())
		},
	}
}

func createLexer(source io.Reader) *lexer {
	return &lexer{
		pos:    0,
		source: newBuffer(source),
		tokens: []Token{},
		patterns: []regexPattern{
			createPattern("<([[:alpha:]]+)", TAG_OPEN_BEGIN),
			createPattern("</([[:alpha:]]+)", TAG_CLOSE_BEGIN),
			createPattern("/>", TAG_CLOSE_END), // Unused atm
			createPattern(">", TAG_END),
			createPattern(`\w+`, IDENTIFIER), // Unused atm
		},
	}
}
