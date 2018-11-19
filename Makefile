# Makefile

pre:
	go get -t -v ./...
	go test -race ./...
	#TODO test for compilation errors examples/*/*

doc: pre
	godoc -http=:6060

bench: pre
	go test -benchmem -bench . github.com/bgadrian/fastfaker/faker github.com/bgadrian/fastfaker/data github.com/bgadrian/fastfaker/randomizer > benchmarks

#run make generate after each new Faker public method change (add/remove/rename)
generate:
	go run cmd/template-variable-generator/main.go go > templates_variables.go
	mv templates_variables.go faker/templates_variables.go
	go run cmd/template-variable-generator/main.go mark > TEMPLATE_VARIABLES.md
