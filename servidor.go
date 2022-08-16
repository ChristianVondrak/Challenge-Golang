package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

// Selector de Canales para mostrar en pantalla
var channel1 = make(chan string, 1)    //Declaracion de Channel1 Global
var channel2 = make(chan string, 1)    //Declaracion de Channel2 Global
var channelclose = make(chan struct{}) //Cierra el metodo recibe

func main() {
	go recibe(channel1, channelclose, 1)
	go recibe(channel2, channelclose, 2)
	//enviar(channel1, "hola")
	channelclose <- struct{}{}

	go servidor()
	var input string
	fmt.Scanln(&input)
}

func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	var msg string
	var channel int

	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
		return
	} //else {

	//fmt.Println("Message: ", msg)
	//}
	err = gob.NewDecoder(c).Decode(&channel)
	if err != nil {
		fmt.Println(err)
		return
	} //else {

	//fmt.Println("Chanel: ", channel)
	//}

	switch channel {
	case 1:
		enviar(channel1, msg)

	case 2:
		enviar(channel2, msg)
	}

}

func recibe(c <-chan string, b <-chan struct{}, nchan int) {
	for {
		select {
		case i := <-c:
			fmt.Println("Channel " + strconv.Itoa(nchan) + ": " + i)
			//fmt.Println("La funcion Recibe termino")
		case <-b:
			break
		}
	}
}

func enviar(c chan string, m string) {
	c <- m
	//fmt.Println("La funcion enviar termino")
}
