package faker

import "strconv"
import "time"

var currentYear = time.Now().Year() - 2000

// CreditCardInfo is a struct containing credit variables
type CreditCardInfo struct {
	Type   string
	Number int
	Exp    string
	Cvv    string
}

// CreditCard will generate a struct full of credit card information
func (f *Faker) CreditCard() *CreditCardInfo {
	return &CreditCardInfo{
		Type:   f.CreditCardType(),
		Number: f.CreditCardNumber(),
		Exp:    f.CreditCardExp(),
		Cvv:    f.CreditCardCvv(),
	}
}

// CreditCardType will generate a random credit card type string
func (f *Faker) CreditCardType() string {
	return f.getRandValue([]string{"payment", "card_type"})
}

// CreditCardNumber will generate a random credit card number int
func (f *Faker) CreditCardNumber() int {
	integer, err := strconv.Atoi(f.replaceWithNumbers(f.getRandValue([]string{"payment", "number"})))
	if err != nil {
		errorLogger.Printf("(CreditCardNumber) %s\n", err)
		return 0
	}
	return integer
}

// CreditCardNumberLuhn will generate a random credit card number int that passes luhn test
func (f *Faker) CreditCardNumberLuhn() int {
	cc := ""
	for i := 0; i < 100000; i++ {
		cc = f.replaceWithNumbers(f.getRandValue([]string{"payment", "number"}))
		if f.luhn(cc) {
			break
		}
	}
	integer, err := strconv.Atoi(cc)
	if err != nil {
		errorLogger.Printf("(CreditCardNumberLuhn) %s\n", err)
		return 0
	}
	return integer
}

// CreditCardExp will generate a random credit card expiration date string
// Exp date will always be a future date
func (f *Faker) CreditCardExp() string {
	month := strconv.Itoa(f.Number(1, 12))
	if len(month) == 1 {
		month = "0" + month
	}
	return month + "/" + strconv.Itoa(f.Number(currentYear+1, currentYear+10))
}

// CreditCardCvv will generate a random CVV number - Its a string because you could have 017 as an exp date
func (f *Faker) CreditCardCvv() string {
	return f.Numerify("###")
}

// luhn check is used for checking if credit card is valid
func (f *Faker) luhn(s string) bool {
	var t = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	odd := len(s) & 1
	var sum int
	for i, c := range s {
		if c < '0' || c > '9' {
			return false
		}
		if i&1 == odd {
			sum += t[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}
