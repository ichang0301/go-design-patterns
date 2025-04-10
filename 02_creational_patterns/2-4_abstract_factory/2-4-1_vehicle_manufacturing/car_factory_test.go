package abstract_factory

import "testing"

func TestLuxuryCar(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d wheels\n", carVehicle.NumWheels())
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumSeats())

	luxuryCar, ok := carVehicle.(Car) // type assertion
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Car vehicle has %d doors\n", luxuryCar.NumDoors())
}

func TestFamilyCar(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.Build(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d wheels\n", carVehicle.NumWheels())
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumSeats())

	familyCar, ok := carVehicle.(Car) // type assertion
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Car vehicle has %d doors\n", familyCar.NumDoors())
}
