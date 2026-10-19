// Package leap provides functions for determining whether a year is a leap year.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear takes a year and returns true if it's a leap year, otherwise returns false.
func IsLeapYear(year int) bool {
    if year % 4 != 0 {
        return false
    } else if year % 100 == 0 && year % 400 != 0 {
        return false
    } else {
        return true
    }
}
