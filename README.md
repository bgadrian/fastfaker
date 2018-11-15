![alt text](https://raw.githubusercontent.com/brianvoe/gofakeit/master/logo.png)

# gofakeit [![Go Report Card](https://goreportcard.com/badge/github.com/brianvoe/gofakeit)](https://goreportcard.com/report/github.com/brianvoe/gofakeit) [![Build Status](https://travis-ci.org/brianvoe/gofakeit.svg?branch=master)](https://travis-ci.org/brianvoe/gofakeit) [![codecov.io](https://codecov.io/github/brianvoe/gofakeit/branch/master/graph/badge.svg)](https://codecov.io/github/brianvoe/gofakeit) [![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit?status.svg)](https://godoc.org/github.com/brianvoe/gofakeit) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/brianvoe/gofakeit/master/LICENSE.txt)
Random data generator written in go

<a href="https://www.buymeacoffee.com/brianvoe" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

### Features
- Every function has an example and a benchmark,
[see benchmarks](https://github.com/brianvoe/gofakeit/blob/master/BENCHMARKS.md)
- Zero dependencies
- Randomizes user defined structs
- Numerous functions for regular use
- Extensible

### 80+ Functions!!!
If there is something that is generic enough missing from this package [add an issue](https://github.com/brianvoe/gofakeit/issues) and let me know what you need.
Most of the time i'll add it!

## Documentation
[![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit?status.svg)](https://godoc.org/github.com/brianvoe/gofakeit)

## Example
```go
package main
import "fmt"
import "github.com/brianvoe/gofakeit"

func main(){
    gofakeit.Global.Name() // Markus Moen
    gofakeit.Global.Email() // alaynawuckert@kozey.biz
    gofakeit.Global.Phone() // (570)245-7485
    gofakeit.Global.BS() // front-end
    gofakeit.Global.BeerName() // Duvel
    gofakeit.Global.Color() // MediumOrchid
    gofakeit.Global.Company() // Moen, Pagac and Wuckert
    gofakeit.Global.CreditCardNumber() // 4287271570245748
    gofakeit.Global.HackerPhrase() // Connecting the array won't do anything, we need to generate the haptic COM driver!
    gofakeit.Global.JobTitle() // Director
    gofakeit.Global.Password(true, true, true, true, true, 32) // WV10MzLxq2DX79w1omH97_0ga59j8!kj
    gofakeit.Global.CurrencyShort() // USD
    // 80+ more!!!
    
    // Create structs with random injected data
    type Foo struct {
        Bar     string
        Baz     string
        Int     int
        Pointer *int
        Skip    *string `fake:"skip"` // Set to "skip" to not generate data for
    }
    var f Foo
    gofakeit.Global.Struct(&f)
    fmt.Printf("f.Bar:%s\n", f.Bar) // f.Bar:hrukpttuezptneuvunh
    fmt.Printf("f.Baz:%s\n", f.Baz) // f.Baz:uksqvgzadxlgghejkmv
    fmt.Printf("f.Int:%d\n", f.Int) // f.Int:-7825289004089916589
    fmt.Printf("f.Pointer:%d\n", *f.Pointer) // f.Pointer:-343806609094473732
    fmt.Printf("f.Skip:%v\n", f.Skip) // f.Skip:<nil>
}
```

## Faker types and concurrency

`gofakeit.Global` and instances of `NewSafeFaker()` are safe to be used from multiple goroutines, but has a small overhead. Similar with [rand.* global functions](https://golang.org/src/math/rand/rand.go?#L288) it uses a mutex.

> Advantage: easy to use, safe.

Another way of using `gofakeit` is having one `Faker` instance for each goroutine. 
```go
go func(){
    myOwnFaker := gofakeit.NewFastFaker()
    
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

## Benchmarks
For a quick overview see [BENCHMARKS.md](./BENCHMARKS.md).
