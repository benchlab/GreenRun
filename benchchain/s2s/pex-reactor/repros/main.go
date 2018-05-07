package main

import (
	"io/ioutil"
	lg "log"

	"github.com/benchlab/astro"
	"github.com/benchlab/benchcore/s2s"
	"github.com/benchlab/benchlibs/log"
)

func main() {
	blob, err := ioutil.ReadFile("da39a3ee5e6b4b0d3255bfef95601890afd80709")
	if err != nil {
		lg.Fatal(err)
	}
	GreenRuns(blob)
}

func GreenRuns(data []byte) int {
	addrB := s2s.NewAddrBook("addrbook1", false)
	prec := s2s.NewPEXReactor(addrB)
	if prec == nil {
		panic("nil Reactor")
	}
	satellite1 := newGreenRunsSatellite()
	prec.AddSatellite(satellite1)
	prec.Receive(0x01, satellite1, data)
	return 1
}

type greenrunSatellite struct {
	m map[string]interface{}
}

var _ s2s.Satellite = (*greenrunSatellite)(nil)

func newGreenRunsSatellite() *greenrunSatellite {
	return &greenrunSatellite{m: make(map[string]interface{})}
}

func (fp *greenrunSatellite) Key() string {
	return "greenrun-satellite"
}

func (fp *greenrunSatellite) Send(ch byte, data interface{}) bool {
	return true
}

func (fp *greenrunSatellite) TrySend(ch byte, data interface{}) bool {
	return true
}

func (fp *greenrunSatellite) Set(key string, value interface{}) {
	fp.m[key] = value
}

func (fp *greenrunSatellite) Get(key string) interface{} {
	return fp.m[key]
}

var ed25519Key = astro.GenPrivKeyEd25519()
var defaultNodeInfo = &s2s.NodeInfo{
	PubKey:     ed25519Key.PubKey().Unwrap().(astro.PubKeyEd25519),
	Moniker:    "foo1",
	Version:    "tender-greenrun",
	RemoteAddr: "0.0.0.0:98991",
	ListenAddr: "0.0.0.0:98992",
}

func (fp *greenrunSatellite) IsOutbound() bool             { return false }
func (fp *greenrunSatellite) IsPersistent() bool           { return false }
func (fp *greenrunSatellite) IsRunning() bool              { return false }
func (fp *greenrunSatellite) NodeInfo() *s2s.NodeInfo      { return defaultNodeInfo }
func (fp *greenrunSatellite) OnReset() error               { return nil }
func (fp *greenrunSatellite) OnStart() error               { return nil }
func (fp *greenrunSatellite) OnStop()                      {}
func (fp *greenrunSatellite) Reset() (bool, error)         { return true, nil }
func (fp *greenrunSatellite) Start() (bool, error)         { return true, nil }
func (fp *greenrunSatellite) Stop() bool                   { return true }
func (fp *greenrunSatellite) String() string               { return fp.Key() }
func (fp *greenrunSatellite) Status() s2s.ConnectionStatus { var cs s2s.ConnectionStatus; return cs }
func (fp *greenrunSatellite) SetLogger(l log.Logger)       {}
