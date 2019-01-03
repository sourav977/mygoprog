package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"Person"`
		Id        int      `xml:"Id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"bio>age"`
		Height    float32  `xml:"bio>height"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42, Height: 170.23}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, " ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}
