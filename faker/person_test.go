package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_SSN() {
	Global.Seed(11)
	fmt.Println(Global.SSN())
	// Output: 296446360
}

func BenchmarkSSN(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.SSN()
	}
}

func ExampleFaker_Gender() {
	Global.Seed(11)
	fmt.Println(Global.Gender())
	// Output: female
}

func BenchmarkGender(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Gender()
	}
}

func ExampleFaker_Person() {
	Global.Seed(11)
	person := Global.Person()
	job := person.Job
	address := person.Address
	contact := person.Contact
	creditCard := person.CreditCard

	fmt.Println(person.FirstName)
	fmt.Println(person.LastName)
	fmt.Println(person.Gender)
	fmt.Println(person.SSN)
	fmt.Println(person.Image)

	fmt.Println(job.Company)
	fmt.Println(job.Title)
	fmt.Println(job.Descriptor)
	fmt.Println(job.Level)

	fmt.Println(address.Address)
	fmt.Println(address.Street)
	fmt.Println(address.City)
	fmt.Println(address.State)
	fmt.Println(address.Zip)
	fmt.Println(address.Country)
	fmt.Println(address.Latitude)
	fmt.Println(address.Longitude)

	fmt.Println(contact.Phone)
	fmt.Println(contact.Email)

	fmt.Println(creditCard.Type)
	fmt.Println(creditCard.Number)
	fmt.Println(creditCard.Exp)
	fmt.Println(creditCard.Cvv)

	// Output: Markus
	//Moen
	//male
	//420776036
	//http://pipsum.com/80x80.jpg
	//Lockman and Sons
	//Developer
	//Global
	//Brand
	//5369 Streamville, Rossieview, Hawaii 42591
	//5369 Streamville
	//Rossieview
	//Hawaii
	//42591
	//Burkina Faso
	//-6.662594491850811
	//23.921575244414612
	//9676941592
	//christyratke@stracke.org
	//Visa
	//4027613217148008
	//11/29
	//982
}

func BenchmarkPerson(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Person()
	}
}
