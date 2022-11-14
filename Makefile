MAKEFILE_DIR = $(dir $(lastword $(MAKEFILE_LIST)))

main:
	go run main.go

benchmark:
	go test -bench=. -benchmem

pprof:
	go tool pprof