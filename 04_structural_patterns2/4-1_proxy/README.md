# Proxy Design Pattern

The Proxy pattern is a simple pattern that provides interesting features and possibilities with very little effort.

## Objectives

- Hide an object behind the proxy so the features can be hidden, restricted, and so on
- Provide a new abstraction layer that is easy to work with, and can be changed easily

## Acceptance criteria

- All accesses to the database of users will be done through the Proxy type.
- A stack of `n` number of recent users will be kept in the Proxy.
- If a user already exists in the stack, it won't query the database, and will return the stored one
- If the queried user doesn't exist in the stack, it will query the database, remove the oldest user in the stack if it's full, store the new one, and return it
