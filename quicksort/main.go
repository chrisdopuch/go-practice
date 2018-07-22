package main

import "fmt"

func main() {
	input := []int{8, 3, 6, 2, 1, 9, 4, 7, 5}
	output := sort(input)
	fmt.Println(output)
}

func sort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	pivot := len(input) / 2
	var right []int
	var left []int
	for i, num := range input {
		if i == pivot {
			continue
		}
		if num > input[pivot] {
			right = append(right, num)
		} else if num <= input[pivot] {
			left = append(left, num)
		}
	}
	return append(append(sort(left), input[pivot]), sort(right)...)
}
