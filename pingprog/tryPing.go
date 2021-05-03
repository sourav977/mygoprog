package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveIPAddr("ip", "google.com")
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("ok1")
	fmt.Println(addr.IP)
	conn, err := net.DialIP("ip4:icmp", nil, addr)
	checkError(err)
	var msg [512]byte
	fmt.Println("ok2")
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum, fix later
	msg[3] = 0  // checksum, fix later
	msg[4] = 0  // identifier[0]
	msg[5] = 13 // identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1]
	len := 8
	fmt.Println("ok3")
	check := checkSum(msg[0:len])
	fmt.Println("ok4")
	msg[2] = byte(check >> 8)
	fmt.Println("ok5")
	msg[3] = byte(check & 255)
	fmt.Println("ok6")
	_, err = conn.Write(msg[0:len])
	checkError(err)
	fmt.Println("ok7")
	_, err = conn.Read(msg[0:])
	checkError(err)
	fmt.Println("ok8")
	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}
	os.Exit(0)
}
func checkSum(msg []byte) uint16 {
	sum := 0
	// assume even for now
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
