package vdns

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
	"github.com/vela-security/vela-dns/client"
	"github.com/vela-security/vela-dns/server"
)

func WithEnv(env assert.Environment) {
	kv := lua.NewUserKV()
	client.WithEnv(env, kv)
	server.LuaInjectApi(env, kv)
	env.Set("dns", kv)
}
