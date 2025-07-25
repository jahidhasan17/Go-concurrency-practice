Multiple Workers

    - Create 3 worker goroutines that receive tasks from a channel
    - Send 10 tasks (numbers 1-10) to these workers
    - Each worker should print the task it's processing and sleep for 1 second
    - Use a WaitGroup to ensure all tasks are completed before the program exits