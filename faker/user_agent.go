package faker

import "strconv"

// Browser will generate a random browser name
func (f *Faker) Browser() string {
	return f.getRandValue([]string{"internet", "browser"})
}

// UserAgent will generate a random browser user agent
func (f *Faker) UserAgent() string {
	randNum := f.Number(0, 4)
	switch randNum {
	case 0:
		return f.ChromeUserAgent()
	case 1:
		return f.FirefoxUserAgent()
	case 2:
		return f.SafariUserAgent()
	case 3:
		return f.OperaUserAgent()
	default:
		return f.ChromeUserAgent()
	}
}

// ChromeUserAgent will generate a random chrome browser user agent string
func (f *Faker) ChromeUserAgent() string {
	randNum1 := strconv.Itoa(f.Number(531, 536)) + strconv.Itoa(f.Number(0, 2))
	randNum2 := strconv.Itoa(f.Number(36, 40))
	randNum3 := strconv.Itoa(f.Number(800, 899))
	return "Mozilla/5.0 " + "(" + f.randomPlatform() + ") AppleWebKit/" + randNum1 + " (KHTML, like Gecko) Chrome/" + randNum2 + ".0." + randNum3 + ".0 Mobile Safari/" + randNum1
}

// FirefoxUserAgent will generate a random firefox broswer user agent string
func (f *Faker) FirefoxUserAgent() string {
	ver := "Gecko/" + f.Date().Format("2006-02-01") + " Firefox/" + strconv.Itoa(f.Number(35, 37)) + ".0"
	platforms := []string{
		"(" + f.windowsPlatformToken() + "; " + "en-US" + "; rv:1.9." + strconv.Itoa(f.Number(0, 3)) + ".20) " + ver,
		"(" + f.linuxPlatformToken() + "; rv:" + strconv.Itoa(f.Number(5, 8)) + ".0) " + ver,
		"(" + f.macPlatformToken() + " rv:" + strconv.Itoa(f.Number(2, 7)) + ".0) " + ver,
	}

	return "Mozilla/5.0 " + f.RandString(platforms)
}

// SafariUserAgent will generate a random safari browser user agent string
func (f *Faker) SafariUserAgent() string {
	randNum := strconv.Itoa(f.Number(531, 536)) + "." + strconv.Itoa(f.Number(1, 51)) + "." + strconv.Itoa(f.Number(1, 8))
	ver := strconv.Itoa(f.Number(4, 6)) + "." + strconv.Itoa(f.Number(0, 2))

	mobileDevices := []string{
		"iPhone; CPU iPhone OS",
		"iPad; CPU OS",
	}

	platforms := []string{
		"(Windows; U; " + f.windowsPlatformToken() + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + f.macPlatformToken() + " rv:" + strconv.Itoa(f.Number(4, 7)) + ".0; en-US) AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + f.RandString(mobileDevices) + " " + strconv.Itoa(f.Number(7, 9)) + "_" + strconv.Itoa(f.Number(0, 3)) + "_" + strconv.Itoa(f.Number(1, 3)) + " like Mac OS X; " + "en-US" + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + strconv.Itoa(f.Number(3, 5)) + ".0.5 Mobile/8B" + strconv.Itoa(f.Number(111, 120)) + " Safari/6" + randNum,
	}

	return "Mozilla/5.0 " + f.RandString(platforms)
}

// OperaUserAgent will generate a random opera browser user agent string
func (f *Faker) OperaUserAgent() string {
	platform := "(" + f.randomPlatform() + "; en-US) Presto/2." + strconv.Itoa(f.Number(8, 13)) + "." + strconv.Itoa(f.Number(160, 355)) + " Version/" + strconv.Itoa(f.Number(10, 13)) + ".00"

	return "Opera/" + strconv.Itoa(f.Number(8, 10)) + "." + strconv.Itoa(f.Number(10, 99)) + " " + platform
}

// linuxPlatformToken will generate a random linux platform
func (f *Faker) linuxPlatformToken() string {
	return "X11; Linux " + f.getRandValue([]string{"computer", "linux_processor"})
}

// macPlatformToken will generate a random mac platform
func (f *Faker) macPlatformToken() string {
	return "Macintosh; " + f.getRandValue([]string{"computer", "mac_processor"}) + " Mac OS X 10_" + strconv.Itoa(f.Number(5, 9)) + "_" + strconv.Itoa(f.Number(0, 10))
}

// windowsPlatformToken will generate a random windows platform
func (f *Faker) windowsPlatformToken() string {
	return f.getRandValue([]string{"computer", "windows_platform"})
}

// randomPlatform will generate a random platform
func (f *Faker) randomPlatform() string {
	platforms := []string{
		f.linuxPlatformToken(),
		f.macPlatformToken(),
		f.windowsPlatformToken(),
	}

	return f.RandString(platforms)
}
