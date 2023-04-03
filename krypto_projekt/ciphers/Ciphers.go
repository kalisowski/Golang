package ciphers

import (
	"fmt"
	"io/ioutil"
	"lab01/fileOperations"
	"log"
	"os"
	"strings"
)

func CaesarCipher(text string, shift int, decrypt bool) string {
	var result strings.Builder
	shift = shift % 26
	if shift < 0 {
		shift += 26
	}

	if decrypt {
		shift = 26 - shift
	}

	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			shifted := 'a' + (char-'a'+rune(shift))%26
			result.WriteRune(shifted)
		} else if char >= 'A' && char <= 'Z' {
			shifted := 'A' + (char-'A'+rune(shift))%26
			result.WriteRune(shifted)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func AffineCipher(text string, a int, b int, decrypt bool) string {
	var result strings.Builder
	if a%2 == 0 || a%13 == 0 {
		log.Fatalln("Invalid value for a. Choose a value that is relatively prime to 26.")
	}
	aInverse := -1
	for i := 1; i < 26; i++ {
		if (a*i)%26 == 1 {
			aInverse = i
			break
		}
	}
	if aInverse == -1 {
		log.Fatalln("Invalid value for a. Choose a value that is relatively prime to 26.")
	}
	shift := 0
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			if decrypt {
				shift = aInverse * (int(char-'a') - b)
			} else {
				shift = a*int(char-'a') + b
			}
			shifted := 'a' + (rune(shift)%26+26)%26
			result.WriteRune(shifted)
		} else if char >= 'A' && char <= 'Z' {
			if decrypt {
				shift = aInverse * (int(char-'A') - b)
			} else {
				shift = a*int(char-'A') + b
			}
			shifted := 'A' + (rune(shift)%26+26)%26
			result.WriteRune(shifted)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func BreakCaesarCipherWithExtra(cipherText string, extraText string) {
	distance := int(cipherText[0]) - int(extraText[0])
	plaintext := CaesarCipher(cipherText, distance, true)

	fileOperations.WriteToFile("data/decrypt.txt", plaintext)
	fileOperations.WriteToFile("data/key-found.txt", fmt.Sprintf("%d %d", distance, 0))

	fmt.Println("Decrypted text written to data/decrypt.txt")
	fmt.Println("Key written to data/key-found.txt")
}

func BruteCeasarCipher(ciphertext string) {
	file, err := os.OpenFile("data/decrypt.txt", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for i := 1; i < 26; i++ {
		decryptedText := CaesarCipher(ciphertext, i, true)
		file.WriteString(decryptedText + "\n")
	}
	fmt.Println("Decrypted text written to data/decrypt.txt")
}

func BreakAffineCipherWithExtra(ciphertext string, extraText string) {
	extraLen := len(extraText)
	if extraLen < 2 {
		log.Fatalln("Extra text must be at least 2 characters long")
	}

	for a := 1; a < 26; a++ {
		if a%2 == 0 || a%13 == 0 {
			continue
		}
		aInverse := -1
		for i := 1; i < 26; i++ {
			if (a*i)%26 == 1 {
				aInverse = i
				break
			}
		}
		if aInverse == -1 {
			continue
		}
		for b := 0; b < 26; b++ {
			plaintext := AffineCipher(ciphertext, a, b, true)
			if len(plaintext) < extraLen {
				continue
			}
			if plaintext[:extraLen] != extraText {
				continue
			}
			keyText := fmt.Sprintf("%d %d", a, b)
			err := ioutil.WriteFile("data/key-found.txt", []byte(keyText), 0644)
			if err != nil {
				fmt.Println("Error saving key to file:", err)
				return
			}

			decryptText := AffineCipher(ciphertext, a, b, true)
			err = ioutil.WriteFile("data/decrypt.txt", []byte(decryptText), 0644)
			if err != nil {
				fmt.Println("Error saving decrypted text to file:", err)
				return
			}

			fmt.Println("Decrypted text written to data/decrypt.txt")
			fmt.Println("Key written to data/key-found.txt")
			return
		}
	}

	fmt.Println("Unable to find key.")
}

func BruteAffineCipher(ciphertext string) {
	file, err := os.OpenFile("data/decrypt.txt", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for a := 1; a < 26; a++ {
		if gcd(a, 26) != 1 {
			continue
		}
		for b := 0; b < 26; b++ {
			decryptedText := AffineCipher(ciphertext, a, b, true)
			file.WriteString(decryptedText + "\n")
		}
	}
	fmt.Println("Decrypted text written to data/decrypt.txt")
}

// helper function to calculate the greatest common divisor of two integers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
