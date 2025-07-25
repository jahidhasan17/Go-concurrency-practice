# Fan-out/Fan-in

- Create a pipeline where:
    * One goroutine generates numbers
    * Three worker goroutines square the numbers
    * One goroutine merges the results and prints them

- Ensure the program closes all channels properly when done

Note: Order is not guaranteed here

# Outcome
1. You will learn the real power of golang concurrently - how multiple worker finish some tasks and merge it
2. The importance of closing `channel`
    Try to remove this code from merge method
    ```
	go func ()  {
		wg.Wait()
		close(mergedCh)
	}()
    ```
3. You will learn how to fix the problem - `fatal error: all goroutines are asleep - deadlock!`