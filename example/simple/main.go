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
	//for all the variables available see:
	//https://github.com/bgadrian/fastfaker/TEMPLATE_VARIABLES.md
	type Foo struct {
		Browser string `fake:"{browser}"`
		Name    string `fake:"{beername}"`
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
