# Singleton Design Pattern

having a unique instance of a type in the entire program

## Objectives

- We need a single, shared value, of some particular type
- We need to restrict object creation of some type to a single unit along the entire program

## For example

- When you want to use **the same connection to a database to make every query**
- When you open a Secure Shell(SSH) connection to a server to do a few tasks, and **don't want to reopen the connection for each task**
- If you need to limit the access to some variable or space, you use a Singleton as the door to this variable
- If you need to limit the number of calls to some places, you create a Singleton instance to make the calls in the accepted window
