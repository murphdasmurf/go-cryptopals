package main

import (
	"fmt"
	"encoding/hex"
)

func main() {
	const s1 = "Burning 'em, if you ain't quick and nimble"
	const s2 = "I go crazy when I hear a cymbal"
	const key = "ICE"
	fmt.Println(hex.EncodeToString([]byte(key)))
	fmt.Println(hex.EncodeToString([]byte(s1)))
	fmt.Println(hex.EncodeToString(xor([]byte(s1), []byte(key))))
	fmt.Println(hex.EncodeToString([]byte(s2)))
	fmt.Println(hex.EncodeToString(xor([]byte(s2), []byte(key))))
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
			if i%len(a) == 0 {
				for n := 0; n < len(a); n++ {
					if (i + n < len(b)) {
						xored[i+n] = b[i+n] ^ a[n]
					} else { continue }
				}
			} else { continue }
		}
	} else { // Must be len(b) <= len(a), so do the above the other way around.
		for i := range a {
			if i%len(b) == 0 {
				for n := 0; n < len(b); n++ {
					if (i + n < len(a)) {
						xored[i+n] = a[i+n] ^ b[n]
					} else {
						continue
					}
				}
			} else { continue }
		}
	} 
	return xored
}
