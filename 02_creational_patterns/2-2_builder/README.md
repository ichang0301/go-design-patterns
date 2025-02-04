# Builder Design Pattern

reusing an algorithm to create many implementations of an interface

## Objectives

- Abstract complex creations so that object creation is separated from the object user
- Create an object step by step by filling its fields and creating the embedded objects
- Reuse the object creation algorithm between many objects

## Risks of using the Builder pattern with unstable algorithms

Try to avoid the Builder pattern when you are not completely sure that the algorithm is going to be more or less stable because any small change in this interface will affect all your builders and it could be awkward if you add a new method that some of your builders need and others Builders do not.
