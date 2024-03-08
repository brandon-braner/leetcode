package main

import "fmt"

func main() {
	result := canPlaceFlowers([]int{0, 0, 0, 0, 0, 1, 0, 0}, 0)
	fmt.Println(result)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i := range flowerbed {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
		if n == 0 {
			return true
		}
	}
	return false
}
