package main

import (
	"bytes"
	"fmt"
	"unsafe"

	"golang.org/x/crypto/blowfish"
)

func main() {
	key := []byte("my private key")
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	src := []byte("1234567890")
	fmt.Println("size of src:", unsafe.Sizeof(src))
	var enc [64]byte
	cipher.Encrypt(enc[0:], src)
	var decrypt [64]byte
	cipher.Decrypt(decrypt[0:], enc[0:])
	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:64])
	fmt.Println(string(result.Bytes()))
}
