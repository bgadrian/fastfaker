# Fast Faker [![Go Report Card](https://goreportcard.com/badge/github.com/bgadrian/fastfaker)](https://goreportcard.com/report/github.com/bgadrian/fastfaker) [![GoDoc](https://godoc.org/github.com/bgadrian/fastfaker?status.svg)](https://godoc.org/github.com/bgadrian/fastfaker/faker) [![Build Status](https://travis-ci.com/bgadrian/fastfaker.svg?branch=master)](https://travis-ci.com/bgadrian/fastfaker) [![codecov](https://codecov.io/gh/bgadrian/fastfaker/branch/master/graph/badge.svg)](https://codecov.io/gh/bgadrian/fastfaker) [![contributions](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/bgadrian/fastfaker/issues)

FastFaker is a data generator written in go. It can generate over 50 data types and has 2 modes of operation: fast or (concurrent) safe.

### Features
- Every function has an example and a benchmark,
[see benchmarks](https://github.com/bgadrian/fastfaker/blob/master/BENCHMARKS.md)
- Zero dependencies (no external calls, no vendor, dep or modules required)
- Randomizes user defined structs
- Over 130 functions for regular use
- Templates with over 110 variables
- Extensible
- Concurrent safe
- Go 1.x compatibility
- Performance
- Idiomatic Go
- predefined popular structures (Address, Person ...)
- 200 unit tests and examples

### 130+ Functions!!!
If there is something that is generic enough missing from this package [let me know!](./CONTRIBUTING.md)

## How to use
#### Library 
The package is meant to be used as a Go package, and imported in your own application, but it can be used as an app or service too:

#### Utility
A small example is provided as an example, and can be used as a binary standalone utility, after compilation:
```bash
$ go build -o fastfaker example/standalone-template/main.go
$ ./fastfaker '{name}'
$ Victoria Berge
```
#### Web service
If you do not want to spend time importing this library in your project, or you are using a different programming language, you can use the [Pseudoservice](https://github.com/bgadrian/pseudoservice) microservice. It is a HTTP Web sever wrapper on the Templates functionality.

```bash
docker run --name pseudoservice -p 8080:8080 -d -e APIKEY=MYSECRET bgadrian/pseudoservice
curl "http://localhost:8080/custom/1?token=MYSECRET&template=~name~"
{"results":["Jefferey Koch"],"seed":-4329827746951412836}
```

## Documentation
All the public functions have comments, examples, tests and benchmarks: https://godoc.org/github.com/bgadrian/fastfaker/faker

You can find more examples in the [Example](./example/) folder.

## Example
```go
	faker.Global.Name()             // Markus Moen
	faker.Global.Email()            // alaynawuckert@kozey.biz
	faker.Global.Phone()            // (570)245-7485
	faker.Global.BS()               // front-end
	faker.Global.BeerName()         // Duvel
	faker.Global.Color()            // MediumOrchid
	faker.Global.Company()          // Moen, Pagac and Wuckert
	faker.Global.CreditCardNumber() // 4287271570245748
	faker.Global.HackerPhrase()     // Connecting the array won't do anything, we need to generate the haptic COM driver!
	faker.Global.JobTitle()         // Director
	faker.Global.PasswordFull()     // +eHQa02X9n3X
	faker.Global.CurrencyShort()    // USD
	// 130+ more!!!

	// Create structs with random injected data
	type Foo struct {
		Browser string `fake:"{browser}"`
		Name    string `fake:"{beername}"`
		Int     int
		Dice    uint8
		Pointer *int
		Skip    *string `fake:"skip"`
	}
	var f Foo
	faker.Global.Struct(&f)
	fmt.Printf("%s\n", f.Browser)  //firefox
	fmt.Printf("%s\n", f.Name)     //Samuel Smith’s Oatmeal Stout
	fmt.Printf("%d\n", f.Int)      //-3651589698752897048
	fmt.Printf("%d\n", f.Dice)     //62
	fmt.Printf("%d\n", *f.Pointer) //-8819218091111228151
	fmt.Printf("%v\n", f.Skip)     //<nil>
```

## Faker types, concurrency and performance
Fast Faker has 2 modes of operating: safe or fast. 
    
##### Safe
`faker.Global` and instances of `NewSafeFaker()` are safe to be used from multiple goroutines, but has a small overhead. Similar with [rand.* global functions](https://golang.org/src/math/rand/rand.go?#L288) it uses a mutex.

> Advantage: easy to use, safe.

##### Fast
Another way of using `faker` is having one `Faker` instance for each goroutine. 
```go
go func(){
    myOwnFaker := faker.NewFastFaker()
    
    myOwnFaker.Name() // Markus Moen
    myOwnFaker.Email() // alaynawuckert@kozey.biz
}()
```

> Advantage: performance

For a heavy usage in a 4 CPU env, the performance benefits of the FastFaker can be up to 10x:
```
go test -benchmem  -bench=BenchmarkNew* github.com/bgadrian/fastfaker/faker
goos: linux
goarch: amd64
pkg: github.com/bgadrian/fastfaker/faker
BenchmarkNewSafeFaker_Parallel-4          500000              2668 ns/op               0 B/op          0 allocs/op
BenchmarkNewFastFaker_Parallel-4        10000000               225 ns/op               0 B/op          0 allocs/op
```

## Templates
Template is the most powerful FastFaker feature. It allows custom patterns/templates (JSON, YAML, HTML, ...) of text to be filled with over 110 random types of [data (variables)](./TEMPLATE_VARIABLES.md).

It can be used directly (faker.Template* methods) or via the faker.Struct fill method and `fake:` tags. 

```go
//instead of:
fmt.Sprintf("I'm %s, call me at %s!", fastFaker.Name(), fastFaker.Numerify("###-###-####"))
//you can do:
fastFaker.Template("I'm {name}, call me at ###-###-####!") // I'm John, call me at 152-335-8761!
```

You can use any type of texts, formats, encoded or not (HTML, JSON, YAML, configs ...). For more details see the [TEMPLATES section](./TEMPLATES.md)

### /Data [![godoc](https://godoc.org/github.com/bgadrian/fastfaker?status.svg)](https://godoc.org/github.com/bgadrian/fastfaker/data)
If you want to use the raw data that the Faker uses internally (like Names, Streets, Countries and Companies) you can import the [Data package](./data) directly, see its documentation.

### /Randomizer [![godoc](https://godoc.org/github.com/bgadrian/fastfaker?status.svg)](https://godoc.org/github.com/bgadrian/fastfaker/randomizer)
The pseudo-data generator has its [own package](./randomizer). At its core it is a `rand.*` wrapper with more functions. 

## gofakeit
This library started as a fork of [gofakeit](https://github.com/brianvoe/gofakeit/), but I had different requirements from such a library, in particular performance and extensibility and could not guarantee [backward compatibility](https://github.com/brianvoe/gofakeit/issues/32). Future sync **will** be performed between the projects.

How is FastFaker different?
* performance (up to 10x improvement)
* code quality (fixed over 50 gometalinter warnings)
* Templates
* error logger (stdErr) instead of silent fails
* new functions
* more documentation, new examples and tests
* usage, instead of `gofakeit.Name()` the calls are `faker.Global.Name()` or `faker.NewFastFaker().Name()`
* `fastfaker` uses the semantic version, making it compatible with go modules
* split /data and /randomizer into their own packages
* version `gofakeit` 3.x is `fastfaker` 1.x
* version `gofakeit` features are brought in `fastfaker` 2.1
* more in the [changelog](./CHANGELOG.md)

## Change log
All the [major changes are found in the CHANGELOG](./CHANGELOG.md).

## Benchmarks
For a quick overview see [BENCHMARKS.md](./BENCHMARKS.md).

## Contributing
For more info see the [contributing readme](./CONTRIBUTING.md)

## TODO
See [issues](https://github.com/bgadrian/fastfaker/issues)

