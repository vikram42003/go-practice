// Package weather provides variables and functions for weather forecasting. 
package weather

var (
    // CurrentCondition stores the current weather condition.
	CurrentCondition string
    // CurrentLocation stores the current location for the weather forecast.
	CurrentLocation  string
)

// Forecast takes strings city and condition as input, assigns them to CurrentLocation and CurrentCondition, and returns a string describing the weather condition at the location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
