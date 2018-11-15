package gofakeit

import "strconv"

// SSN will generate a random Social Security Number
func (f *Faker) SSN() string {
	return strconv.Itoa(f.randIntRange(100000000, 999999999))
}

// Gender will generate a random gender string "male" or "female"
func (f *Faker) Gender() string {
	if f.Bool() == true {
		return "male"
	}

	return "female"
}

// PersonInfo is a struct of person information
type PersonInfo struct {
	FirstName  string
	LastName   string
	Gender     string
	SSN        string
	Image      string
	Job        *JobInfo
	Address    *AddressInfo
	Contact    *ContactInfo
	CreditCard *CreditCardInfo
}

// Person will generate a struct with person information
func (f *Faker) Person() *PersonInfo {
	return &PersonInfo{
		FirstName:  f.FirstName(),
		LastName:   f.LastName(),
		Gender:     f.Gender(),
		SSN:        f.SSN(),
		Image:      f.ImageURL(300, 300) + "/people",
		Job:        f.Job(),
		Address:    f.Address(),
		Contact:    f.Contact(),
		CreditCard: f.CreditCard(),
	}
}
