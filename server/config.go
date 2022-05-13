package server

import (
	"fmt"
	"github.com/vela-security/vela-public/auxlib"
	"github.com/vela-security/vela-public/lua"
	"github.com/vela-security/vela-public/pipe"
	"sync"
)

type config struct {
	Name     string
	Bind     auxlib.URL
	Resolver string
	mutex    sync.RWMutex
	direct   map[string]lua.LValue
	route    []route

	pipe *pipe.Px
	sdk  *pipe.Px
	co   *lua.LState
}

func newConfig(L *lua.LState) *config {
	tab := L.CheckTable(1)

	bind, _ := auxlib.NewURL("udp://127.0.0.1:53?timeout=1")
	cfg := &config{
		Bind:     bind,
		Resolver: "114.114.114.114:53",
		direct:   make(map[string]lua.LValue),
		pipe:     pipe.New(),
		sdk:      pipe.New(),
	}

	tab.Range(func(key string, val lua.LValue) {
		switch key {
		case "name":
			cfg.Name = val.String()
		case "bind":
			cfg.Bind = auxlib.CheckURL(val, L)
		case "resolver":
			cfg.Resolver = val.String()
		}
	})

	if e := cfg.verify(); e != nil {
		L.RaiseError("%v", e)
		return nil
	}

	cfg.co = xEnv.Clone(L)
	return cfg
}

func (cfg *config) verify() error {
	if e := auxlib.Name(cfg.Name); e != nil {
		return e
	}

	if cfg.Bind.IsNil() {
		return fmt.Errorf("not found bind url")
	}

	return nil
}
