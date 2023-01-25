# OPA Envoy Panic repro

1. Run Envoy
```sh
envoy --config-path envoy_config.yaml
```

2. Run OPA agent
```sh
go run main.go
```

3. Send request
```sh
curl localhost:10000/ -v 
```

4. See the result
```
string runtime...
{"addrs":[":9999"],"diagnostic-addrs":[],"level":"info","msg":"Initializing server.","time":"2023-01-25T16:06:30+01:00"}
{"addr":"127.0.0.1:8888","dry-run":false,"enable-reflection":false,"level":"info","msg":"Starting gRPC server.","path":"envoy/authz/allow","query":"","time":"2023-01-25T16:06:30+01:00"}
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x20 pc=0x1056ddfd4]

goroutine 9 [running]:
github.com/open-policy-agent/opa/topdown.evalVirtual.eval({0x14000146c00, {0x140003fc920, 0x4, 0x4}, {0x140003fcba0, 0x4, 0x4}, 0x3, 0x14000352c60, 0x14000348378, ...}, ...)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:2266 +0x84
github.com/open-policy-agent/opa/topdown.evalTree.next({0x14000146c00, {0x140003fc920, 0x4, 0x4}, {0x140003fcba0, 0x4, 0x4}, 0x3, 0x14000352c60, 0x14000348378, ...}, ...)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:2101 +0x260
github.com/open-policy-agent/opa/topdown.evalTree.eval({0x14000146c00, {0x140003fc920, 0x4, 0x4}, {0x140003fcba0, 0x4, 0x4}, 0x3, 0x14000352c60, 0x14000348378, ...}, ...)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:2054 +0xd0
github.com/open-policy-agent/opa/topdown.evalTree.next({0x14000146c00, {0x140003fc920, 0x4, 0x4}, {0x140003fcba0, 0x4, 0x4}, 0x2, 0x14000352c60, 0x14000348378, ...}, ...)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:2107 +0x2b0
github.com/open-policy-agent/opa/topdown.evalTree.eval({0x14000146c00, {0x140003fc920, 0x4, 0x4}, {0x140003fcba0, 0x4, 0x4}, 0x2, 0x14000352c60, 0x14000348378, ...}, ...)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:2054 +0xd0
github.com/open-policy-agent/opa/topdown.evalTree.next({0x14000146c00, {0x140003fc920, 0x4, 0x4}, {0x140003fcba0, 0x4, 0x4}, 0x1, 0x14000352c60, 0x14000348378, ...}, ...)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:2107 +0x2b0
github.com/open-policy-agent/opa/topdown.evalTree.eval({0x14000146c00, {0x140003fc920, 0x4, 0x4}, {0x140003fcba0, 0x4, 0x4}, 0x1, 0x14000352c60, 0x14000348378, ...}, ...)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:2054 +0xd0
github.com/open-policy-agent/opa/topdown.(*eval).biunifyRef(0x14000146c00, 0x105e39ca0?, 0x14000348378, 0x14000352c60, 0x14000352c60, 0x1052ec97c?)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:1041 +0x338
github.com/open-policy-agent/opa/topdown.(*eval).biunifyValues(0x14000146c00, 0x14000348360, 0x14000348378, 0x14000352c60, 0x14000352c60, 0x14000352d80)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:957 +0xc4
github.com/open-policy-agent/opa/topdown.(*eval).biunify(0x14000146c00, 0x1056ce0cc?, 0x28?, 0x105d6aa00?, 0x140000d4d01?, 0x104e45bbc?)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:853 +0x5ec
github.com/open-policy-agent/opa/topdown.(*eval).unify(0x140000d4d48?, 0x104e45bbc?, 0x140000d4d58?, 0x104e45bbc?)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:842 +0x28
github.com/open-policy-agent/opa/topdown.(*eval).evalStep(0x14000146c00, 0x14000318ee0)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:360 +0x61c
github.com/open-policy-agent/opa/topdown.(*eval).evalExpr(0x14000146c00, 0x14000318ed0)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:341 +0xe4
github.com/open-policy-agent/opa/topdown.(*eval).eval(0x140000102a0?, 0x10595778d?)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:304 +0x1c
github.com/open-policy-agent/opa/topdown.(*eval).Run(0x14000146c00, 0x14000318ec0)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/eval.go:105 +0xb0
github.com/open-policy-agent/opa/topdown.(*Query).Iter(0x140000d5150, {0x105e37cb8?, 0x14000596750}, 0x140003fcb60)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/topdown/query.go:524 +0x9a8
github.com/open-policy-agent/opa/rego.(*Rego).eval(0x1400002db00, {0x105e37cb8?, 0x14000596750}, 0x140000bcd80)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/rego/rego.go:2007 +0x5fc
github.com/open-policy-agent/opa/rego.PreparedEvalQuery.Eval({{0x1400002db00?, 0x140003188d0?}}, {0x105e37cb8, 0x14000596750}, {0x140000d5540?, 0x140003fc760?, 0x1400057d780?})
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa@v0.48.0/rego/rego.go:411 +0x104
github.com/open-policy-agent/opa-envoy-plugin/envoyauth.Eval({0x105e37cb8, 0x14000596750}, {0x105e3e300, 0x14000578fd0}, {0x105e394a0?, 0x140001cfb80}, 0x140006ef0a0, {0x0, 0x0, 0x0})
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa-envoy-plugin@v0.48.0-envoy/envoyauth/evaluation.go:78 +0x628
github.com/open-policy-agent/opa-envoy-plugin/internal.(*envoyExtAuthzGrpcServer).check(0x14000578fd0, {0x105e37cb8?, 0x14000596750}, {0x105d70fc0, 0x14000596780})
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa-envoy-plugin@v0.48.0-envoy/internal/internal.go:351 +0x54c
github.com/open-policy-agent/opa-envoy-plugin/internal.(*envoyExtAuthzGrpcServer).Check(0x105d70fc0?, {0x105e37cb8?, 0x14000596750?}, 0x140001d0c60?)
	/Users/jakub/go/pkg/mod/github.com/open-policy-agent/opa-envoy-plugin@v0.48.0-envoy/internal/internal.go:291 +0x30
github.com/envoyproxy/go-control-plane/envoy/service/auth/v3._Authorization_Check_Handler({0x105de1d40?, 0x14000578fd0}, {0x105e37cb8, 0x14000596750}, 0x140006ee0e0, 0x0)
	/Users/jakub/go/pkg/mod/github.com/envoyproxy/go-control-plane@v0.10.2-0.20220325020618-49ff273808a1/envoy/service/auth/v3/external_auth.pb.go:692 +0x170
google.golang.org/grpc.(*Server).processUnaryRPC(0x1400016c1e0, {0x105e3dbe0, 0x14000182820}, 0x140001165a0, 0x14000172ed0, 0x106759e10, 0x0)
	/Users/jakub/go/pkg/mod/google.golang.org/grpc@v1.51.0/server.go:1340 +0xb7c
google.golang.org/grpc.(*Server).handleStream(0x1400016c1e0, {0x105e3dbe0, 0x14000182820}, 0x140001165a0, 0x0)
	/Users/jakub/go/pkg/mod/google.golang.org/grpc@v1.51.0/server.go:1713 +0x82c
google.golang.org/grpc.(*Server).serveStreams.func1.2()
	/Users/jakub/go/pkg/mod/google.golang.org/grpc@v1.51.0/server.go:965 +0x84
created by google.golang.org/grpc.(*Server).serveStreams.func1
	/Users/jakub/go/pkg/mod/google.golang.org/grpc@v1.51.0/server.go:963 +0x290
exit status 2
```
