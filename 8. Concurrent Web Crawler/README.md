# Concurrent Web Crawler

Simulate a web crawler that:

    - Takes a list of URLs as input
    - Uses goroutines to fetch each URL (simulate with random sleep)
    - Limits concurrency to 4 simultaneous fetches
    - Collects results in a thread-safe manner