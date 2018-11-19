/*
Package faker contains a random data generator (faker) with over 130 functions and has 2 modes: fast or (concurrent) safe.
Can be used to mock data, test or stress test your systems or databases.
You can generate almost anything, from names, addresses to prices and vehicle transmission types.

If you are looking for a HTTP service with the same functionality, not a package see Pseudoservice, it is a wrapper on FastFaker https://github.com/bgadrian/pseudoservice

Usage:
	import "github.com/bgadrian/fastfaker/faker"

Method 1, use the Global instance (simple and concurrent safe):
	fmt.Println("Name:", Global.Name()) //Name: Markus Moen

Method 2, instantiate a concurrent safe version:
	faker := NewSafeFaker()
	fmt.Println("Email:", faker.Email()) //kimsteuber@jones.com
Method 3, create a fast faker that is Not safe to use from multiple goroutines:
	faker := NewFastFaker()
	fmt.Println(Global.Template("{name} drinks {beername}.")) //Adrian drinks Duvel.

Fakers take their data from the github.com/bgadrian/fastfaker/data package.
The pseudo-random algorithm is provided by the github.com/bgadrian/fastfaker/randomizer
package, that is a wrapper of the rand package from the standard library.

For a more advanced overview of the Template feature see https://github.com/bgadrian/fastfaker/blob/master/TEMPLATES.md
*/
package faker
