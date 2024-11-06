package lexer

import (
	"bufio"
	"fmt"
	"io"
)

// Wraps an io.Reader, and allows creating multiple readers that all read from a
// specific point in the stream, i.e.
//
// this type is itself not a an io.Reader
// TODO, performance could _potentially_ be improved by letting multiple threads
// search for tokens in parallel; in which case this needs to be made thread safe
type forkableReader struct {
	reader io.Reader
	cache  []byte
	eof    bool
	pos    int
}

func (b *forkableReader) debug() {
	fmt.Printf("Pos %d\nConsumed: %s\nRemaining:%s\n", b.pos, b.cache[:b.pos], b.cache[b.pos:])
}

type fork struct {
	buffer *forkableReader
	pos    int
}

func newBuffer(reader io.Reader) *forkableReader {
	return &forkableReader{reader: reader}
}

func (b *forkableReader) fork() *fork {
	return &fork{b, b.pos}
}

func (b *forkableReader) forkRuneReader() io.RuneReader {
	return bufio.NewReader(b.fork())
}

func (b *forkableReader) advanceCache(count int) error {
	tmp := make([]byte, count)
	read, err := b.reader.Read(tmp)
	b.cache = append(b.cache, tmp[:read]...)
	if err == io.EOF {
		b.eof = true
		err = nil
	}
	return err
}

func (b *forkableReader) advance(length int) {
	if len(b.cache) < (length + b.pos) {
		panic("Cannot advance beyone what has been read")
	}
	b.pos += length
}

func (b *forkableReader) subString(start int, end int) string {
	return (string)(b.cache[b.pos+start : b.pos+end])
}

func (f *fork) Read(b []byte) (count int, err error) {
	if f.buffer.eof && f.pos >= len(f.buffer.cache) {
		return 0, io.EOF
	}
	cacheSize := len(f.buffer.cache)
	desiredCacheSize := len(b) // TODO: This is wrong
	if desiredCacheSize > cacheSize {
		err = f.buffer.advanceCache(desiredCacheSize - cacheSize)
	}
	count = copy(b, f.buffer.cache[f.pos:])
	f.pos += count
	if f.buffer.eof && err == nil {
		err = io.EOF
	}
	return
}
