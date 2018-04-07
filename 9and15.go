package main

import (
	"fmt"
	"bytes"
	"errors"
)

// PKCS7 errors.
var (
	ErrInvalidBlockSize = errors.New("invalid blocksize")
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

func main() {
	// Test with our favorite 16-byte key.
	str := []byte("YELLOW SUBMARINE")
	pad, err := pkcs7Pad(str, 20)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Padded string: ", string(pad))
	unpad, err := pkcs7Unpad(pad, 20)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Remove the padding: ", string(unpad))
}

// Right-pads the given byte slice with 0-len(block)-1
// bytes, per RFC 2315, section 10.3.
func pkcs7Pad(unpadded []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if unpadded == nil || len(unpadded) == 0 || len(unpadded)%blocksize == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	padding := blocksize - (len(unpadded) % blocksize)
	padded := make([]byte, len(unpadded)+padding)
	// Copy the unpadded block contents in, then pad with the
	// integer value of the length of the padding.
	copy(padded, unpadded)
	copy(padded[len(unpadded):], bytes.Repeat([]byte{byte(padding)}, padding))
	return padded, nil
}

// Validates and unpads data from the given bytes slice.
func pkcs7Unpad(padded []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if padded == nil || len(padded) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	if len(padded)%blocksize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}
	// The padding value should be the repeated integer value
	// of the length of the padding.
	padding := padded[len(padded)-1]
	paddingsize := int(padding)
	if paddingsize == 0 || paddingsize > len(padded) {
		return nil, ErrInvalidPKCS7Padding
	}
	// Validates that the padding complies with RFC 2315.
	// That is, if there is a value other than the length of
	// the padding in the padding area, throw an error.
	for i := 0; i < paddingsize; i++ {
		if padded[len(padded)-paddingsize+i] != padding {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return padded[:len(padded)-paddingsize], nil
}
