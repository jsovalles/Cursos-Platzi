package net

import (
	"fmt"
	"net"
)

func NetExample() {
	// Escanear cada puerto y hacer una conexión
	for port := 0; port < 100; port++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", port))
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Println("Port", port, "is open")
	}
}
