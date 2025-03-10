// Package weather provides functions and methods to track and forecast
// the weather conditions of various cities in Goblinocus.
package weather

// CurrentCondition stores information abouot the current weather conditions.
var CurrentCondition string

// CurrentLocation stores the specific location where the weather conditions
// refers to.
var CurrentLocation string

// Forecast shows the current weather forecast for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
