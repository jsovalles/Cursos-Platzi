package rangecloseselect

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func ChannelsExample() {

	c := make(chan string, 2)
	c <- "Mansaje1"
	c <- "Mensaje2"
	//len is used to check how many channels does the variable has, cap checks how many channels the variable can use
	fmt.Println(len(c), cap(c))

	//close is used to close a channel, so it doesn't admit any more values (if it applies)
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go message("mansaje1", email1)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}
