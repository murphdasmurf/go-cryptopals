package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const hex_str = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	raw, err := hex.DecodeString(hex_str)
	if err != nil {
		log.Fatal(err)
	}
	b64_str := base64.StdEncoding.EncodeToString(raw)
	fmt.Println(b64_str)

}
