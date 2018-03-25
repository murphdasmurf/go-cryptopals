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
