package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/benchlab/astro"
	"github.com/benchlab/benchcore/s2s"
)

var privKey = astro.GenPrivKeyEd25519()

type rwNopCloser struct {
	io.ReadWriter
}

func (rwn *rwNopCloser) Close() error { return nil }

func main() {
	source := flag.String("src", "", "the file containing the bytes to use")
	flag.Parse()

	data, err := ioutil.ReadFile(*source)
	if err != nil {
		log.Fatalf("reading %q err: %v", *source, err)
	}

	rwc := &rwNopCloser{new(bytes.Buffer)}
	sc, err := s2s.MakeSecretConnection(rwc, privKey)
	if err != nil {
		panic(fmt.Errorf("for some reason the connection making failed, err: %v", err))
	}
	defer sc.Close()
	n, err := sc.Write(data)
	if err != nil {
		panic(err)
	}
	if g, w := n, len(data); g != w {
		panic(fmt.Errorf("n: got %d; want %d", g, w))
	}
}
