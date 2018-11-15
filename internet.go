package fastfaker

import (
	"fmt"
	"strings"
)

// DomainName will generate a random url domain name
func (f *Faker) DomainName() string {
	return strings.ToLower(f.JobDescriptor()+f.BS()) + "." + f.DomainSuffix()
}

// DomainSuffix will generate a random domain suffix
func (f *Faker) DomainSuffix() string {
	return f.getRandValue([]string{"internet", "domain_suffix"})
}

// URL will generate a random url string
func (f *Faker) URL() string {
	url := "http" + f.RandString([]string{"s", ""}) + "://www."
	url += f.DomainName()

	// Slugs
	num := f.Number(1, 4)
	slug := make([]string, num)
	for i := 0; i < num; i++ {
		slug[i] = f.BS()
	}
	url += "/" + strings.ToLower(strings.Join(slug, "/"))

	return url
}

// HTTPMethod will generate a random http method
func (f *Faker) HTTPMethod() string {
	return f.getRandValue([]string{"internet", "http_method"})
}

// IPv4Address will generate a random version 4 ip address
func (f *Faker) IPv4Address() string {
	num := func() int { return 2 + f.Intn(254) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

// IPv6Address will generate a random version 6 ip address
func (f *Faker) IPv6Address() string {
	num := 65536
	return fmt.Sprintf("2001:cafe:%x:%x:%x:%x:%x:%x", f.Intn(num), f.Intn(num), f.Intn(num), f.Intn(num), f.Intn(num), f.Intn(num))
}

// Username will genrate a random username based upon picking a random lastname and random numbers at the end
func (f *Faker) Username() string {
	return f.getRandValue([]string{"person", "last"}) + f.replaceWithNumbers("####")
}
