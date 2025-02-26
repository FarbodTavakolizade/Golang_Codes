package main

import (
    "fmt"

)

func bubblesort(arr[]int){
n:=len(arr)
        for i:=0; i<n-1; i++{
            for j:=0; j<n-i-1; j++{
                if arr[j] > arr[j+1]{
                    arr[j], arr[j+1] = arr[j+1], arr[j]
                }
            }
        }
    }
func main() {
	nums := []int{65,34,25,12,22,87,11}

fmt.Println("unsorted array: ",nums)

bubblesort(nums)

fmt.Println("sorted array: ",nums)



}