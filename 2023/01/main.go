package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numberWords := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4", "five": "5",
		"six": "6", "seven": "7", "eight": "8", "nine": "9",
	}

	total := 0
	check := 0
	for scanner.Scan() {
		line := scanner.Text()

		line = replaceWords(line, numberWords)
		fmt.Println("=> Updated:", line)

		runes := []rune(line)
		check += 1
		fmt.Println(runes)
		forwardStopped, backwardStopped := false, false
		forwardDigit, backwardDigit := "", ""
		for i, j := 0, len(runes)-1; i <= j; {
			if !forwardStopped {
				if unicode.IsDigit(runes[i]) {
					forwardStopped = true
					forwardDigit = string(runes[i])
				} else {
					i++
				}
			}

			if !backwardStopped {
				if unicode.IsDigit(runes[j]) {
					backwardStopped = true
					backwardDigit = string(runes[j])

				} else {
					j--
				}
			}

			if forwardStopped && backwardStopped {
				combinedNumber, _ := strconv.Atoi(forwardDigit + backwardDigit)
				total += combinedNumber
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total:", total, "Lines Processed:", check)
}

func replaceWords(input string, numberWords map[string]string) string {
	for word, num := range numberWords {
		re := regexp.MustCompile(word)
		input = re.ReplaceAllStringFunc(input, func(match string) string {
			// Replace the second character of the match
			return match[:1] + num + match[2:]
		})
	}
	return input
}
