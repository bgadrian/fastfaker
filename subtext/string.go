package subtext

const hashtag = '#'
const questionmark = '?'

// Letter will generate a single random lower case ASCII letter
func (f *FakerText) Letter() string {
	return string(byte(f.textRand.Intn(26)) + 'a')
}

// Digit will generate a single ASCII digit
func (f *FakerText) Digit() string {
	return string(byte(f.textRand.Intn(9)) + '0')
}

// Replace ? with ASCII lowercase letters
func (f *FakerText) Lexify(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == questionmark {
			//TODO modify to allow multi-byte letters
			letter := f.Letter()
			bytestr[i] = []byte(letter)[0]
		}
	}

	return string(bytestr)
}

// Numerify will replace # with random numerical values (0-9 digits)
func (f *FakerText) Numerify(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			//TODO handle multi-byte digits !?
			bytestr[i] = []byte(f.Digit())[0]
		}
	}
	if bytestr[0] == '0' {
		bytestr[0] = byte(f.textRand.Intn(8)+1) + '0'
	}

	return string(bytestr)
}

// ShuffleStrings will randomize a slice of strings
func (f *FakerText) ShuffleStrings(a []string) {
	swap := func(i, j int) {
		a[i], a[j] = a[j], a[i]
	}
	//to avoid upgrading to 1.10 I copied the algorithm
	n := len(a)
	if n <= 1 {
		return
	}

	//if size is > int32 probably it will never finish, or ran out of entropy
	i := n - 1
	for ; i > 0; i-- {
		j := int(f.textRand.Int32Range(0, int32(i+1)))
		swap(i, j)
	}
}

// RandString will take in a slice of string and return a randomly selected value
func (f *FakerText) RandString(a []string) string {
	size := len(a)
	if size == 0 {
		return ""
	}
	return a[f.textRand.Intn(size)]
}
