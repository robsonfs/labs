package lasagna

func countElements[T comparable](arr []T, item T) int {
	counter := 0

	for k := 0; k < len(arr); k++ {
		if arr[k] == item {
			counter += 1
		}
	}

	return counter
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg_prep_timer int) int {
	if avg_prep_timer > 0 {
		return len(layers) * avg_prep_timer
	} else {
		return len(layers) * 2
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	amountNoodles := countElements(layers, "noodles")
	amountSauce := countElements(layers, "sauce")

	return amountNoodles * 50, float64(amountSauce) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions_number int) (new_quantities []float64) {
	if portions_number < 1 {
		panic("Nothing to cook!")
	}

	new_quantities = []float64{}

	for k := 0; k < len(quantities); k++ {
		scaled_quantity := quantities[k] * (float64(portions_number) / 2.0)
		new_quantities = append(new_quantities, scaled_quantity)
	}

	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
