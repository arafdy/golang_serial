package main

import (
	"bufio"
	"github.com/tarm/serial"
	"log"
)

func main()  {


	c := &serial.Config{Name: "COM7", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(s)

	for scanner.Scan(){

		hupla := scanner.Text()
		log.Println(hupla)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}





}

