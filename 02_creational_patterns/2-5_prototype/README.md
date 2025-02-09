# Prototype

## Objective

The prototype pattern is a powerful tool to build caches and default objects.
The main objective for the prototype design pattern is to avoid repetitive object creation.

- Maintain a set of objects that will be cloned to crete new instances
- Provide a default value of some type to start working on top of it
- Free CPU of complex object initialization to take more memory resources

## Difference between prototype and builder pattern

The key difference between prototype and builder pattern is that objects are cloned for the user instead of building them at runtime. You can also build a cache-like solution, storing information using a prototype.
