package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func copySlice(src []string) []string {
	cpy := make([]string, len(src))
	copy(cpy, src)
	return cpy
}

func isReportSafe(report []string) bool {
	increasing, decreasing := false, false
	isCorrect := true
	for j := 0; j < len(report)-1; j++ {
		first, err := strconv.Atoi(report[j])
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.Atoi(report[j+1])
		if err != nil {
			log.Fatal(err)
		}
		difference := int(math.Abs(float64(first) - float64(second)))
		if difference < 1 || difference > 3 {
			isCorrect = false
		} else if first > second && increasing {
			isCorrect = false
		} else if first < second && decreasing {
			isCorrect = false
		} else if !increasing && !decreasing {
			if first > second {
				decreasing = true
			} else if first < second {
				increasing = true
			} else if first == second {
				isCorrect = false

			}
		}
		if !isCorrect {
			break
		}
	}
	return isCorrect
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var reports = strings.Split(string(input), "\n")
	var safe_reports int
	var isSafe bool
	for i := 0; i < len(reports); i++ {
		report := strings.Split(reports[i], " ")
		isSafe = isReportSafe(report)
		if !isSafe {
			for j := 0; j < len(report); j++ {
				faultyless := append(copySlice(report[:j]), report[j+1:]...)
				isSafe = isReportSafe(faultyless)
				if isSafe {
					break
				}
			}
		}
		if isSafe {
			safe_reports++
		}
	}
	fmt.Printf("Safe reports: %d\n", safe_reports)
}
