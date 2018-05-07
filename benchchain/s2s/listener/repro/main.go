package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/benchlab/benchcore/s2s"
	lg "github.com/benchlab/benchlibs/log"
)

func main() {
	inFile := flag.String("if", "", "the file with the <protocol>,<address> combination")
	flag.Parse()

	data, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatalf("%q: %v", *inFile, err)
	}
	data = bytes.TrimSuffix(data, []byte("\n"))
	splits := strings.Split(string(data), ",")
	if len(splits) < 1 {
		log.Fatal("expecting `data-csv` to be of the form <protocol>,<address>")
	}
	protocol, addr := splits[0], splits[1]
	buf := new(bytes.Buffer)
	llg := lg.NewBCLogger(buf)
	ln := s2s.NewDefaultListener(protocol, addr, true, llg)
	ln.Stop()
}
