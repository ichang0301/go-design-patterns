# Concurrency Patterns 1

Concurrency patterns, we will mostly manage the timing execution and order execution of applications that has more than one `flow`.

## Barrier Design Pattern

Barrier pattern is a very common pattern, especially when we have to wait for more than one response from different Goroutines before letting the program continue.

## Future Design Pattern

Future pattern allows us to write an algorithm that will be executed eventually in time (or not) by the same Goroutine or a different one.

## Pipeline Design Pattern

Pipeline pattern is a powerful pattern to build complex synchronous flows of Goroutines that are connected with each other according to some logic.
