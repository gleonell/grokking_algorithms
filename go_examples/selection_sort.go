package main

import "fmt"

func selectionSort(arr []int) []int{

	for i := 0; i < len(arr); i++{
		for j := i + 1; j < len(arr); j++{
			if arr[j] < arr[i]{
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

func main(){
	var input_arr []int = []int{1, 3, 9, 2, 7, 12, 324, 2, 54 , -54}
	fmt.Println(selectionSort(input_arr))
}