/*
Package faker contains a random data generator (faker) with over 130 functions and has 2 modes: fast or (concurrent) safe.

Usage:
	import "github.com/bgadrian/fastfaker/faker"

Method 1, use the Global instance (simple and concurrent safe):
	fmt.Println("Name:", Global.Name()) //Name: Markus Moen

Method 2, instantiate a concurrent safe version:
	faker := NewSafeFaker()
	fmt.Println("Email:", faker.Email()) //kimsteuber@jones.com
Method 3, create a fast faker that is Not safe to use from multiple goroutines:
	faker := NewFastFaker()
	fmt.Println(Global.BeerName()) //Duvel

Fakers take their data from the github.com/bgadrian/fastfaker/data package.
The pseudo-random algorithm is provided by the github.com/bgadrian/fastfaker/randomizer
package, that is a wrapper of the rand package from the standard library.

*/
package faker
