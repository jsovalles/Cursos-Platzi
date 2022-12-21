package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func BufferedAndUnbufferedChannelsExample() {
	// Unbuffered channel
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

	// Buffered channel
	ch2 := make(chan int, 1)
	ch2 <- 2
	fmt.Println(<-ch2)

	ch2 <- 3
	fmt.Println(<-ch2)

}

/*
traffic light.

this uses channels and wait groups to 1. execute only 2 doSomething() func
at a time and 2. be able to wait for all of them.

in order of execution it'll:
c := [][] -- two free spaces
c := [routine][] -- one free space
c := [routine][routine] -- all occupied
c := [][routine] -- one free space
*/

func doSomethingStoplight(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id: %d -> started...\n", i)
	time.Sleep(time.Second * 4)
	fmt.Printf("Id: %d -> finished...\n", i)

	<-c // frees the space for new routines
}

func BufferedChannelsAsStoplight() {
	c := make(chan int, 2) // creates a buffered channel with a capacity of two
	var wg sync.WaitGroup  // creates wait group

	for i := 0; i < 10; i++ {
		c <- 1    // allocate a new "instance" in the free space
		wg.Add(1) // adds to the wait group
		go doSomethingStoplight(i, &wg, c)
	}

	wg.Wait()
}
