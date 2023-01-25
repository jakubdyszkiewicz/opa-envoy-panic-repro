package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa-envoy-plugin/plugin"
	"github.com/open-policy-agent/opa/runtime"
	"github.com/open-policy-agent/opa/storage"
)

const policies = 2
const policy = `
package envoy.authz

import input.attributes.request.http as http_request

default allow = false

allow {
  http_request.method == "GET"
}
`

func main() {
	ctx := context.Background()
	runtime.RegisterPlugin(plugin.PluginName, plugin.Factory{})
	params := runtime.NewParams()
	params.Addrs = &[]string{":9999"}
	params.ConfigOverrides = append(params.ConfigOverrides, "plugins.envoy_ext_authz_grpc.addr=127.0.0.1:8888")
	rt, err := runtime.NewRuntime(ctx, params)
	if err != nil {
		panic(err)
	}
	tx, err := rt.Store.NewTransaction(ctx, storage.TransactionParams{Write: true})
	if err != nil {
		panic(err)
	}

	for i := 0; i < policies; i++ {
		if err := rt.Store.UpsertPolicy(ctx, tx, fmt.Sprintf("policy-%d", i), []byte(policy)); err != nil {
			panic(err)
		}
	}
	if err := rt.Store.Commit(ctx, tx); err != nil {
		panic(err)
	}

	println("string runtime...")
	if err := rt.Serve(ctx); err != nil {
		panic(err)
	}
}
