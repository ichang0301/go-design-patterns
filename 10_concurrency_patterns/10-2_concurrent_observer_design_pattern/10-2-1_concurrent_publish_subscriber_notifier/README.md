# 10-2-1_concurrent_publish_subscriber_notifier

## Acceptance criteria

- We must have a publisher with a `PublishingCh` method that returns a channel to send messages through and triggers a `Notify` method on every observer subscribed
- We must have a method to add new subscribers to the publisher
- We must have a method to remove new subscribers from the publisher
- We must have a method to stop a subscriber
- We must have a method to stop a `Publisher` interface that will also stop all subscribers
- All inter Goroutine communication must be synchronized so that no Goroutine is locked waiting for a response. In such cases, an error is returned after the specified timeout period has passed

## Tips

### How to test the functionality prints to the `stdout`

In three ways:

- Capturing the `stdout` method
- Injecting an `io.Writer` interface to print to it. This is the preferred solution, as it makes the code more manageable
- Redirecting the `stdout` method to a different file
