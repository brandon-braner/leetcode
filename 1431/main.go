package main

import (
	"fmt"
	"slices"
)

func main() {
	result := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	fmt.Println(result)
}


func kidsWithCandies(candies []int, extraCandies int) []bool {
    // get max number from array
	maxCandy := slices.Max(candies)

	var candiesBools []bool
	for _, candy := range candies {
		if candy+extraCandies >= maxCandy {
			candiesBools = append(candiesBools, true)
		} else {
			candiesBools = append(candiesBools, false)
		}
	}
	return candiesBools
}