/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanln(&N)
	A := make([]int, N)
	scanner := bufio.NewScanner(os.Stdin)
	str := readString(scanner)
	newstr := strings.Split(str, " ")
	for i := 0; i < N; i++ {
		s := newstr[i]
		A[i], _ = strconv.Atoi(s)
	}
	fmt.Println("array:", A)
	min := A[0]
	for _, v := range A {
		if v < min {
			min = v
		}
	}
	fmt.Println("Closest to Zero is:", min)

}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
