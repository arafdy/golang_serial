package xxx

import (
	"github.com/tarm/serial"
	"log"
)

func main(){
	c := &serial.Config{Name: "COM6", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
}
