package chance

import "math/rand"

func intRandInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func floatRandInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func shuffleSlice[T any](arr []T) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return intRandInRange(1, 20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return floatRandInRange(0.0, 12.0)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	shuffleSlice(animals)

	return animals
}
