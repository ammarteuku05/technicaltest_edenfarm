package main

import "fmt"

func main() {
	var a = []int{11, 4, 7, 5}

	max := FindMax(a)
	fmt.Println("Max : ", max)
}

func FindMax(a []int) int {
	max := a[0]

	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
