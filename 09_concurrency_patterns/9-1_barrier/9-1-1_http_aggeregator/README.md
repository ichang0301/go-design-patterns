# 9-1-1_http_aggregator

## Acceptance criteria

Our main objectives in this app is to get a merged response of two different calls, so we can describe our acceptance criteria like this:

- Print on the console the merged result of the two calls to <http://httpbin.org/headers> and <http://httpbin.org/user-agent> URLs. These are a couple of public endpoints that respond with data from the incoming connections. They are very popular for testing purposes. You will need an internet connection to do this exercise
- If any of the calls fails, it must not print any result. it must print just the error message (or error messages if both calls failed).
- The output must be printed as a composed result when both calls have finished. It means that we cannot print the result of one call and then the other.
