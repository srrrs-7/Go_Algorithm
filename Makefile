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