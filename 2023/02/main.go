package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// limits := map[string]int{
	// 	"blue": 14, "red": 12, "green": 13,
	// }

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ";", ",")
		// gameNumber := strings.Split(strings.Split(line, ":")[0], " ")[1]
		// gameNum, _ := strconv.Atoi(gameNumber)
		diceSegments := strings.Split(line, ":")[1:]
		for _, segment := range diceSegments {
			redMax := 0
			greenMax := 0
			blueMax := 0
			power := 1
			numColorPairs := strings.Split(segment, ",")
			for _, pair := range numColorPairs {
				pair = strings.TrimSpace(pair)
				numAndColor := strings.Fields(pair)
				if len(numAndColor) == 2 {
					color := numAndColor[1]
					num, _ := strconv.Atoi(numAndColor[0])
					switch color {
					case "red":
						if num > redMax {
							redMax = num
						}
					case "green":
						if num > greenMax {
							greenMax = num
						}
					case "blue":
						if num > blueMax {
							blueMax = num
						}
					}

				}
			}
			power = redMax * greenMax * blueMax
			total += power
		}
	}
	fmt.Println(total)
}
