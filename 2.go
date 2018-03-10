package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const s1 = "1c0111001f010100061a024b53535009181c"
	const s2 = "686974207468652062756c6c277320657965"
	byte1, err := hex.DecodeString(s1)
	if err != nil {
		log.Fatal(err)
	}
	byte2, err := hex.DecodeString(s2)
	if err != nil {
		log.Fatal(err)
	}
	xor_byte := xor(byte1, byte2)
	xor_str := hex.EncodeToString(xor_byte)
	fmt.Printf("%s\n", xor_str)
}

// XORs 2 byte arrays, assuming that both arguments are of equal length.
// TODO: check this assumption and throw error if not true.
// Consider accepting an optional third argument for position within the
// longer array to XOR the shorter array against.
func xor(a []byte, b []byte) []byte {
	c := make([]byte, len(a)+1)
	for i := range b {
		c[i] = b[i]
	}
	for i := range a {
		c[i] ^= a[i]
	}
	return c
}
