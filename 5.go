package main

import (
	"fmt"
	"encoding/hex"
)

func main() {
	const s1 = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	const key = "ICE"
	fmt.Println(hex.EncodeToString(xor([]byte(s1), []byte(key))))
}

// Returns the integer length of the longer slice.
func max_len(a []byte, b []byte) int {
	if len(a) < len(b) {
		return len(b)
	}
	return len(a)
}

// Returns the integer length of the shorter slice.
func min_len(a []byte, b []byte) int {
	if len(a) > len(b) {
		return len(b)
	}
	return len(a)
}

// Determines the shorter slice and repeatedly XORs the longer with it.
func xor(a []byte, b []byte) []byte {
	// Make a slice the length of the longer slice to hold the XORed value.
	result := make([]byte, max_len(a, b), max_len(a, b))
	longer := make([]byte, max_len(a, b), max_len(a, b))
	shorter := make([]byte, min_len(a, b), min_len(a, b))
	if len(a) < len(b) {
	  copy(shorter, a)
	  copy(longer, b)
	} else {
	  copy(shorter, b)
	  copy(longer, a)
	}
	// Iterate along in blocks of length shorter.
	for i := 0; i < len(longer); i += len(shorter) {
		for n := 0; n < len(shorter); n++ {
	   		if (i + n >= len(longer)) {
			  	//fmt.Printf("XOR: I'm out! %d >= %d \n", i+n, len(longer))
				  break
			  } else {
				  result[i+n] = longer[i+n] ^ shorter[n]
			  }
		  } // End iterating through characters in key.
	} // End iterating through ciphertext.
	return result
}
