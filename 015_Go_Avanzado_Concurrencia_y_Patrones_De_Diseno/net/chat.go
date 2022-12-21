package net

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incoming       = make(chan Client)
	leaving        = make(chan Client)
	serverMessages = make(chan string)
)

// Client1 -> Server -> HandleConnection(Client1)
// HanldeConnnection handles the client connection of a single user
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	// Create the client channel
	clientMessage := make(chan string)
	go MessageWrite(conn, clientMessage)

	// Get the client's name
	clientName := conn.RemoteAddr().String()

	// Send just to the client his name
	clientMessage <- "Welcome to the server, your name is " + clientName

	// Send the message to all the clients
	serverMessages <- clientName + " has joined!"

	// Add the client to the list of clients
	incoming <- clientMessage

	// Read the messages from the client
	// If the loop breaks, that means that the client has disconnected
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		serverMessages <- clientName + ": " + inputMessage.Text()
	}

	// Remove the client from the list of clients
	leaving <- clientMessage
	serverMessages <- clientName + " has left!"
}

// MessageWrite recives messages from the channel and writes them to the client
func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}

// Broadcast sends the message to all the clients, and handles incoming
// and outgoing connections
func Broadcast() {

	// Map the clients to a boolean
	clients := make(map[Client]bool)

	for {
		// Multiplex the messages
		select {
		// We get a new message
		case msg := <-serverMessages:
			// Send the message to all the clients
			for client := range clients {
				client <- msg
			}
		// We get a new client
		case client := <-incoming:
			clients[client] = true
		// A client that has disconnected
		case client := <-leaving:
			delete(clients, client)
			close(client)
		}
	}
}

func Server() {
	flag.Parse()

	// Create the server and listen to it
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// Start the broadcast
	go Broadcast()

	// Listen for connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go HandleConnection(conn)
	}
}
