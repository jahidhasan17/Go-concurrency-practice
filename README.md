# Goroutine concurrency practice

A curated set of Go concurrency exercises designed to help you master goroutines, channels, and synchronization primitives. Each folder contains a focused problem, sample code, and explanations to reinforce key concepts in concurrent programming.

## What’s Inside

- [**Basic Goroutine Communication**](./Basic-Goroutine-Communication/README.md): Learn how to send and receive data between goroutines using channels.
- [**Channel Summation**](./Channel-Summation/README.md): Generate numbers and sum them concurrently.
- [**Multiple Workers**](./Multiple-Workers/README.md): Distribute tasks among multiple worker goroutines.
- [**Prime Number Sieve**](./Prime-Number-Sieve/README.md): Implement the Sieve of Eratosthenes using goroutines and channels.
- [**Timeout Handling**](./Timeout-Handling/README.md): Use channels and timeouts to handle slow operations.
- [**Fan-out/Fan-in**](./Fan-out-Fan-in/README.md): Process data with multiple workers and merge results.
- [**Rate Limiter**](./Rate-Limiter/README.md): Control the rate of operations using a token bucket pattern.
- [**Concurrent Web Crawler**](./Concurrent-Web-Crawler/README.md): Simulate a web crawler with limited concurrency.
- [**Pub/Sub System**](./Pub-Sub-System/README.md): Build a simple publish-subscribe system using channels.
- [**Concurrent Merge Sort**](./Concurrent-Merge-Sort/README.md): Sort data concurrently using merge sort and channels.
- [**MQTT Example**](./MQTT-Example/README.md): Demonstrates concurrent publish/subscribe using MQTT protocol.

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