package xxx

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
)

func main()  {
	c := &serial.Config{Name: "COM6", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	var input string
	fmt.Scanln(&input)
	_, err = s.Write([]byte(input))
	if err != nil {
		log.Fatal(err)
}
