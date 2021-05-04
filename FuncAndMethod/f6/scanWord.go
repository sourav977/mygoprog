package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("enter words separated by space")
	//scanner := bufio.NewScanner(os.Stdin)
	for {
		//scanner, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		str := scanner.Text()
		//fmt.Println("input:", scanner)
		fmt.Println("input:", str)
		if str == "exit" || str == "quit" {
			os.Exit(0)
		}
	}
}
