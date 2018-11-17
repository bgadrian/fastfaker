# Makefile

pre:
	go get -t -v ./...
	go test -race ./...
	#TODO test for compilation errors examples/*/*

doc: pre
	godoc -http=:6060

bench: pre
	go test -benchmem -bench . github.com/bgadrian/fastfaker/faker github.com/bgadrian/fastfaker/data github.com/bgadrian/fastfaker/randomizer > benchmarks

