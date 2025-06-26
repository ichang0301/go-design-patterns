# 4-2-1_pizza

With the Decorator pattern, we can keep stacking `IngredientAdds` which call their inner pointer to add functionality to `PizzaDecorator`.
We won't touching the core type either, nor modifying or implementing new things.
All the new features will be implemented by an external type.

## Acceptance criteria

- We must have the main interface that all decorators will implement. This interface will be called `IngredientAdd`, and it will have the `AddIngredient() string` method
- We must have a core `PizzaDecorator` type (the decorator) that we will add ingredients to
- We must have an ingredient "onion" implementing the same `IngredientAdd` interface that will add the string `onion` to the returned pizza.
- We must have a ingredient "meat" implementing the `IngredientAdd` interface that will add the string `meat` to the returned pizza.
- When calling `AddIngredient` method on the top object, it must return a fully decorated `pizza` with the text `Pizza with the following ingredients: meat, onion`
