package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(10, 3))
}

func distributeCandies(candies int, num_people int) []int {

	result := make([]int, num_people)
	var index int
	// 显然不是最优解，但是直观
	for i := 0; i < candies; i++ {
		index = i % num_people

		result[index] += i + 1
		candies -= i + 1
		if candies < i+2 {
			result[(index+1)%num_people] += candies
		}
	}
	return result
}
