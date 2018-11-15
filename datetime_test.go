package gofakeit

import (
	"fmt"
	"testing"
	"time"
)

func ExampleFaker_Date() {
	Global.Seed(11)
	fmt.Println(Global.Date())
	// Output: 1989-01-07 04:14:25.685339029 +0000 UTC
}

func BenchmarkDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Date()
	}
}

func ExampleFaker_DateRange() {
	Global.Seed(11)
	fmt.Println(Global.DateRange(time.Unix(0, 484633944473634951), time.Unix(0, 1431318744473668209))) // May 10, 1985 years to May 10, 2015
	// Output: 2012-02-04 14:10:37.166933216 +0000 UTC
}

func BenchmarkDateRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.DateRange(time.Now().AddDate(-30, 0, 0), time.Now())
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
	for i := 0; i < b.N; i++ {
		Global.Month()
	}
}

func ExampleFaker_Day() {
	Global.Seed(11)
	fmt.Println(Global.Day())
	// Output: 12
}

func BenchmarkDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Day()
	}
}

func BenchmarkWeekDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.WeekDay()
	}
}

func ExampleFaker_Year() {
	Global.Seed(11)
	fmt.Println(Global.Year())
	// Output: 1989
}

func BenchmarkYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Year()
	}
}

func ExampleFaker_Hour() {
	Global.Seed(11)
	fmt.Println(Global.Hour())
	// Output: 0
}

func BenchmarkHour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Hour()
	}
}

func ExampleFaker_Minute() {
	Global.Seed(11)
	fmt.Println(Global.Minute())
	// Output: 0
}

func BenchmarkMinute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Minute()
	}
}

func ExampleFaker_Second() {
	Global.Seed(11)
	fmt.Println(Global.Second())
	// Output: 0
}

func BenchmarkSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Second()
	}
}

func ExampleFaker_NanoSecond() {
	Global.Seed(11)
	fmt.Println(Global.NanoSecond())
	// Output: 196446360
}

func BenchmarkNanoSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.NanoSecond()
	}
}

func ExampleFaker_TimeZone() {
	Global.Seed(11)
	fmt.Println(Global.TimeZone())
	// Output: Kaliningrad Standard Time
}

func BenchmarkTimeZone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.TimeZone()
	}
}

func ExampleFaker_TimeZoneFull() {
	Global.Seed(11)
	fmt.Println(Global.TimeZoneFull())
	// Output: (UTC+03:00) Kaliningrad, Minsk
}

func BenchmarkTimeZoneFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.TimeZoneFull()
	}
}

func ExampleFaker_TimeZoneAbv() {
	Global.Seed(11)
	fmt.Println(Global.TimeZoneAbv())
	// Output: KST
}

func BenchmarkTimeZoneAbv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.TimeZoneAbv()
	}
}

func ExampleFaker_TimeZoneOffset() {
	Global.Seed(11)
	fmt.Println(Global.TimeZoneOffset())
	// Output: 3
}

func BenchmarkTimeZoneOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.TimeZoneOffset()
	}
}
