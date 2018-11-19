package faker

import (
	"fmt"
	"testing"
	"time"
)

func ExampleFaker_Date() {
	Global.Seed(11)
	fmt.Println(Global.Date().String())
	// Output: 1989-01-07 04:14:25.685339029 +0000 UTC
}
func ExampleFaker_DateStr() {
	Global.Seed(11)
	fmt.Println(Global.DateStr())
	// Output: 1989-01-07 04:14:25.685339029 +0000 UTC
}

func BenchmarkDate(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Date()
	}
}

func ExampleFaker_DateCurrentYear() {
	Global.Seed(11)
	fmt.Println(Global.DateCurrentYear().String())
	// Output: 2018-06-25 05:16:14.244961305 +0000 UTC
}
func ExampleFaker_DateCurrentYearStr() {
	Global.Seed(11)
	fmt.Println(Global.DateCurrentYearStr())
	// Output: 2018-06-25 05:16:14.244961305 +0000 UTC
}

func BenchmarkDateCurrentYear(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.DateCurrentYear()
	}
}

func ExampleFaker_DateRange() {
	Global.Seed(11)
	fmt.Println(Global.DateRange(time.Unix(0, 484633944473634951), time.Unix(0, 1431318744473668209))) // May 10, 1985 years to May 10, 2015
	// Output: 2012-02-04 14:10:37.166933216 +0000 UTC
}

func BenchmarkDateRange(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.DateRange(time.Now().AddDate(-30, 0, 0), time.Now())
	}
}

func ExampleFaker_Month() {
	Global.Seed(11)
	fmt.Println(Global.Month())
	// Output: January
}

func ExampleFaker_WeekDay() {
	Global.Seed(11)
	fmt.Println(Global.WeekDay())
	// Output: Friday
}

func BenchmarkMonth(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Month()
	}
}

func ExampleFaker_Day() {
	Global.Seed(11)
	fmt.Println(Global.Day())
	// Output: 12
}

func BenchmarkDay(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Day()
	}
}

func BenchmarkWeekDay(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.WeekDay()
	}
}

func ExampleFaker_Year() {
	Global.Seed(11)
	fmt.Println(Global.Year())
	// Output: 1989
}

func BenchmarkYear(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Year()
	}
}

func ExampleFaker_Hour() {
	Global.Seed(11)
	fmt.Println(Global.Hour())
	// Output: 0
}

func BenchmarkHour(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Hour()
	}
}

func ExampleFaker_Minute() {
	Global.Seed(11)
	fmt.Println(Global.Minute())
	// Output: 0
}

func BenchmarkMinute(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Minute()
	}
}

func ExampleFaker_Second() {
	Global.Seed(11)
	fmt.Println(Global.Second())
	// Output: 0
}

func BenchmarkSecond(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Second()
	}
}

func ExampleFaker_NanoSecond() {
	Global.Seed(11)
	fmt.Println(Global.NanoSecond())
	// Output: 196446360
}

func BenchmarkNanoSecond(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.NanoSecond()
	}
}

func ExampleFaker_TimeZone() {
	Global.Seed(11)
	fmt.Println(Global.TimeZone())
	// Output: Kaliningrad Standard Time
}

func BenchmarkTimeZone(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.TimeZone()
	}
}

func ExampleFaker_TimeZoneFull() {
	Global.Seed(11)
	fmt.Println(Global.TimeZoneFull())
	// Output: (UTC+03:00) Kaliningrad, Minsk
}

func BenchmarkTimeZoneFull(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.TimeZoneFull()
	}
}

func ExampleFaker_TimeZoneAbv() {
	Global.Seed(11)
	fmt.Println(Global.TimeZoneAbv())
	// Output: KST
}

func BenchmarkTimeZoneAbv(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.TimeZoneAbv()
	}
}

func ExampleFaker_TimeZoneOffset() {
	Global.Seed(11)
	fmt.Println(Global.TimeZoneOffset())
	// Output: 3
}

func BenchmarkTimeZoneOffset(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.TimeZoneOffset()
	}
}
