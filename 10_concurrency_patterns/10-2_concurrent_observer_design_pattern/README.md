# Concurrent Observer Design Pattern

We will implement the observer design pattern that we implemented previously on 07_behavioral_patterns, but with a concurrent structure and thread safety.

## Description

- Now, the access to the list of subscribers must be serialized. If we are reading the list with one Goroutine, we cannot be removing a subscriber from it or we will have a race
- When a subscriber is removed, the subscriber's Goroutine must be closed too, or it will keep iterating forever and we will run into Goroutine lakes
- When stopping the publisher, all subscribers must stop their Goroutines, too.

## Objectives

- Providing an event-driven architecture where one event can trigger one or more actions
- Uncoupling the actions that are performed from the event that triggers them
- Providing more than one source event that triggers the same action
