package goroutines

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func GoRoutinesExample() {

	var wg sync.WaitGroup

	fmt.Println("Hello")
	// We define a goroutine so the execution of main waits for the other one to end
	wg.Add(1)
	// With go it doesn't stay on the same thread, so it won't execute with the main thread unless we control the goroutine in the program
	go say("world", &wg)
	// Tells the main thread (main execution) to wait until all the goroutines end on their execution
	wg.Wait()

	// anonymous function, it won't execute since is not on the WaitGroup (it can be executed if we use a time.sleep)
	go func(text string) {
		fmt.Println(text)
	}("Adios")
}
