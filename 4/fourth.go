package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func findXmas(input []string) int {
	var validXmas int
	isLeftDiagMas, isRightDiagMas := false, false
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			isLeftDiagMas = isMas(string(input[i][j]), string(input[i-1][j-1]), string(input[i+1][j+1]))
			isRightDiagMas = isMas(string(input[i][j]), string(input[i-1][j+1]), string(input[i+1][j-1]))
			if isLeftDiagMas && isRightDiagMas {
				validXmas++
			}
			isLeftDiagMas, isRightDiagMas = false, false
		}
	}
	return validXmas
}

func isMas(center, left, right string) bool {
	isMasValid := false
	if center == "A" {
		if ((left == "M") && (right == "S")) || ((left == "S") && (right == "M")) {
			isMasValid = true
		}
	}
	return isMasValid
}

func validXmasByRegexp(input string) int {
	var valids int
	r := regexp.MustCompile(`(XMAS)`)
	matches := r.FindAllString(input, -1)
	valids += len(matches)
	r = regexp.MustCompile(`(SAMX)`)
	matches = r.FindAllString(input, -1)
	valids += len(matches)
	return valids
}

func horizontalXmas(input []string) int {
	var validXmas int
	var subString strings.Builder
	for i := 0; i < len(input); i++ {
		subString.Reset()
		for j := 0; j < len(input[i]); j++ {
			subString.WriteByte(input[i][j])
		}
		validXmas += validXmasByRegexp(subString.String())
		subString.Reset()
	}
	return validXmas
}

func verticalXmas(input []string) int {
	var validXmas int
	var subString strings.Builder
	for i := 0; i < len(input); i++ {
		subString.Reset()
		for j := 0; j < len(input[i]); j++ {
			subString.WriteByte(input[j][i])
		}
		validXmas += validXmasByRegexp(subString.String())
		subString.Reset()
	}
	return validXmas
}

func diagonalLeftXmas(input []string) int {
	var validXmas int
	var subString strings.Builder
	for i := 0; i < len(input); i++ {
		subString.Reset()
		for j := 0; i+j < len(input[i]); j++ {
			subString.WriteByte(input[i+j][j])
		}
		validXmas += validXmasByRegexp(subString.String())
		subString.Reset()
		for j := 0; i+j < len(input[i]); j++ {
			if i == 0 {
				continue
			}
			subString.WriteByte(input[j][i+j])
		}
		validXmas += validXmasByRegexp(subString.String())
		subString.Reset()
	}
	return validXmas
}

func diagonalRightXmas(input []string) int {
	var validXmas int
	var subString strings.Builder

	for i := 0; i < len(input); i++ {
		subString.Reset()
		for j := 0; j <= i; j++ {
			subString.WriteByte(input[j][i-j])
		}
		validXmas += validXmasByRegexp(subString.String())
		subString.Reset()
		if i == len(input)-1 {
			continue
		}
		for j := 0; j <= i; j++ {
			subString.WriteByte(input[len(input)-1-j][len(input)-1+j-i])
		}
		validXmas += validXmasByRegexp(subString.String())
		subString.Reset()
	}
	return validXmas
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputDivided := strings.Split(string(input), "\n")
	var xmasSum, realXmasSum int
	xmasSum += horizontalXmas(inputDivided)
	xmasSum += verticalXmas(inputDivided)
	xmasSum += diagonalLeftXmas(inputDivided)
	xmasSum += diagonalRightXmas(inputDivided)
	realXmasSum = findXmas(inputDivided)
	fmt.Println("Total valid XMAS:", xmasSum)
	fmt.Println("Total valid REAL XMAS:", realXmasSum)
}
