MAKEFILE_DIR = $(dir $(lastword $(MAKEFILE_LIST)))

run:
	go run main.go

benchmark:
	go test -bench=. -benchmem

pprof:
	go tool pprof

debug:
	dlv debug main.go

dlvhelp:
	dlv help

mecab:
	export CGO_LDFLAGS="`mecab-config --libs`"
	export CGO_CFLAGS="-I`mecab-config --inc-dir`"
	go get github.com/shogo82148/go-mecab