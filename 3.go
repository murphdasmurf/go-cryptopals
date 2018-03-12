package main

import (
	"encoding/hex"
	"strings"
	"fmt"
	"log"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "

func main() {
	const s1 = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	key := get_key(s1)
	fmt.Println(key)
}

// Returns the integer length of the longer slice.
func max_len(a []byte, b []byte) int {
	if len(a) < len(b) {
		return len(b)
	}
	return len(a)
}

// Sum of the number of alphabets in the string.
func sum_letters(str string) int {
	sum := 0
	for _, letter := range Alphabet {
		sum += strings.Count(string(str), string(letter))
	}
	return sum
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
	return xored
}

// Takes an encrypted hex string as input, tries every English letter against it,
// then counts the number of letters and spaces in the resultant string to guess
// if it's the correct plaintext. Returns the letter used arrive at the guessed answer.
func get_key(a string) string {
	byte1, err := hex.DecodeString(a)
	if err != nil {
		log.Fatal(err)
	}
	// Store the highest number of letters in the string.
	most_letters := 0
	// Store the character that was the correct XOR.
	correct_key := ""
	// Store the decrypted sentence.
	plaintext := ""
	// Split up the alphabet to try with XOR.
	alphabet_array := strings.Split(Alphabet, "")
	for _, letter := range alphabet_array {
		byte2 := []byte(letter)
		xored_slice := xor(byte1, byte2)
		sum := sum_letters(string(xored_slice))
		if sum > most_letters {
			most_letters = sum
			correct_key = letter
			plaintext = string(xored_slice)
		}
	}
	fmt.Println(plaintext)
	return correct_key
}
