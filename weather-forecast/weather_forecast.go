// Package weather This is a weather package.
package weather

var (
	// CurrentCondition describes the condition at a place at a time.
	CurrentCondition string
	// CurrentLocation describes the location.
	CurrentLocation string
)

// Forecast returns a string with the weather condition at a certain place.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
