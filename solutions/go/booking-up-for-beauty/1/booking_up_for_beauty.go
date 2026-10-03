package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(d string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", d)
    if err != nil {
        panic("Time parsing failed")
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(d string) bool {
    date, err := time.Parse("January 2, 2006 15:04:05", d)
    if err != nil {
        panic("Time parsing failed")
    }
    return date.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(d string) bool {
    date, err := time.Parse("Monday, January 2, 2006 15:04:05", d)
    if err != nil {
        panic("Time parsing failed")
    }
	hour := date.Hour()
    if hour >= 12 && hour <= 18 {
        return true
    } else {
        return false
    }
}

// Description returns a formatted string of the appointment time.
func Description(d string) string {
    t := Schedule(d)
	return fmt.Sprintf("You have an appointment on %s.", t.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
