// Package weather provides functionality to forecast the current weather conditions
// for various cities in Goblinocus. It maintains the current location and weather
// condition as package-level variables.
package weather

var (
	// CurrentCondition stores the latest weather condition for the current location.
	CurrentCondition string
	// CurrentLocation stores the name of the city for which the weather condition is being forecasted.
	CurrentLocation string
)

// Forecast updates the current location and weather condition, and returns a formatted string
// describing the current weather condition for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
