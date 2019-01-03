package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// Person struct
type Person struct {
	Name  Name
	Email []Email
}

// Name struct
type Name struct {
	Family   string
	Personal string
}

//Email struct
type Email struct {
	Kind    string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	service := "127.0.0.1:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)
		for n := 0; n < 2; n++ {
			var person Person
			decoder.Decode(&person)
			fmt.Println(person.String())
			person.Name.Family = "sourav"
			person.Name.Personal = "july"
			//person.Name.Email[0].Address = "spatnaik@rythmos.com"
			encoder.Encode(person)
		}
		conn.Close() // we're finished
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
