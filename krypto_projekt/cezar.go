// Author: Kamil Lisowski

package main

import (
	"fmt"
	"lab01/ciphers"
	"lab01/fileOperations"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		log.Fatalln("\nInvalid Arguments! Usage---> cypher method: [-c|-a] action: [-e|-d|-j|-k]")
	}

	args := os.Args[1:]
	method := args[0]
	action := args[1]

	if method != "-c" && method != "-a" {
		fmt.Println("Unknown method - methods: [-c|-a]")
		return
	}

	if action != "-e" && action != "-d" && action != "-j" && action != "-k" {
		fmt.Println("Unknown action - actions: [-e|-d|-j|-k]")
		return
	}

	if method == "-c" {
		if action == "-e" {
			a, _ := fileOperations.ReadKey(method)
			plainText := fileOperations.ReadFromFile("data/plain.txt")
			cipherText := ciphers.CaesarCipher(plainText, a, false)
			fileOperations.WriteToFile("data/crypto.txt", cipherText)
			fmt.Println("Encryption succeeded")
			return
		}
		if action == "-d" {
			a, _ := fileOperations.ReadKey(method)
			encryptedText := fileOperations.ReadFromFile("data/crypto.txt")
			cipherText := ciphers.CaesarCipher(encryptedText, a, true)
			fileOperations.WriteToFile("data/decrypt.txt", cipherText)
			fmt.Println("Decryption succeeded")
			return
		}
		if action == "-j" {
			encryptedText := fileOperations.ReadFromFile("data/crypto.txt")
			extraText := fileOperations.ReadFromFile("data/extra.txt")
			ciphers.BreakCaesarCipherWithExtra(encryptedText, extraText)
		}
		if action == "-k" {
			encryptedText := fileOperations.ReadFromFile("data/crypto.txt")
			ciphers.BruteCeasarCipher(encryptedText)
		}
	}

	if method == "-a" {
		if action == "-j" {
			encryptedText := fileOperations.ReadFromFile("data/crypto.txt")
			extraText := fileOperations.ReadFromFile("data/extra.txt")
			ciphers.BreakAffineCipherWithExtra(encryptedText, extraText)
		}
		if action == "-k" {
			encryptedText := fileOperations.ReadFromFile("data/crypto.txt")
			ciphers.BruteAffineCipher(encryptedText)
		}
		if action == "-e" {
			a, b := fileOperations.ReadKey(method)
			plainText := fileOperations.ReadFromFile("data/plain.txt")
			cipherText := ciphers.AffineCipher(plainText, a, b, false)
			fileOperations.WriteToFile("data/crypto.txt", cipherText)
			fmt.Println("Encryption succeeded")
			return
		}
		if action == "-d" {
			a, b := fileOperations.ReadKey(method)
			encryptedText := fileOperations.ReadFromFile("data/crypto.txt")
			cipherText := ciphers.AffineCipher(encryptedText, a, b, true)
			fileOperations.WriteToFile("data/decrypt.txt", cipherText)
			fmt.Println("Decryption succeeded")
			return
		}

	}
}
