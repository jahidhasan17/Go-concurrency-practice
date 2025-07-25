# Goroutine-And-Channel-Practice-Problems

# Goroutine-And-Channel-Practice-Problems

## Problems

- [1. Basic Goroutine Communication](1.%20Basic%20Goroutine%20Communication/README.md)
- [2. Channel Summation](2.%20Channel%20Summation/README.md)
- [3. Multiple Workers](3.%20Multiple%20Workers/README.md)

Keep in mind - 

    - Always clean up goroutines properly
    - Handle channel closing correctly
    - Use synchronization primitives like WaitGroups when needed
    - Consider edge cases and error handling

# Some Concepts
1. Channels are Passed by Value (But Have Reference-like Behavior)
    * When you pass a channel to a function, Go makes a copy of the channel reference. So if we change the variable after call the method, the channel in the method doesn't affect.