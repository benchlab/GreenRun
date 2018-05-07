package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net"

	"github.com/benchlab/benchcore/s2s"
)

var addrBook = s2s.NewAddrBook("./testdata/addrbook.json", true)

// random octet
func ro() byte { return byte(rand.Intn(256)) }

func randAddress() *s2s.NetAddress {
	return &s2s.NetAddress{
		IP:   net.IPv4(ro(), ro(), ro(), ro()),
		Port: uint16(rand.Intn(1 << 16)),
	}
}

func main() {
	srcStr := flag.String("src", "", "the JSON for the source")
	biasVar := flag.Int("bias", 412, "the bias to use in picking the address")
	addrStr := flag.String("addr", "", "the JSON for the address")
	flag.Parse()

	srcAddr := new(s2s.NetAddress)
	if err := json.Unmarshal([]byte(*srcStr), srcAddr); err != nil {
		log.Fatal(err)
	}
	addr := new(s2s.NetAddress)
	if err := json.Unmarshal([]byte(*addrStr), addr); err != nil {
		log.Fatal(err)
	}
	addrBook.AddAddress(addr, srcAddr)
	bias := rand.Intn(*biasVar)
	if addr := addrBook.PickAddress(bias); addr == nil {
		panic("picked a nil address")
	}
}
