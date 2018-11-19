package main

import (
	"fmt"
	"sync"

	"github.com/bgadrian/fastfaker/faker"
)

func main() {
	/**
		To use the Faker in a concurrent app you have 2 choices:
	* Use a SafeFaker that is an instance shared by more goroutines.
	* A FastFaker will be more efficient but you cannot share it between
		goroutines

	*/

	workers := 5
	count := 4

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(id int) {
			myFaker := faker.NewFastFaker()
			for r := 0; r < count; r++ {
				randomLoto := myFaker.Uint8()
				fmt.Printf("routine %d: %d\n", id, randomLoto)
			}
			wg.Done()
		}(i + 1)
	}

	wg.Wait()
}
