package faker

import (
	"strconv"
	"time"
)

// Date will generate a random time.Time struct between 1900-Now
func (f *Faker) Date() time.Time {
	return time.Date(f.Year(), time.Month(f.Number(0, 12)), f.Day(), f.Hour(), f.Minute(), f.Second(), f.NanoSecond(), time.UTC)
}

// DateStr will generate a random date-time between 1900-Now
// the format is: "2006-01-02 15:04:05.999999999 -0700 MST"
func (f *Faker) DateStr() string {
	return f.Date().String()
}

// DateCurrentYear will generate a random time.Time within the current year (UTC)
func (f *Faker) DateCurrentYear() time.Time {
	return time.Date(time.Now().UTC().Year(), time.Month(f.Number(0, 12)), f.Day(), f.Hour(), f.Minute(), f.Second(), f.NanoSecond(), time.UTC)
}

// DateCurrentYearStr will generate a random date-time within the current year (UTC)
// the format is: "2019-01-02 15:04:05.999999999 -0700 MST"
func (f *Faker) DateCurrentYearStr() string {
	return f.DateCurrentYear().String()
}

// DateRange will generate a random time.Time struct between a start and end date
func (f *Faker) DateRange(start, end time.Time) time.Time {
	return time.Unix(0, int64(f.Number(int(start.UnixNano()), int(end.UnixNano())))).UTC()
}

// Month will generate a random month string in English (January, February ...)
func (f *Faker) Month() string {
	return time.Month(f.Number(1, 12)).String()
}

// Day will generate a random day between 1 - 31
func (f *Faker) Day() int {
	return f.Number(1, 31)
}

// WeekDay will generate a random weekday string (Monday-Sunday)
func (f *Faker) WeekDay() string {
	return time.Weekday(f.Number(0, 6)).String()
}

// Year will generate a random year between 1900 - current year
func (f *Faker) Year() int {
	return f.Number(1900, time.Now().Year())
}

// Hour will generate a random hour - in military time
func (f *Faker) Hour() int {
	return f.Number(0, 23)
}

// Minute will generate a random minute
func (f *Faker) Minute() int {
	return f.Number(0, 59)
}

// Second will generate a random second
func (f *Faker) Second() int {
	return f.Number(0, 59)
}

// NanoSecond will generate a random nano second
func (f *Faker) NanoSecond() int {
	return f.Number(0, 999999999)
}

// TimeZone will select a random timezone string
func (f *Faker) TimeZone() string {
	return f.getRandValue([]string{"timezone", "text"})
}

// TimeZoneFull will select a random full timezone string
func (f *Faker) TimeZoneFull() string {
	return f.getRandValue([]string{"timezone", "full"})
}

// TimeZoneAbv will select a random timezone abbreviation string
func (f *Faker) TimeZoneAbv() string {
	return f.getRandValue([]string{"timezone", "abr"})
}

// TimeZoneOffset will select a random timezone offset
func (f *Faker) TimeZoneOffset() float32 {
	value, err := strconv.ParseFloat(f.getRandValue([]string{"timezone", "offset"}), 32)
	if err != nil {
		errorLogger.Printf("(TimeZoneOffset) %s\n", err)
		return 0
	}
	return float32(value)
}
