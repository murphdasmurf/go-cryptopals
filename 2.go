package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const s1 = "1c0111001f010100061a024b53535009181c"
	const s2 = "686974207468652062756c6c277320657965"
	fmt.Println(hexor(s1, s2))
}

// XORs 2 byte slices, assuming that both arguments are of equal length.
func xor(a []byte, b []byte) []byte {
	if len(a) != len(b) {
		log.Fatal("Slices not equal length!")
	}
	c := make([]byte, len(a)+1)
	for i := range b {
		c[i] = b[i]
	}
	for i := range a {
		c[i] ^= a[i]
	}
	return c
}

// XORs two hex strings using the function above.
func hexor(s1, s2 string) string {
	byte1, err := hex.DecodeString(s1)
	if err != nil {
		log.Fatal(err)
	}
	byte2, err := hex.DecodeString(s2)
	if err != nil {
		log.Fatal(err)
	}
	xor_byte := xor(byte1, byte2)
	xor_hex := hex.EncodeToString(xor_byte)
	return xor_hex
}
