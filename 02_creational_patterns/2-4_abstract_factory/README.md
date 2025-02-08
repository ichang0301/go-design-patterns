# Abstract Factory

a factory of factories

## Objectives

- Provide a new layer of encapsulation for Factory methods that return a common interface for all factories
- Group common factories into a super Factory (also called a factory of factories)

## Builder pattern vs Abstract factory

The abstract factory and builder patterns can both resolve the same problem, but your particular needs will help you find the slight differences that should lead you to take on solution or the other.

### Builder pattern

We had an unstructured list of objects (cars with motorbikes in the same factory) with the builder pattern.
Also, we encouraged reusing the building algorithm in the builder pattern.

## Abstract factory

We have a very structured list of vehicles ( the factory for motorbikes and a factory for cars) in the abstract factory.
We also didn't mix the creation of cars with motorbikes, providing more flexibility in the creation process.

## Tips

When you have an interface instance, which is essentially a pointer to a struct, you just have access to the interface methods. With type assertion, you can tell the compiler the type of the pointed struct, so you can access the entire struct fields and methods.
Type assertion is also known as `casting` in other languages.
