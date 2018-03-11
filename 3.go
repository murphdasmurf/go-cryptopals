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
	// Split up the letters to try with XOR.
	ascii_array := strings.Split(Alphabet, "")
	for _, letter := range ascii_array {
		byte2 := []byte(letter)
		xored_slice := xor(byte1, byte2)
		fmt.Printf("%s\n", string(xored_slice[:]))
	}
}

// Returns the integer length of the longer slice.
func max_len(a []byte, b []byte) int {
	if len(a) < len(b) {
		fmt.Println(len(b))
		return len(b)
	}
	return len(a)
}

// Determines the shorter slice (assumed 1 byte) and XORs the longer with it.
func xor(a []byte, b []byte) []byte {
	// Make the slice to hold the XOR the length of the longer slice.
	xored := make([]byte, max_len(a, b), max_len(a, b))
	if len(a) < len(b) {
		// Copy in the value of the longer slice XORed with the single byte other value.
		for i := range b {
        		xored[i] = b[i] ^ a[0]
		}
	}
	// Do it all again in reverse!
	if len(b) < len(a) {
		for i := range a {
        		xored[i] = a[i] ^ b[0]
		}
	}
	// If equal, both must be one byte.
	if len(a) == len(b) {
		xored[0] = a[0] ^ b[0]
	}
	return xored
}
