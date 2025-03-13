package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := arr[len(arr)/2]

	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr
}

func main() {
	nums := []int{2, 8, 75, 46, 92., 13, 15, 4}
	fmt.Println(nums)            //unsorted array
	fmt.Println(quickSort(nums)) //sort array

}
