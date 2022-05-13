package client

import (
	"github.com/miekg/dns"
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
)

var xEnv assert.Environment

func (dc *dnsClient) Query(L *lua.LState) lua.LValue {

	str := L.CheckString(1) + "."
	msg := dns.Msg{}
	msg.SetQuestion(str, dns.TypeANY)
	cli := dc.Client()

	r, rtt, err := cli.Exchange(&msg, dc.cfg.Resolve)
	return &Reply{r, rtt, err}
}

func newLuaDnsClient(L *lua.LState) int {
	cfg := newConfig(L)
	cli := newDnsClient(cfg)
	L.Push(L.NewAnyData(cli))
	return 1
}

func WithEnv(env assert.Environment, kv lua.UserKV) {
	kv.Set("client", lua.NewFunction(newLuaDnsClient))
	xEnv = env
}
