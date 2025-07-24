package main

import (
	"io"
	"log"
	"strings"
	"testing"
)

type LimitedReader struct {
	r io.Reader
	n int
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	log.Println("start to read")
	if l.n <= 0 {
		return 0, io.EOF
	}

	p = make([]byte, l.n)
	log.Printf("-------- p? len %d, cap %d\n", len(p), cap(p))
	readBytes, err := l.r.Read(p)
	log.Println(readBytes, string(p))
	log.Printf("adjusted p? len %d, cap %d\n", len(p), cap(p))

	return readBytes, io.EOF
}

func TestLimitedReader(t *testing.T) {
	log.SetFlags(0)

	src := strings.NewReader("string to read")
	ll := &LimitedReader{r: src, n: 1024}
	readed, err := io.ReadAll(ll)
	if err != nil {
		t.Fail()
	}

	log.Printf("read - %s", string(readed))

	// src2 := strings.NewReader("ok")
	// ll2 := &LimitedReader{r: src2, n: 5}
	// io.ReadAll(ll2)
}
