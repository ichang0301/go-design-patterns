# Flyweight design pattern

If you have to create and store too many objects of some heavy type that are fundamentally equal, then you can easily solve it with the flyweight pattern using `pointer`.

It's very commonly used in computer graphics and the video game industry.
But it's not so much in enterprise applications.

The flyweight pattern is additional help of the factory pattern that is usually in charge of encapsulating object creation.
cf. `/02_creational_patterns/2-3_factory`

## What's the difference between singleton and flyweight

With the singleton pattern, we ensure that the same type is created only once. Also, the singleton pattern is a creational pattern.

With flyweight, which is a structural pattern, we arn't worried about how the objects are created or how they do their business, but about how to structure a type to contain heavy information in a light way.
