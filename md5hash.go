package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	fmt.Printf("%T\n", hash)
	bytes := []byte("sourav")
	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	fmt.Println(string(hashValue))
	hashSize := hash.Size()
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
	data := []byte("sourav")
	fmt.Printf("%x", md5.Sum(data))
}
