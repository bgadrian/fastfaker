# Fast Faker [![Go Report Card](https://goreportcard.com/badge/github.com/bgadrian/fastfaker)](https://goreportcard.com/report/github.com/bgadrian/fastfaker) [![GoDoc](https://godoc.org/github.com/bgadrian/fastfaker?status.svg)](https://godoc.org/github.com/bgadrian/fastfaker/faker) [![Build Status](https://travis-ci.com/bgadrian/fastfaker.svg?branch=master)](https://travis-ci.com/bgadrian/fastfaker)

FastFaker is a data generator written in go. It can generate over 50 data types and has 2 modes of operation: fast or (concurrent) safe.

### Library Features
- Every function has an example and a benchmark,
[see benchmarks](https://github.com/bgadrian/fastfaker/blob/master/BENCHMARKS.md)
- Zero dependencies (no external calls, no vendor, dep or modules required)
- Randomizes user defined structs
- Over 130 functions for regular use
- Extensible
- Concurrent safe
- Go 1.x compatibility
- Performance
- Idiomatic Go
- predefined popular structures (Address, Person ...)

### 130+ Functions!!!
If there is something that is generic enough missing from this package [let me know!](./CONTRIBUTING.md)

### As a service
If you do not want to spend time implementing this library you can use the [Pseudoservice](https://github.com/bgadrian/pseudoservice) microservice. You can build it yourself or use the docker image
```bash
docker run --name pseudoservice -p 8080:8080 -d -e APIKEY=MYSECRET bgadrian/pseudoservice
curl "http://localhost:8080/users/1?token=MYSECRET&seed=42"
```

## Documentation
All the public functions have comments, examples, tests and benchmarks: https://godoc.org/github.com/bgadrian/fastfaker/faker

You can find more examples in the [Example](./example/) folder.

## Example
```go
package main

import (
	"fmt"

	"github.com/bgadrian/fastfaker/faker"
)

func main() {
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
		Browser string `fake:"{internet.browser}"`
		Name    string `fake:"{beer.name}"`
		Int     int
		Dice    uint8
		Pointer *int
		Skip    *string `fake:"skip"`
	}
	var f Foo
	faker.Global.Seed(42)
	faker.Global.Struct(&f)
	fmt.Printf("%s\n", f.Browser)  //firefox
	fmt.Printf("%s\n", f.Name)     //Samuel Smith’s Oatmeal Stout
	fmt.Printf("%d\n", f.Int)      //-3651589698752897048
	fmt.Printf("%d\n", f.Dice)     //62
	fmt.Printf("%d\n", *f.Pointer) //-8819218091111228151
	fmt.Printf("%v\n", f.Skip)     //<nil>
}

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
BenchmarkNewSafeFaker_Parallel-4           50000             30893 ns/op               0 B/op          0 allocs/op
BenchmarkNewFastFaker_Parallel-4          300000              3829 ns/op               0 B/op          0 allocs/op
```

### /Data [![godoc](https://godoc.org/github.com/bgadrian/fastfaker?status.svg)](https://godoc.org/github.com/bgadrian/fastfaker/data)
If you want to use the raw data that the Faker uses internally (like Names, Streets, Countries and Companies) you can import the [Data package](./data) directly,see its documentation.

### /Randomizer [![godoc](https://godoc.org/github.com/bgadrian/fastfaker?status.svg)](https://godoc.org/github.com/bgadrian/fastfaker/randomizer)
The pseudo-data generator has its [own package](./randomizer) so it can be easily replaced. At its core it is a `rand.*` wrapper with more functions. 

## gofakeit
This library started as a fork of [gofakeit](https://github.com/brianvoe/gofakeit/), but I had different requirements from such a library, in particular performance and extensibility and could not guarantee [backward compatibility](https://github.com/brianvoe/gofakeit/issues/32). Future sync **will** be performed between the projects.

Differences between `gofakeit` and `fastfaker`
* import path, name
* version `gofakeit` 3.x is `fastfaker` 1.x
* different documentation, new examples and tests
* non deterministic (using the same `Seed` on `fastfaker` may lead to different results, eg Name(), than the same seed with `gofakeit`)
* usage, instead of `gofakeit.Name()` the calls are `faker.Global.Name()` or `faker.NewFastFaker().Name()`
* versioning, `fastfaker` uses the semantic version, making it compatible with go modules
* `fastfaker` generates Unicode strings (multi-byte runes)
* `fastfaker` may return non-english data and non-US addresses

## Benchmarks
For a quick overview see [BENCHMARKS.md](./BENCHMARKS.md).

## Contributing
For more info see the [contributing readme](./CONTRIBUTING.md)

## TODO
See [issues](https://github.com/bgadrian/fastfaker/issues)

