go run ./cmd/go-printf-func-name/main.go -- ./testdata/src/p/p.go

go build -ldflags="-s -w"  -buildmode=plugin  plugin/go-lint-nill-err.go
golangci-lint run -Ego-lint-nil-err --timeout 100m  --disable-all  pkg/nonce/service.go

