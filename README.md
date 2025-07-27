# Go concurrency practice

A curated set of Go concurrency exercises designed to help you master goroutines, channels, and synchronization primitives. Each folder contains a focused problem, sample code, and explanations to reinforce key concepts in concurrent programming.

## What’s Inside

- [**Basic Goroutine Communication**](./1.%20Basic%20Goroutine%20Communication): Learn how to send and receive data between goroutines using channels.
- [**Channel Summation**](./2.%20Channel%20Summation): Generate numbers and sum them concurrently.
- [**Multiple Workers**](./3.%20Multiple%20Workers): Distribute tasks among multiple worker goroutines.
- [**Prime Number Sieve**](./4.%20Prime%20Number%20Sieve): Implement the Sieve of Eratosthenes using goroutines and channels.
- [**Timeout Handling**](./5.%20Timeout%20Handling): Use channels and timeouts to handle slow operations.
- [**Fan-out/Fan-in**](./6.%20Fan-out-Fan-in): Process data with multiple workers and merge results.
- [**Rate Limiter**](./7.%20Rate%20Limiter): Control the rate of operations using a token bucket pattern.
- [**Concurrent Web Crawler**](./8.%20Concurrent%20Web%20Crawler): Simulate a web crawler with limited concurrency.
- [**Pub/Sub System**](./9.%20PubSub%20System): Build a simple publish-subscribe system using channels.
- [**Concurrent Merge Sort**](./10.%20Concurrent%20Merge%20Sort): Sort data concurrently using merge sort and channels.
- [**MQTT Example**](./MQTT): Demonstrates concurrent publish/subscribe using MQTT protocol.

## How to Use

1. Clone the repository.
2. Each folder contains a standalone Go program. Read the `README.md` inside each folder for problem details.
3. Run any example with:
   ```sh
   go run main.go
   ```
4. Explore, modify, and experiment to deepen your understanding of Go concurrency.


## Best Practices

- Always clean up goroutines properly.
- Handle channel closing correctly.
- Use synchronization primitives like WaitGroups when needed.
- Consider edge cases and error handling.

# Some tricky Concepts
1. Channels are Passed by Value (But Have Reference-like Behavior)
    * When you pass a channel to a function, Go makes a copy of the channel reference. So if we change the variable after call the method, the channel in the method doesn't affect.

2. Dereferencing (*) creates a copy, so be careful when you dereference a variable.
