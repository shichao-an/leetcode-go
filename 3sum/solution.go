package main

import "fmt"
import "sort"

func threeSum(num []int) [][]int {
	n := len(num)
	sort.Ints(num)
	res := [][]int{}
	for i := 0; i < n-2; i++ {
		if i == 0 || i > 0 && num[i-1] != num[i] {
			left := i + 1
			right := n - 1
			for left < right {
				s := num[i] + num[left] + num[right]
				if s == 0 {
					res = append(res, []int{num[i], num[left], num[right]})
					left++
					right--
					for left < right && num[left] == num[left-1] {
						left++
					}
					for right > left && num[right] == num[right+1] {
						right--
					}
				} else if s < 0 {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}

func main() {
	var num []int = []int{-1, -3, 0, 1, 3, 5, 2}
	fmt.Println(num)
	fmt.Println(threeSum(num))
}
