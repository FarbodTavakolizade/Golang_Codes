package main
import "fmt"
func main() {
/*
	prices := []int{10,20,30,40}
	prices[2]= 50
	fmt.Println(prices)
	fmt.Println(len(prices))
	fmt.Println(cap(prices))

	//slice_name = append(slice_name,elements)

    prices = append(prices,20,21)
    fmt.Println(prices)
    fmt.Println(len(prices))
    fmt.Println(cap(prices))


*/
/*
slice:=[]int{1,2,3,4,5}

index:=2

slice = append(slice[:index] , slice[index+1:]...)
fmt.Println(slice)

*/

numbers:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

numbers2 :=numbers[:len(numbers)-10]

numbers_copy := make ([]int ,len(numbers2))
copy(numbers_copy,numbers2)
fmt.Println(numbers_copy)














}