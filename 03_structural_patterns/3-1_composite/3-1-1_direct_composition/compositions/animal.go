package compositions

import "fmt"

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}
