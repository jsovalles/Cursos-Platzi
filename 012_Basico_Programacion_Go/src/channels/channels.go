package channels

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func ChannelExample() {
	c := make(chan string, 1)
	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)
}
