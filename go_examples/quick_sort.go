package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		var less []int
		var greater []int
		for _, value := range arr[1:] {
			if value <= pivot {
				less = append(less, value)
			} else if value > pivot {
				greater = append(greater, value)
			}
		}
		arr = append(quickSort(less), pivot)
		arr = append(arr, quickSort(greater)...)
		return arr
	}
}

func main() {
	arr := []int{10, 5, 5, 2, 3, 11, -11, 11}

	fmt.Println(quickSort(arr))
}
