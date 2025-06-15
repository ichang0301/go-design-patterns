# Future Design Pattern (Promise Design Pattern)

The future design pattern is a quick and easy way to achieve concurrent structures for asynchronous programming.
Defining each possible behaviour of an action before executing them in different Goroutines provides event-driven programming by default(e.g. Node.js). This approach is to achieve using a `fire-and-forget` that handles all possible results in an action.

We can launch a new Future within a Future and embed as many Futures as we want in the same Goroutine (or new ones).

| Future |                  |            |        |              |
| ------ | ---------------- | ---------- | ------ | ------------ |
|        | Success function |            |        |              |
|        |                  | Go routine |        |              |
|        |                  |            | Future |              |
|        |                  |            |        | Success func |
|        |                  |            |        | Fail func    |
|        |                  |            |        | Execute func |
|        | Fail function    |            |        |              |
|        | Execute function |            |        |              |

## Objectives

- Delegate the action handler to a different Goroutine
- Stack many asynchronous calls between them (an asynchronous call that calls another asynchronous call in its results)
