package abstract_factory

import "testing"

func TestSportMotorbike(t *testing.T) {
	MotorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := MotorbikeF.Build(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())
	t.Logf("Motorbike vehicle has %d seats\n", motorbikeVehicle.NumSeats())

	sportBike, ok := motorbikeVehicle.(Motorbike) // type assertion
	if !ok {
		t.Fatal("Motorbike vehicle is not of type Motorbike")
	}

	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())
}

func TestCruiseMotorbike(t *testing.T) {
	MotorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := MotorbikeF.Build(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())
	t.Logf("Motorbike vehicle has %d seats\n", motorbikeVehicle.NumSeats())

	cruiseBike, ok := motorbikeVehicle.(Motorbike) // type assertion
	if !ok {
		t.Fatal("Motorbike vehicle is not of type Motorbike")
	}

	t.Logf("Sport motorbike has type %d\n", cruiseBike.GetMotorbikeType())
}
