package main

import (
	"fmt"
	"strings"
)

func main() {
	// Challenge-provided test strings.
	s1 := "this is a test"
	s2 := "wokka wokka!!!"
	
	// Make sure the Hamming distance comes out to 37.
	fmt.Println(hamming(s1, s2))
	
}

// The sum of the number of differing bits (where XOR is true)
// is called the Hamming distance.
func hamming(a string, b string) int {
	return strings.Count(xor([]byte(a), []byte(b)), "1")
}

// Returns a string of the binary representation of the ASCII string.
func stringToBin(s string) (binString string) {
    for _, c := range s {
        binString = fmt.Sprintf("%s%b",binString, c)
    }
    return 
}

// Returns the integer length of the longer slice.
func max_len(a []byte, b []byte) int {
	if len(a) < len(b) {
		return len(b)
	}
	return len(a)
}

// Determines the shorter slice and repeatedly XORs the longer with it.
func xor(a []byte, b []byte) string {
	// Make a slice the length of the longer slice to hold the XORed value.
	xored := make([]byte, max_len(a, b), max_len(a, b))
	if len(a) < len(b) {
		for i := range b {
			// Operate in block the length of the shorter slice.
			if i%len(a) != 0 {
				continue
			} else {
				for n := 0; n < len(a); n++ {
					// Make sure we don't go out of bounds of the longer slice.
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
	return stringToBin(string(xored))
}
