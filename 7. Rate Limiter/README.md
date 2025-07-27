# Rate Limiter

Implement a rate limiter using a buffered channel as a token bucket:

    - Allow up to 5 operations per second
    - Each operation is simulated with a sleep and print
    - Use a ticker to replenish tokens