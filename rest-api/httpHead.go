package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "URL")
		os.Exit(1)
	}
	url := os.Args[1]

	response, err := http.Head(url) //to retrive resource information
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println("req status:", response.Status)
	fmt.Println("http responce protocol:", response.Proto)
	for k, v := range response.Header { //response.Header stores resource information
		fmt.Println(k+":", v)
	}

	//Usually, we are want to retrieve a resource rather than just get information about it.
	//The "GET" request will do this

	os.Exit(0)
}
