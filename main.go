package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 8, 9, 3, 2, 1}
	max := FindMax(a)
	fmt.Println(max)
}

func FindMax(a []int) int {
	var res []int
	for i, _ := range a {
		if a[i] == a[len(a)-1-i] {
			res = append(res, a[i])
		} else if a[i] == a[len(a)-1-i] && len(a) != 0 {
			break
		}
	}
	var max int
	for i := 0; i < len(res)/2; i++ {
		for j := i + 1; j < len(res)/2; j++ {
			if res[i] > res[j] && res[i] > max {
				max = res[i]
			} else if res[i] < res[j] && res[j] > max {
				max = res[j]
			}
		}
	}
	return max
}
