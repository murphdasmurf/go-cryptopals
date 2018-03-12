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

// Returns the integer length of the longer slice.
func max_len(a []byte, b []byte) int {
	if len(a) < len(b) {
		return len(b)
	}
	return len(a)
}

// Determines the shorter slice and repeatedly XORs the longer with it.
func xor(a []byte, b []byte) []byte {
	// Make a slice the length of the longer slice to hold the XORed value.
	xored := make([]byte, max_len(a, b), max_len(a, b))
	if len(a) < len(b) {
		for i := range b {
			// Operate in block the length of the shorter slice.
			if i%len(a) != 0 {
				continue
			} else {
				for n := 0; n < len(a); n++ {
					// Make sure to stay within bounds of the longer slice.
					if (i + n >= len(b)) {
						break
					}
					xored[i+n] = b[i+n] ^ a[n]
				}
			}
		}
	} else { // Must be len(b) <= len(a), so do the above the other way around.
		for i := range a {
			if i%len(b) != 0 {
				continue
			} else {
				for n := 0; n < len(b); n++ {
					if (i + n >= len(a)) {
						break
					}
					xored[i+n] = a[i+n] ^ b[n]
				}
			}
		}
	} 
	return xored
}

// XORs two hex strings.
func hexor(s1, s2 string) string {
	byte1, err := hex.DecodeString(s1)
	if err != nil {
		log.Fatal(err)
	}
	byte2, err := hex.DecodeString(s2)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(xor(byte1, byte2))
}
