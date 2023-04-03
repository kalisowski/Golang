package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"todo/fileOperations"
)

func main() {

	fmt.Print("\033[H\033[2J")
	option := flag.String("option", "", "list, add, remove, done")
	noColor := flag.Bool("nocolor", false, "print output without color")

	flag.Parse()
	if *option == "" {
		fmt.Println("No option provided. Exiting...")
		return
	}
	res := fileOperations.ReadFromFile("todo.txt")
	data := strings.Split(res, "\n")
	switch *option {
	case "list":
		if *noColor {
			list(data, false)
		} else {
			list(data, true)
		}
	case "add":
		add()
	case "remove":
		remove(data)
		// case "done":
		// done()
	case "exit":
		return
	default:
		log.Fatalln("Unknown option: ", *option)
	}
}

func list(res []string, useColor bool) {
	for i, v := range res {
		var color string
		if useColor {
			if v[0] == '#' {
				color = colorCodes.magenta
			} else if v[0] == '[' && v[2] == ']' {
				switch v[1] {
				case 'X':
					color = colorCodes.green
				case ' ':
					color = colorCodes.red
				case '+':
					color = colorCodes.yellow
				case '-':
					color = colorCodes.blue
				}
			}
		}
		fmt.Printf("%d%s %s%s\n", i, color, v, colorCodes.reset)
	}
}

func add() {
	fmt.Println("Podaj notatkę do dodania: ")
	reader := bufio.NewReader(os.Stdin)
	note, _ := reader.ReadString('\n')
	note = strings.TrimSpace(note)
	line := "[ ] " + note
	fileOperations.AppendToFile("todo.txt", line)
}
func remove(res []string) {
	fmt.Println("Podaj numer pozycji do usunięcia: ")
	var index int
	fmt.Scanln(&index)
	println(len(res))
	if index < 0 || index > len(res)-1 {
		log.Fatalln("Podano nieprawidłowy numer linii")
	}
	removed := fileOperations.RemoveLine("todo.txt", index)
	fmt.Println("Usunięto: ", removed)
}

func done() {
	fmt.Println("done")
}

var colorCodes = struct {
	red, green, yellow, blue, magenta, cyan, reset string
}{
	"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m", "\033[0m",
}
