package functions

import (
	"fmt"
	"time"
)

func FuncionAnonimaExample() {
	func() {
		println("Hello")
	}()

	y := func(x int) int {
		return x * 2
	}(5)
	fmt.Println(y)

	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(2 * time.Second)
		fmt.Println("Finishing function")
		c <- 1
	}()
	fmt.Println(<-c)
}
