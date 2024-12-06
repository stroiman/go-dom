package main

import (
	"fmt"
	"io"
)

func WriteHeader(b *builder) {
	b.Printf("package scripting\n\n")
	b.Printf("// This file is generated. Do not edit.\n\n")
}

type builder struct {
	io.Writer
	indentLvl int
}

func newBuilder(w io.Writer) *builder {
	return &builder{w, 0}
}

func (b builder) Printf(format string, args ...interface{}) {
	for i := 0; i < b.indentLvl; i++ {
		fmt.Fprint(b.Writer, "\t")
	}
	fmt.Fprintf(b.Writer, format, args...)
}

func (b *builder) indent() {
	b.indentLvl++
}

func (b *builder) indentF(format string, args ...interface{}) {
	b.indent()
	b.Printf(format, args...)
}

func (b *builder) unIndent() {
	b.indentLvl--
	if b.indentLvl < 0 {
		panic("More unindentation than indentation")
	}
}

func (b *builder) unIndentF(format string, args ...interface{}) {
	b.unIndent()
	b.Printf(format, args...)
}
