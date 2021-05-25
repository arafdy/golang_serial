package main

import (
	"github.com/tarm/serial"
	"time"

	//"fmt"
	"log"
)

func serialWrite()  {
	c := &serial.Config{Name: "COM6", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}
}

func serialRead()  {

	c := &serial.Config{Name: "COM6", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])


}

func main(){
	for{
		serialWrite()
		time.Sleep(10)
	}

}