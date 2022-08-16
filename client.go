package main

import (
	//"bufio"
	"encoding/gob"
	"fmt"
	"net"
	//"os"
)

func client() {
	for {
		c, err := net.Dial("tcp", ":9999")
		if err != nil {
			fmt.Println(err)
			return
		}
		var msg string
		var channel int
		fmt.Println("Ingresa el mensaje a enviar: ")
		fmt.Scanln(&msg)

		err = gob.NewEncoder(c).Encode(msg)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ingresa Channel a enviar (1 o 2): ")
		fmt.Scanln(&channel)
		err = gob.NewEncoder(c).Encode(channel)
		if err != nil {
			fmt.Println(err)
		}
		c.Close()
	}
}

func main() {
	client()

	//var input string
	//fmt.Scanln(&input)
}
