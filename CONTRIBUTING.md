## New data
If you want to add a new data type and you do not have the time for a PR please make a CSV or a Go map and send it to me, I will add it to the project. Example: https://github.com/bgadrian/fastfaker/blob/master/data/colors.go

## New functionality
Add an Issue https://github.com/bgadrian/fastfaker/issues

After we agreed on the implementation you can make a Pull Request with the description:
```
<Functionality description>
Fixes #issueID
```
## Requirements
* Go (preferable the latest version)
* bash-like environment with `make`
* activate the `go fmt` and `go imports` file watchers
* keep the 100% test coverage
* each Public method must examples or tests, and a benchmark
* keep the Go 1.0 compatibility

## Caveats
The templates system is built on top of generated code. Make sure you run 
```bash
$ make generate
```
each time you update a `Faker` public method.