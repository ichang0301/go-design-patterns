# Structural Patterns 2

## Proxy design pattern

- Used when wrapping to hide something like secrets or complexity
- Used when you want to create a new abstraction layer to make changes easier

## Decorator design pattern

- Decorate a type dynamically. The decoration may or may not be there, or it may be composed of one or many types
- The decorator pattern is more flexible than proxy pattern. A decorator might implement the entire interface that the type it decorates also implements **or not**. So if you have an interface with 10 methods and a decorator that just implements one of them and it will still be valid. But the decorator pattern is weaker than proxy pattern, because you could have errors at runtime, if you forget to implement any interface method

## Facade design pattern

Grouping multiple proxies into a single point, e.g. an HTTP REST API.

## Flyweight design pattern

It is used when dealing with a lot of 'heavy but similarly structured' data in fields such as games and graphics.

We will practice this flyweight design pattern with a game match program that manages the history of each team and player using pointers.
In order not to refer to the database every time a large data is requested, we will store the team pointer of the data acquired once in `map[int]*Team`.
