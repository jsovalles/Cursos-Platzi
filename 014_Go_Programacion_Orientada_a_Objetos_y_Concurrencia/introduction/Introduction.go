package introduction

import (
	"fmt"
	"strconv"
	"time"
)

func Introduction() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	// Capturando valor y error
	myValue, err := strconv.ParseInt("NaN", 0, 64)

	// Validando errores.
	if err != nil {
		fmt.Println("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Mapa clave valor.
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	// Slice de enteros.
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	//c := make(chan int)
	//go doSomething(c)
	//<-c

	g := 25
	fmt.Println(g) // imprime el valor entero 25
	h := &g
	fmt.Println(h) // imprimer la direccion de memoria.
	i := *h
	fmt.Println(i) // Imprime el valor por dentro de g

	ChannelExample()
}

func ChannelExample() {
	c := make(chan int)
	go doSomething(c)
	go doOtherThing()
	go doSomething(c)
	<-c
	print("Lo imprimirá?")

}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	println("Do something")
	c <- 1
}

func doOtherThing() {
	println("quien se imprimirá primero?")
}
