package main

import "fmt"

func main() {

array:=[]int{1,2,3,4}

n:=4

sum:=0

for i:=0; i<n ; i++{
	sum+=array[i]
}
avg:=float64(sum) / float64(n)

fmt.Println("Sum of array elements is:", sum,"and avg of elements is:", avg)








}