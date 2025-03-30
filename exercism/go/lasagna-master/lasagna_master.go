package lasagna

func countElements[T comparable](arr []T, item T) int {
	counter := 0

	for _, value := range arr {
		if value == item {
			counter++
		}
	}

	return counter
}

// PreparationTime caclculates the total estimated time to prepare the meal.
func PreparationTime(layers []string, avg_prep_timer int) int {
	if avg_prep_timer > 0 {
		return len(layers) * avg_prep_timer
	} else {
		return len(layers) * 2
	}
}

// Quantities computes the amount of noodles and sauce needed
func Quantities(layers []string) (int, float64) {
	amountNoodles := countElements(layers, "noodles")
	amountSauce := countElements(layers, "sauce")

	return amountNoodles * 50, float64(amountSauce) * 0.2
}

// AddSecretIngredient adds the secretsingredient from my friends list
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales quiantities up or down based on portions number
func ScaleRecipe(quantities []float64, portions_number int) (new_quantities []float64) {
	if portions_number < 1 {
		panic("Nothing to cook!")
	}

	new_quantities = []float64{}

	for _, value := range quantities {
		scaled_quantity := value * (float64(portions_number) / 2.0)
		new_quantities = append(new_quantities, scaled_quantity)
	}

	return
}
