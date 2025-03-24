package speed

// Car describes a car and his attibutes in the game.
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	var localDistance int
	if car.battery < car.batteryDrain {
		localDistance = 0
	} else {
		localDistance = car.distance + car.speed
	}

	var localBattery int
	if (car.battery - car.batteryDrain) < 0 {
		localBattery = car.battery
		localDistance = car.distance
	} else {
		localBattery = car.battery - car.batteryDrain
	}
	return Car{
		battery:      localBattery,
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
		distance:     localDistance,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	panic("Please implement the CanFinish function")
}
