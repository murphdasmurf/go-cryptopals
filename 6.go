package main

import (
	"fmt"
	"sort"
	"strings"
	"encoding/base64"
)

func main() {
	ciphertext, err := base64.StdEncoding.DecodeString(b64_ciphertext)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("key: ", get_full_key(ciphertext, guess_key_size(ciphertext, 3, 41)))
	fmt.Println(decrypt(ciphertext, get_full_key(ciphertext, guess_key_size(ciphertext, 3, 41))))
}

// Decrypt a given ciphertext given a key. Returns the plaintext.
func decrypt(ciphertext []byte, key string) string {
  return xor(ciphertext, []byte(key))
}

// The sum of the number of differing bits (where XOR is true).
func hamming(a []byte, b []byte) int {
	return strings.Count(xor_bin(a, b), "1")
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

// Returns the integer length of the shorter slice.
func min_len(a []byte, b []byte) int {
	if len(a) > len(b) {
		return len(b)
	}
	return len(a)
}

// Sum of the number of ASCII characters in the string.
func sum_letters(str string) int {
	sum := 0
	alphabet_array := strings.Split(Alphabet, "")
	for _, letter := range alphabet_array {
		sum += strings.Count(str, letter)
	}
	return sum
}

// Returns a string of the binary representation of the XOR.
func xor_bin(a []byte, b []byte) string {
	return stringToBin(xor(a, b))
}
// Determines the shorter slice and repeatedly XORs the longer with it.
func xor(a []byte, b []byte) string {
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
		// While iterating through this shorter block, if we ever
		// arrive at an out of range indez on the longer, stop.
		for n := 0; n < len(shorter); n++ {
	   		if (i + n >= len(longer)) {
				break
			} else {
				result[i+n] = longer[i+n] ^ shorter[n]
			}
		  } // End iterating through characters in key.
	} // End iterating through ciphertext.
	return string(result)
}

// Compute the average normalized Hamming distance of the first 10 blocks.
// If we don't have enough bytes to average them all once fully, exit with 10.
func normalized_hamming (ciphertext []byte, keysize int) float64 {
  if 10*keysize > len(ciphertext) {
    fmt.Println("Key longer than a tenth of the length of the ciphertext, returning 10.")
    return 10.0
  }
  // Hold the distances, as a float64. Normalize later.
  normalized := 0
	// Compare the first 10 keylength chunks against the previous.
	for i := 0; i < 10; i++ {
	  first_chunk := ciphertext[i*keysize:(i+1)*keysize-1]
	  next_chunk := ciphertext[(i+1)*keysize:(i+2)*keysize-1]
	  normalized += hamming(first_chunk, next_chunk)
	}
	normal_float := float64(normalized)/float64(10*keysize)
	return normal_float
}

// Take some ciphertext and upper and lower bounds of key length to guess, in bytes.
// Returns the most likely key length (smallest normalized Hamming distance between bytes).
func guess_key_size(ciphertext []byte, lower int, upper int) int {
	normalized := make(map[int]float64)
	for i := lower; i < upper + 1; i++ {
		normalized[i] = normalized_hamming(ciphertext, i)
	}
	// Return the 4 most likely lengths.
	sorted := sort_map(normalized, 4)
	plaintext := ""
	key := ""
	// Get the likeliest keys for the likeliest lengths,
	// and compare which is most likely the plaintext.
	for length, _ := range sorted {
	  this_key := get_full_key(ciphertext, length)
	  this_plaintext := decrypt(ciphertext, this_key)
	  if sum_letters(this_plaintext) > sum_letters(plaintext) {
	    plaintext = this_plaintext
	    key = this_key
	  }
	}
	return len(key)
}

// Sort the map of normalized Hamming distances
func sort_map(input map[int]float64, length int) map[int]float64 {
  // Map to hold the sorted results.
  sorted := make(map[int]float64)
  var hamming []float64
  // Read the values into the slice.
  for _, normalized := range input {
    hamming = append(hamming, normalized)
  }
  // Sort the normalized edit distances.
  sort.Float64s(hamming)
  // Trim it to the desired length.
  hamming = hamming[:length-1]
  // For each distance, determine the key length.
  for i := 0; i < length-1; i++ {
    for j, normalized := range input {
      if normalized == hamming[i] {
        // Add to some new sorted map.
        sorted[j] = normalized
      }
    }
  }
  return sorted
}

// Takes an encrypted byte slice as input, tries all non-special characters against
// it, then counts the number of letters and spaces in the resultant string to guess
// if it's the correct plaintext. Returns the guessed key.
func get_key(ciphertext []byte) string {
	// Store the highest number of letters in the string.
	most_letters := 0
	// Store the decryption key.
	key := ""
	// Try all non-specials with XOR.
	alphabet_array := strings.Split(Alphabet, "")
	for _, letter := range alphabet_array {
		xored := string(xor(ciphertext, []byte(letter)))
		sum := sum_letters(xored)
		if sum > most_letters {
			most_letters = sum
			key = letter
		}
	}
	return key
}

// Returns the multi-byte key.
func get_full_key(ciphertext []byte, key_length int) string {
  // Hold the return key.
  key := ""
  // Break the ciphertext into blocks by position according to key_length.
  for i := 0; i < key_length; i++ {
    // Slice to hold the same-position blocks.
    var block []byte
    for j := i; j+key_length < len(ciphertext); j += key_length {
      block = append(block,ciphertext[j])
    }
    // Append the single-block answer to the multi-byte key.
    key += get_key(block)
  }
  return key
}

// Just iterating through ASCII characters doesn't guess correctly because specials
// throw it off. Use a static list of acceptable characters instead.
const Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ .,?!:'"

const b64_ciphertext = "HUIfTQsPAh9PE048GmllH0kcDk4TAQsHThsBFkU2AB4BSWQgVB0dQzNTTmVSBgBHVBwNRU0HBAxTEjwMHghJGgkRTxRMIRpHKwAFHUdZEQQJAGQmB1MANxYGDBoXQR0BUlQwXwAgEwoFR08SSAhFTmU+Fgk4RQYFCBpGB08fWXh+amI2DB0PQQ1IBlUaGwAdQnQEHgFJGgkRAlJ6f0kASDoAGhNJGk9FSA8dDVMEOgFSGQELQRMGAEwxX1NiFQYHCQdUCxdBFBZJeTM1CxsBBQ9GB08dTnhOSCdSBAcMRVhICEEATyBUCHQLHRlJAgAOFlwAUjBpZR9JAgJUAAELB04CEFMBJhAVTQIHAh9PG054MGk2UgoBCVQGBwlTTgIQUwg7EAYFSQ8PEE87ADpfRyscSWQzT1QCEFMaTwUWEXQMBk0PAg4DQ1JMPU4ALwtJDQhOFw0VVB1PDhxFXigLTRkBEgcKVVN4Tk9iBgELR1MdDAAAFwoFHww6Ql5NLgFBIg4cSTRWQWI1Bk9HKn47CE8BGwFTQjcEBx4MThUcDgYHKxpUKhdJGQZZVCFFVwcDBVMHMUV4LAcKQR0JUlk3TwAmHQdJEwATARNFTg5JFwQ5C15NHQYEGk94dzBDADsdHE4UVBUaDE5JTwgHRTkAUmc6AUETCgYAN1xGYlUKDxJTEUgsAA0ABwcXOwlSGQELQQcbE0c9GioWGgwcAgcHSAtPTgsAABY9C1VNCAINGxgXRHgwaWUfSQcJABkRRU8ZAUkDDTUWF01jOgkRTxVJKlZJJwFJHQYADUgRSAsWSR8KIgBSAAxOABoLUlQwW1RiGxpOCEtUYiROCk8gUwY1C1IJCAACEU8QRSxORTBSHQYGTlQJC1lOBAAXRTpCUh0FDxhUZXhzLFtHJ1JbTkoNVDEAQU4bARZFOwsXTRAPRlQYE042WwAuGxoaAk5UHAoAZCYdVBZ0ChQLSQMYVAcXQTwaUy1SBQsTAAAAAAAMCggHRSQJExRJGgkGAAdHMBoqER1JJ0dDFQZFRhsBAlMMIEUHHUkPDxBPH0EzXwArBkkdCFUaDEVHAQANU29lSEBAWk44G09fDXhxTi0RAk4ITlQbCk0LTx4cCjBFeCsGHEETAB1EeFZVIRlFTi4AGAEORU4CEFMXPBwfCBpOAAAdHUMxVVUxUmM9ElARGgZBAg4PAQQzDB4EGhoIFwoKUDFbTCsWBg0OTwEbRSonSARTBDpFFwsPCwIATxNOPBpUKhMdTh5PAUgGQQBPCxYRdG87TQoPD1QbE0s9GkFiFAUXR0cdGgkADwENUwg1DhdNAQsTVBgXVHYaKkg7TgNHTB0DAAA9DgQACjpFX0BJPQAZHB1OeE5PYjYMAg5MFQBFKjoHDAEAcxZSAwZOBREBC0k2HQxiKwYbR0MVBkVUHBZJBwp0DRMDDk5rNhoGACFVVWUeBU4MRREYRVQcFgAdQnQRHU0OCxVUAgsAK05ZLhdJZChWERpFQQALSRwTMRdeTRkcABcbG0M9Gk0jGQwdR1ARGgNFDRtJeSchEVIDBhpBHQlSWTdPBzAXSQ9HTBsJA0UcQUl5bw0KB0oFAkETCgYANlVXKhcbC0sAGgdFUAIOChZJdAsdTR0HDBFDUk43GkcrAAUdRyonBwpOTkJEUyo8RR8USSkOEENSSDdXRSAdDRdLAA0HEAAeHQYRBDYJC00MDxVUZSFQOV1IJwYdB0dXHRwNAA9PGgMKOwtTTSoBDBFPHU54W04mUhoPHgAdHEQAZGU/OjV6RSQMBwcNGA5SaTtfADsXGUJHWREYSQAnSARTBjsIGwNOTgkVHRYANFNLJ1IIThVIHQYKAGQmBwcKLAwRDB0HDxNPAU94Q083UhoaBkcTDRcAAgYCFkU1RQUEBwFBfjwdAChPTikBSR0TTwRIEVIXBgcURTULFk0OBxMYTwFUN0oAIQAQBwkHVGIzQQAGBR8EdCwRCEkHElQcF0w0U05lUggAAwANBxAAHgoGAwkxRRMfDE4DARYbTn8aKmUxCBsURVQfDVlOGwEWRTIXFwwCHUEVHRcAMlVDKRsHSUdMHQMAAC0dCAkcdCIeGAxOazkABEk2HQAjHA1OAFIbBxNJAEhJBxctDBwKSRoOVBwbTj8aQS4dBwlHKjUECQAaBxscEDMNUhkBC0ETBxdULFUAJQAGARFJGk9FVAYGGlMNMRcXTRoBDxNPeG43TQA7HRxJFUVUCQhBFAoNUwctRQYFDE43PT9SUDdJUydcSWRtcwANFVAHAU5TFjtFGgwbCkEYBhlFeFsABRcbAwZOVCYEWgdPYyARNRcGAQwKQRYWUlQwXwAgExoLFAAcARFUBwFOUwImCgcDDU5rIAcXUj0dU2IcBk4TUh0YFUkASEkcC3QIGwMMQkE9SB8AMk9TNlIOCxNUHQZCAAoAHh1FXjYCDBsFABkOBkk7FgALVQROD0EaDwxOSU8dGgI8EVIBAAUEVA5SRjlUQTYbCk5teRsdRVQcDhkDADBFHwhJAQ8XClJBNl4AC1IdBghVEwARABoHCAdFXjwdGEkDCBMHBgAwW1YnUgAaRyonB0VTGgoZUwE7EhxNCAAFVAMXTjwaTSdSEAESUlQNBFJOZU5LXHQMHE0EF0EABh9FeRp5LQdFTkAZREgMU04CEFMcMQQAQ0lkay0ABwcqXwA1FwgFAk4dBkIACA4aB0l0PD1MSQ8PEE87ADtbTmIGDAILAB0cRSo3ABwBRTYKFhROHUETCgZUMVQHYhoGGksABwdJAB0ASTpFNwQcTRoDBBgDUkksGioRHUkKCE5THEVCC08EEgF0BBwJSQoOGkgGADpfADETDU5tBzcJEFMLTx0bAHQJCx8ADRJUDRdMN1RHYgYGTi5jMURFeQEaSRAEOkURDAUCQRkKUmQ5XgBIKwYbQFIRSBVJGgwBGgtzRRNNDwcVWE8BT3hJVCcCSQwGQx9IBE4KTwwdASEXF01jIgQATwZIPRpXKwYKBkdEGwsRTxxDSToGMUlSCQZOFRwKUkQ5VEMnUh0BR0MBGgAAZDwGUwY7CBdNHB5BFwMdUz0aQSwWSQoITlMcRUILTxoCEDUXF01jNw4BTwVBNlRBYhAIGhNMEUgIRU5CRFMkOhwGBAQLTVQOHFkvUkUwF0lkbXkbHUVUBgAcFA0gRQYFCBpBPU8FQSsaVycTAkJHYhsRSQAXABxUFzFFFggICkEDHR1OPxoqER1JDQhNEUgKTkJPDAUAJhwQAg0XQRUBFgArU04lUh0GDlNUGwpOCU9jeTY1HFJARE4xGA4LACxSQTZSDxsJSw1ICFUdBgpTNjUcXk0OAUEDBxtUPRpCLQtFTgBPVB8NSRoKSREKLUUVAklkERgOCwAsUkE2Ug8bCUsNSAhVHQYKUyI7RQUFABoEVA0dWXQaRy1SHgYOVBFIB08XQ0kUCnRvPgwQTgUbGBwAOVREYhAGAQBJEUgETgpPGR8ELUUGBQgaQRIaHEshGk03AQANR1QdBAkAFwAcUwE9AFxNY2QxGA4LACxSQTZSDxsJSw1ICFUdBgpTJjsIF00GAE1ULB1NPRpPLF5JAgJUVAUAAAYKCAFFXjUeDBBOFRwOBgA+T04pC0kDElMdC0VXBgYdFkU2CgtNEAEUVBwTWXhTVG5SGg8eAB0cRSo+AwgKRSANExlJCBQaBAsANU9TKxFJL0dMHRwRTAtPBRwQMAAATQcBFlRlIkw5QwA2GggaR0YBBg5ZTgIcAAw3SVIaAQcVEU8QTyEaYy0fDE4ITlhIJk8DCkkcC3hFMQIEC0EbAVIqCFZBO1IdBgZUVA4QTgUWSR4QJwwRTWM="
