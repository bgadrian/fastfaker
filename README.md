# Fast Faker [![Go Report Card](https://goreportcard.com/badge/github.com/bgadrian/fastfaker)](https://goreportcard.com/report/github.com/bgadrian/fastfaker) [![GoDoc](https://godoc.org/github.com/bgadrian/fastfaker?status.svg)](https://godoc.org/github.com/bgadrian/fastfaker) [![Build Status](https://travis-ci.com/bgadrian/fastfaker.svg?branch=master)](https://travis-ci.com/bgadrian/fastfaker) [![codecov](https://codecov.io/gh/bgadrian/fastfaker/branch/master/graph/badge.svg)](https://codecov.io/gh/bgadrian/fastfaker)

FastFaker is a data generator written in go. It can generate over 100 data types and has 2 modes of operation: fast or (concurrent) safe.

### Library Features
- Every function has an example and a benchmark,
[see benchmarks](https://github.com/bgadrian/fastfaker/blob/master/BENCHMARKS.md)
- Zero dependencies (no vendor, dep or modules required)
- Randomizes user defined structs
- Numerous functions for regular use
- Extensible
- Concurrent safe
- Go 1.x compatibility
- Performance
- Idiomatic Go

### 130+ Functions!!!
If there is something that is generic enough missing from this package [let me know!](./CONTRIBUTING.md)

### As a service
If you do not want to spend time implementing this library you can use the [Pseudoservice](https://github.com/bgadrian/pseudoservice) microservice. You can build it yourself or use the docker image
```bash
docker run --name pseudoservice -p 8080:8080 -d -e APIKEY=MYSECRET bgadrian/pseudoservice
curl "http://localhost:8080/users/1?token=MYSECRET&seed=42"
```

## Documentation
All the public functions have godocs and Examples: https://godoc.org/github.com/bgadrian/fastfaker

## Example
```go
package main
import "fmt"
import "github.com/bgadrian/fastfaker"

func main(){
    fastfaker.Global.Name() // Markus Moen
    fastfaker.Global.Email() // alaynawuckert@kozey.biz
    fastfaker.Global.Phone() // (570)245-7485
    fastfaker.Global.BS() // front-end
    fastfaker.Global.BeerName() // Duvel
    fastfaker.Global.Color() // MediumOrchid
    fastfaker.Global.Company() // Moen, Pagac and Wuckert
    fastfaker.Global.CreditCardNumber() // 4287271570245748
    fastfaker.Global.HackerPhrase() // Connecting the array won't do anything, we need to generate the haptic COM driver!
    fastfaker.Global.JobTitle() // Director
    fastfaker.Global.Password(true, true, true, true, true, 32) // WV10MzLxq2DX79w1omH97_0ga59j8!kj
    fastfaker.Global.CurrencyShort() // USD
    // 100+ more!!!
    
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
    fastfaker.Global.Seed(42)
    fastfaker.Global.Struct(&f)
    fmt.Printf("%s\n", f.Browser) //firefox
    fmt.Printf("%s\n", f.Name) //Samuel Smith’s Oatmeal Stout
    fmt.Printf("%d\n", f.Int) //-3651589698752897048
    fmt.Printf("%d\n", f.Dice) //62
    fmt.Printf("%d\n", *f.Pointer) //-8819218091111228151
    fmt.Printf("%v\n", f.Skip) //<nil>
}
```

## Faker types, concurrency and performance
Fast Faker has 2 modes of operating: safe or fast. 
    
##### Safe
`fastfaker.Global` and instances of `NewSafeFaker()` are safe to be used from multiple goroutines, but has a small overhead. Similar with [rand.* global functions](https://golang.org/src/math/rand/rand.go?#L288) it uses a mutex.

> Advantage: easy to use, safe.

##### Fast
Another way of using `fastfaker` is having one `Faker` instance for each goroutine. 
```go
go func(){
    myOwnFaker := fastfaker.NewFastFaker()
    
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
## Thanks
This library started as a fork of [gofakeit](https://github.com/brianvoe/gofakeit/), but I had different requirements from such a library, in particular performance and extensibility.

Because it has braking changes and it was published under a new name/repo the version started from 1 (fastfaker v1 is gofakeit v3.x)


## Benchmarks
For a quick overview see [BENCHMARKS.md](./BENCHMARKS.md).

## Contributing
For more info see the [contributing readme](./CONTRIBUTING.md)

## TODO
See [issues](https://github.com/bgadrian/fastfaker/issues)

