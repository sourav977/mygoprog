package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("sample.env")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file")
	}
	scanner := bufio.NewScanner(file) //NewScanner returns a new Scanner to read from r. The split function defaults to ScanLines.

	var readLine []string //to store values after reading
	//readLine := make([]string, 1)
	for scanner.Scan() {
		readLine = append(readLine, scanner.Text())
	}
	fmt.Println("reading line by line from a file ...")
	for _, v := range readLine {
		if v == "" {
			continue
		}
		fmt.Println(v)
	}

	fmt.Println("\n\ndisplaying values only ...")
	for _, v := range readLine {
		if v == "" {
			continue
		}
		newV := strings.Split(v, "=")
		fmt.Println(newV[1])
	}
}
