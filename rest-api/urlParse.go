package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	fmt.Println("before: ", u.String())
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"    //changed from http to https
	u.Host = "google.com" // changed from bing.com to google.com
	q := u.Query()
	q.Set("q", "golang")    //changed query string from dotnet to golang
	u.RawQuery = q.Encode() //overwrite URL
	fmt.Println("after: ", u.String())
}
