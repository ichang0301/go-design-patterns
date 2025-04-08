package main

import (
	"fmt"

	"github.com/ichang0301/go-design-patterns/03_structural_patterns1/3-1_composite/3-1-1_direct_composition/compositions"
)

func Swim() {
	fmt.Println("Swimming")
}

func main() {
	localSwim := Swim

	swimmerA := compositions.CompositeSwimmerA{
		MySwim: &localSwim,
	}
	swimmerA.MyAthlete.Train()
	(*swimmerA.MySwim)()

	swimmerB := compositions.CompositeSwimmerB{
		Trainer: &compositions.Athlete{},
		Swimmer: &compositions.SwimmerImplementor{},
	}
	swimmerB.Trainer.Train()
	swimmerB.Swimmer.Swim()

	// In the case of swimmerC, costs are saved at runtime.
	// However, problems may not be found during compilation.
	swimmerC := compositions.CompositeSwimmerC{
		&compositions.Athlete{},
		&compositions.SwimmerImplementor{},
	}
	swimmerC.Train()
	swimmerC.Swim()

	fish := compositions.Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}
