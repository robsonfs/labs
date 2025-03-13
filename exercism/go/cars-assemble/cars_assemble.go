package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	ratePerHour := CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(ratePerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount < 10 {
		return uint(carsCount * 10000)
	}

	r := carsCount % 10

	t := (carsCount - r) / 10

	return uint((r * 10000) + (t * 95000))
}
