package pizza

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (d *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("an IngredientAdd is needed in the Ingredient field of the Onion")
	}

	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", s, "onion"), nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("an IngredientAdd is needed in the Ingredient field of the Meat")
	}

	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", s, "meat"), nil
}
