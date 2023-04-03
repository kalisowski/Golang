package fileOperations

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func WriteToFile(filename string, text string) error {
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		log.Fatalln("\nError writing to file")
	}

	return nil
}

func ReadFromFile(filename string) string {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("\nError opening file '%s': %s", filename, err)
	}
	defer file.Close()
	plainBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("\nError reading file '%s': %s", filename, err)
	}
	if len(plainBytes) == 0 {
		log.Fatalf("\nFile '%s' is empty", filename)
	}
	plainText := strings.TrimSpace(string(plainBytes))

	return plainText
}

func ReadKey(method string) (int, int) {
	keyFile, err := os.OpenFile("data/key.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening key file")
		os.Exit(1)
	}
	defer keyFile.Close()

	keyBytes, err := ioutil.ReadAll(keyFile)
	if err != nil {
		fmt.Println("Error reading key file")
		os.Exit(1)
	}

	keyStrings := strings.Fields(string(keyBytes))
	if len(keyStrings) != 2 {
		fmt.Println("Key file does not contain two integers")
		os.Exit(1)
	}

	key1, err := strconv.Atoi(keyStrings[0])
	if err != nil {
		fmt.Println("Key file does not contain two integers")
		os.Exit(1)
	}

	key2, err := strconv.Atoi(keyStrings[1])
	if err != nil {
		fmt.Println("Key file does not contain two integers")
		os.Exit(1)
	}

	return key1, key2
}

func WriteAndReplaceToFile(filename string, text string) error {
	file, err := os.OpenFile("data/decrypt.txt", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	return nil
}
