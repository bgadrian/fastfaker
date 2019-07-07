package faker

import (
	"fmt"
	"sync"
	"testing"
)

func Example() {
	Global.Seed(11)
	fmt.Println("Name:", Global.Name())
	fmt.Println("Email:", Global.Email())
	fmt.Println("Phone:", Global.Phone())
	fmt.Println("Address:", Global.Address().Address)
	fmt.Println("BS:", Global.BS())
	fmt.Println("Beer Name:", Global.BeerName())
	fmt.Println("Color:", Global.Color())
	fmt.Println("Company:", Global.Company())
	fmt.Println("Credit Card:", Global.CreditCardNumber())
	fmt.Println("Hacker Phrase:", Global.HackerPhrase())
	fmt.Println("Job Title:", Global.JobTitle())
	fmt.Println("Password:", Global.Password(true, true, true, true, true, 32))
	currency := Global.Currency()
	fmt.Printf("Currency: %s - %s", currency.Short, currency.Long)
	// Output:
	// Name: Markus Moen
	//Email: alaynawuckert@kozey.biz
	//Phone: 7625013191
	//Address: 8995 Wellport, North Denesik, Indiana 30042
	//BS: networks
	//Beer Name: Arrogant Bastard Ale
	//Color: Black
	//Company: Pagac, Ratke and Stracke
	//Credit Card: 6502027613217148
	//Hacker Phrase: Use the open-source JSON system, then you can bypass the optical microchip!
	//Job Title: Strategist
	//Password: SvH*aQzqA *snbo2?z-O=zJn1L# iu u
	//Currency: ZAR - South Africa Rand
}

func TestSeed(t *testing.T) {
	Global.Seed(0)
}

func TestNewSafeFaker_Deadlocks(t *testing.T) {
	wg := sync.WaitGroup{}
	steps := 8000
	threads := 4
	safe := NewSafeFaker()
	wg.Add(threads)

	for t := 0; t < threads; t++ {
		go func() {
			for i := 0; i < steps; i++ {
				safe.Float64Range(-33.2, 10034.23)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func ExampleNewFastFaker() {
	//create only one instance per goroutine!
	faker := NewFastFaker()
	faker.Seed(42)

	fmt.Println("Name:", faker.Name())
	fmt.Println("Email:", faker.Email())
	// Output:
	// Name: Jeromy Schmeler
	// Email: kimsteuber@jones.com
}

func BenchmarkNewSafeFaker_Parallel(b *testing.B) {
	safe := NewSafeFaker()

	b.RunParallel(func(pb *testing.PB) {
		//each thread has its own slice
		arr := []int{5, 6, 7, 8, 9, 0, 2, 3, 4, 5, 6, 7}

		for pb.Next() {
			//we choose this func because it does the most calls to the rand.Int* functions
			safe.ShuffleInts(arr)
		}
	})
}

func BenchmarkNewFastFaker_Parallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		//each thread has its own slice
		arr := []int{5, 6, 7, 8, 9, 0, 2, 3, 4, 5, 6, 7}
		//each thread has its own faker
		fast := NewFastFaker()

		for pb.Next() {
			//we choose this func because it does the most calls to the rand.Int* functions
			fast.ShuffleInts(arr)
		}
	})
}
