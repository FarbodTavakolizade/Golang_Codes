package main

import "fmt"


func sum (a, b int) int {
result := a + b
fmt.Println("result is ", result)
return 0
}
// defer is lifo(last in first out) filo

func main() {
fmt.Println("hello world")

defer sum(5,10)
defer sum(5,7)




}