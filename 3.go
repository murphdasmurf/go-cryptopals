package main

import (
	"encoding/hex"
	"strings"
	"fmt"
	"log"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	// Convert the given hex string to a byte slice.
	const s1 = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	byte1, err := hex.DecodeString(s1)
	if err != nil {
		log.Fatal(err)
	}
	// Split up the alphabet to try with XOR.
	letter_array := strings.Split(Alphabet, "")
	for _, letter := range letter_array {
		byte2 := []byte(letter)
		xored_slice := xor(byte1, byte2)
		fmt.Println(frequency(string(xored_slice)))
	}
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
					xored[i+n] = a[i+n] ^ b[n]
				}
			}
		}
	} 
	return xored
}

// Converts everything to lowercase to measure the frequency of each letter.
func frequency(a string) []int {
	lower_case := strings.Split(Alphabet[:26], "")
	freq := make([]int, len(lower_case), cap(lower_case))
	for i, letter := range lower_case {
		freq[i] = strings.Count(strings.ToLower(a), letter)
	}
	return freq
}
