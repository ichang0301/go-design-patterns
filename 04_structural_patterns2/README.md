# Structural Patterns 2

## Proxy design pattern

- Used when wrapping to hide something like secrets or complicated information
- Used when you want to create a new abstraction layer to make changes easier

## Decorator design pattern

- Decorate a type dynamically. The decoration may or may not be there, or it may be composed of one or many types
- The decorator pattern is more flexible than proxy pattern. A decorator might implement the entire interface that the type it decorates also implements **or not**. So if you have an interface with 10 methods and a decorator that just implements one of them and it will still be valid. But the decorator pattern is weaker than proxy pattern, because you could have errors at runtime, if you forget to implement any interface method
