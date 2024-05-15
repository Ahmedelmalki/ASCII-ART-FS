package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("please enter tow args")
	}
	getInput := getInput(os.Args[1])
	readfile := readfile(os.Args[2])
	PrintOutput(getInput, readfile)
}

func getInput(input string) []string {
	var splitedInput []string
	for _, runes := range input {
		if runes < 32 || runes > 126 {
			fmt.Println("invalid input")
			os.Exit(1)
		}
		splitedInput = strings.Split(input, "\\n")
	}
	return splitedInput
}

func readfile(filename string) map[rune][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("someting went wrong")
	}
	defer file.Close()
	output := make(map[rune][]string)
	s := bufio.NewScanner(file)
	i := 31
	line := 0
	for s.Scan() {
		if s.Text() == "" {
			i++
			continue
		}
		// if line == 8 {
		// 	line = 0

		// 	continue
		// }
		output[rune(i)] = append(output[rune(i)], s.Text())
		line++
		if i == 127 {
			break
		}
	}
	return output
}

func PrintOutput(getInput []string, readfile map[rune][]string) {
	is_printed := false
	for idx, word := range getInput {
		if word != "" {
			i := 0
			for i < 8 {
				for _, char := range word {
					line := readfile[char][i]
					fmt.Print(line)
				}
				fmt.Println()
				i++
				is_printed = true
			}
		} else {
			if idx == len(getInput)-1 && !is_printed {
				continue
			} else {
				fmt.Println()
			}
		}
	}
}
