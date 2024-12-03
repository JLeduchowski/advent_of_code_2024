package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputString := string(input)
	// inputStringTest := inputString
	// splittedString := strings.Split(inputStringTest, "don't()")
	// var validMuls strings.Builder
	// validMuls.WriteString(splittedString[0])
	// for i := 0; i < len(splittedString); i++ {
	// 	splittedDos := splittedString[i]
	// 	if len(splittedDos) > 1{
	// 		for j := 1; j < len(splittedDos); j++ {
	// 			validMuls.WriteString(string(splittedDos[j]))
	// 		}
	// 	}
	// }
	var sum int
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllString(inputString, -1)
	for i := 0; i < len(matches); i++ {
		match := matches[i]
		splitted := strings.Split(match, ",")
		first, second := splitted[0], splitted[1]
		first = strings.Replace(first, "mul(", "", 1)
		second = strings.Replace(second, ")", "", 1)
		firstNumber, err := strconv.Atoi(first)
		if err != nil {
			log.Fatal(err)
		}
		secondNumber, err := strconv.Atoi(second)
		if err != nil {
			log.Fatal(err)
		}
		sum += firstNumber * secondNumber
	}
	fmt.Printf("Sum of valid multiplies: %d\n", sum)
}
