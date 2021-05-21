go run ./cmd/go-printf-func-name/main.go -- ./testdata/src/p/p.go

go build -ldflags="-s -w"  -buildmode=plugin  plugin/go-lint-nill-err.go
golangci-lint run -Ego-lint-nil-err --timeout 100m  --disable-all  pkg/nonce/service.go
# go-printf-func-name

The Go linter `go-printf-func-name` checks that printf-like functions are named with `f` at the end.

For example, `myLog` should be named `myLogf` by Go convention:

```go
package main

import "log"

func myLog(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix + format, args...)
}
```
