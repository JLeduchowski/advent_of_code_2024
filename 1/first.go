package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func get_distance(x, y int) uint64 {
	return uint64(x - y)
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	pairs := strings.Split(string(file), "\n")
	var sum_distance, similarity_score uint64
	var first_list, second_list []int
	for i := 0; i < len(pairs); i++ {
		current_pair := strings.Split(pairs[i], "   ")
		first, err := strconv.Atoi(current_pair[0])
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.Atoi(current_pair[1])
		if err != nil {
			log.Fatal(err)
		}
		first_list = append(first_list, first)
		second_list = append(second_list, second)

	}

	sort.Ints(first_list)
	sort.Ints(second_list)

	for i := 0; i < len(first_list); i++ {
		if first_list[i] > second_list[i] {
			sum_distance += get_distance(first_list[i], second_list[i])
		} else if first_list[i] < second_list[i] {
			sum_distance += get_distance(second_list[i], first_list[i])
		}
	}

	for i := 0; i < len(first_list); i++ {
		var occurencies int
		for j := 0; j < len(second_list); j++ {
			if second_list[j] == first_list[i] {
				occurencies += 1
			}
		}
		similarity_score += (uint64(first_list[i]) * uint64(occurencies))
	}

	fmt.Printf("Final distance is: %d\n", sum_distance)
	fmt.Printf("Similarity score is: %d\n", similarity_score)
}
