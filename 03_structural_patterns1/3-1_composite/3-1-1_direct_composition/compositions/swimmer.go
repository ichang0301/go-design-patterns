package compositions

import "fmt"

type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func() // closure. In Go, functions are first-class citizens and they can be used as parameters, fields, or arguments just like any variable.
}

type CompositeSwimmerB struct {
	Trainer Trainer
	Swimmer Swimmer
}

type CompositeSwimmerC struct {
	Trainer
	Swimmer
}

type Shark struct {
	Animal
	Swim func()
}
