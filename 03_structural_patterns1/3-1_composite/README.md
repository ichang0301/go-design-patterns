# Composite design pattern

The composite design pattern favours composition (commonly defined as a `has a` relationship) over inheritance (an `is a` relationship).
All in all, Go doesn't have inheritance because it doesn't need it.

## Objectives

The objective of the composition is to avoid hierarchy hell where the complexity of an application could grow too much, and the clarity of the code is affected.

## Two types of composition

- the `direct` composition
- the `embedding` composition

## Why is there no test?

It doesn't have much to test apart from the structure itself.
We won't write unit tests, and we'll simply describe the ways to create those compositions in Go.

## Composite pattern versus inheritance

In go, you can simply composite the son struct (sub class in the object oriented programming) with the parent without embedding so that you can access the Parent instance later:

```go
type Parent struct {
    SomeField int
}

type Son struct {
    P Parent
}

func GetParentField(p *Parent) int {
    fmt.Println(p.SomeField)
}

son := Son{}
GetParentField(son.P)
```
