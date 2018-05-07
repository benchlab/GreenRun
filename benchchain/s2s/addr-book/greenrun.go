package addr

import (
	"encoding/json"
	"fmt"
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

func GreenRuns(data []byte) int {
	srcAddr := new(s2s.NetAddress)
	if err := json.Unmarshal(data, srcAddr); err != nil {
		return -1
	}
	addr := randAddress()
	addrBook.AddAddress(addr, srcAddr)
	bias := rand.Intn(1000)
	if p := addrBook.PickAddress(bias); p == nil {
		blob, _ := json.Marshal(addr)
		panic(fmt.Sprintf("picked a nil address with bias: %d, netAddress: %s addr: %s", bias, data, blob))
	}
	return 0
}
