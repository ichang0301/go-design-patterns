# Facade design pattern

Imagine that we group many proxies in a single point such as a file or a library. This could be a facade pattern.

## What is facade

In architectural terms, a facade is the front wall that hides the rooms and corridors of a building.
It protects its inhabitants from cold and rain, and provides them privacy.
It orders and divides the dwellings.

The facade design pattern does the same, but in our code.
It shields the code from unwanted access, orders some calls, and hides the complexity scope from the user.

## Objectives

- When you want to decrease the complexity of some parts of our code. You hide that complexity behind the facade by providing a more easy-to-use method
- When you want to group actions that are cross-related in a single place
- When you want to build a library so that others can use your products without worrying about how it all works
