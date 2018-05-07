package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/benchlab/mir/client"
	"github.com/benchlab/benchcore/config"
	"github.com/benchlab/benchcore/mempool"
	"github.com/benchlab/benchcore/proxy"
	"github.com/benchlab/benchcore/types"
)

func runServer(addr string) chan<- bool {
	shutdownChan := make(chan bool)
	go func() {
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			// Shutdown
			<-shutdownChan
			ln.Close()
		}()

		i := uint64(0)
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("conn err:: %d: %v", i, err)
			} else {
				go handleConnection(conn)
			}
		}
	}()
	return shutdownChan
}

func handleConnection(conn net.Conn) {
	data, err := ioutil.ReadAll(conn)
	log.Printf("err: %v data: %d\n", err, len(data))
}

func main() {
	path := flag.String("file", "", "the input file")
	flag.Parse()

	const addr = "0.0.0.0:8181"
	shutdownChan := runServer(addr)
	defer func() {
		shutdownChan <- true
	}()

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d\n", greenrunIt(addr, data))
	foo := make(chan bool)
	go func() {
		<-foo
		close(foo)
	}()
	<-foo
}

func greenrunIt(addr string, data []byte) int {
	c, err := mircli.NewClient(addr, "socket", false)
	if err != nil {
		panic(err)
	}

	app := proxy.NewAppConnMempool(c)
	mcfg := config.DefaultMempoolConfig()
	mp := mempool.NewMempool(mcfg, app, 1)

	if err := mp.AuditNtx(types.Ntx(data), nil); err != nil {
		return 0
	}

	return 1
}
