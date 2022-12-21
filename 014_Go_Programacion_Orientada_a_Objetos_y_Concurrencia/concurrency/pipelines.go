package concurrency

import "fmt"

func PipelineExample() {
	generator := make(chan int)
	doubles := make(chan int)
	go Generator(generator)
	go Double(generator, doubles)
	PrintChannel(doubles)
}

//Write channel
func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

//Read & write channels
func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

// Read channel
func PrintChannel(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}
