# Workers Pool Design Pattern

Goroutines are light, but the work they perform could be very heavy.
A workers pool helps us to solve the problem that we cannot let an app create an unlimited amount of Goroutines.

## Description

With a pool of workers, we want to bound the amount of Goroutines available so that we have a deeper control of the pool of resources.
This is easy to achieve by creating a channel for each worker and having workers with either an idle or busy status.

## Objectives

- Control access to shared resources using quotas
- Create a limited amount of Goroutines per app
- Provide more parallelism capabilities to other concurrent structures
