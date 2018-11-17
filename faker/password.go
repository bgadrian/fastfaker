package faker

const lowerStr = "abcdefghijklmnopqrstuvwxyz"
const upperStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numericStr = "0123456789"
const specialStr = "!@#$%&*+-=?"
const spaceStr = "   "

// PasswordFull will generate a random password of length 12 that contains at least
// 1 digit, 1 special character, 1 Lower and 1 Upper case letters.
func (f *Faker) PasswordFull() string {
	return f.Password(true, true, true, true, false, 12)
}

// Password will generate a random password
// Minimum number length of 5 if less than
func (f *Faker) Password(lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	// Make sure the num minimum is at least 5
	if num < 5 {
		num = 5
	}
	i := 0
	b := make([]byte, num)
	var passString string

	//it is NOT unicode safe! only works for ASCII
	randomByte := func(source string) byte {
		return source[f.Int64Positive()%int64(len(source))]
	}

	if lower {
		passString += lowerStr
		b[i] = randomByte(lowerStr)
		i++
	}
	if upper {
		passString += upperStr
		b[i] = randomByte(upperStr)
		i++
	}
	if numeric {
		passString += numericStr
		b[i] = randomByte(numericStr)
		i++
	}
	if special {
		passString += specialStr
		b[i] = randomByte(specialStr)
		i++
	}
	if space {
		passString += spaceStr
		b[i] = randomByte(spaceStr)
		i++
	}

	// Set default if empty
	if passString == "" {
		passString = lowerStr + numericStr
	}

	// Loop through and add it up
	for i <= num-1 {
		b[i] = randomByte(passString)
		i++
	}

	// Shuffle bytes
	for i := range b {
		j := f.Intn(i + 1)
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
