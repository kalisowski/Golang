package fileOperations

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

func RemoveLine(filename string, index int) string {
	plainText := ReadFromFile(filename)
	data := strings.Split(plainText, "\n")
	removed := data[index]
	data = append(data[:index], data[index+1:]...)
	plainText = strings.Join(data, "\n")
	WriteToFile(filename, plainText)
	return removed
}

func WriteToFile(filename string, text string) error {
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		log.Fatalln("\nError writing to file")
	}

	return nil
}

func AppendToFile(filename string, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString("\n" + text); err != nil {
		return err
	}

	return nil
}
