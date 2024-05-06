package advancedgo

import "fmt"

// Define an interface for vehicles (e.g., cars, trucks)
type Vehicle interface {
	drive()
}

// Concrete type: Car struct with a Drive method
type Car struct{}

func (c *Car) drive() {
	fmt.Println("The car is driving...")
}

// Another concrete type: Truck struct with a different Drive method
type Truck struct{}

func (t *Truck) drive() {
	fmt.Println("The truck is hauling heavy loads...")
}

// Define an interface for vehicles that can tow trailers
type TrailerTowingVehicle interface {
	Vehicle // Inherit the Vehicle interface
	tow()
}

// Implement the TrailerTowingVehicle interface on Car struct
func (c *Car) tow() {
	fmt.Println("The car is towing a trailer...")
}

func InterfaceUseCase() {
	vehicles := []Vehicle{
		&Car{},
		&Truck{},
	}

	for _, v := range vehicles {
		switch v.(type) {
		case *Car:
			c := v.(*Car)
			c.drive()
			// You can call c.tow() if Car implements TrailerTowingVehicle
		case *Truck:
			t := v.(*Truck)
			t.drive()
			// You can't call t.tow() because Trucks don't implement TrailerTowingVehicle
		}
	}
}
