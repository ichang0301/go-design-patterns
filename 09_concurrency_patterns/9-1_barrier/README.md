# Barrier Design Pattern

Barrier patten's purpose: put-up a barrier so that nobody passes until we have all the results we need, something quite common in concurrent applications.

Barrier pattern is not only useful to make network requests, we could also use it to split some task into multiple Goroutines. For example, an expensive operation could be split into a few smaller operations distributed in different Goroutines to maximize parallelism and achieve better performance.

## Objectives

- Compose the value of a type with the data coming from one or more Goroutines
- Control the correctness of any of those incoming data pipes so that no inconsistent data is returned. We don't want a partially filled result because one of the pipes has returned an error
