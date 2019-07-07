package faker

import (
	"strconv"
	"strings"
)

// ContactInfo struct full of contact info
type ContactInfo struct {
	Phone string
	Email string
}

// Contact will generate a struct with information randomly populated contact information
func (f *Faker) Contact() *ContactInfo {
	return &ContactInfo{
		Phone: f.Phone(),
		Email: f.Email(),
	}
}

// Phone will generate a 10 digit random phone number string
func (f *Faker) Phone() string {
	return strconv.FormatInt(f.Int64Range(1000000000, 9999999999), 10)
}

// PhoneFormatted will generate a random phone number string
func (f *Faker) PhoneFormatted() string {
	return f.replaceWithNumbers(f.getRandValue([]string{"contact", "phone"}))
}

// Email will generate a random email string
func (f *Faker) Email() string {
	email := f.getRandValue([]string{"person", "first"}) + f.getRandValue([]string{"person", "last"})
	email += "@"
	email += f.getRandValue([]string{"person", "last"}) + "." + f.getRandValue([]string{"internet", "domain_suffix"})

	return strings.ToLower(email)
}
