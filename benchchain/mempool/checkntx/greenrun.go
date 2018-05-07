package checktx

import (
	"github.com/benchlab/mir/client"
	"github.com/benchlab/benchcore/config"
	"github.com/benchlab/benchcore/mempool"
	"github.com/benchlab/benchcore/proxy"
	"github.com/benchlab/benchcore/types"
)

const addr = "0.0.0.0:8080"

func GreenRuns(data []byte) int {
	c, err := mircli.NewClient(addr, "socket", false)
	if err != nil {
		panic(err)
	}

	app := proxy.NewAppConnMempool(c)
	mcfg := config.DefaultMempoolConfig()
	mp := mempool.NewMempool(mcfg, app, 1)
	defer mp.CloseWAL()

	if err := mp.AuditNtx(types.Ntx(data), nil); err != nil {
		return 0
	}

	return 1
}
