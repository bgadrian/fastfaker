# Makefile

pre:
	env go test -race ./...
	#TODO test for compilation errors examples/*/*

doc: pre
	godoc -http=:6060
