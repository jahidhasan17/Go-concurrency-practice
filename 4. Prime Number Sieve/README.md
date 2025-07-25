# Prime Number Sieve

* Implement the Sieve of Eratosthenes using goroutines and channels:
* The first goroutine generates numbers 2, 3, 4, ...
* Each prime number creates a new goroutine that filters out its multiples
* Print the prime numbers as they're found

# Test Yourself
```go
prime, ok := <- ch
if !ok {
    break
}
```
Here `ch` is the latest channel, and we're checking that if the ch is close then break the loop. It works. But how it works even though we just check the last channel is close or not.