package main

import "fmt"

func main() {

/*fmt.Println("please provide a number between 1 and seven")
var number int 
fmt.Scanln(&number)


switch number {

case 1:
	fmt.Println("today is saturday")
case 2:
	fmt.Println("today is sunday")
case 3:
	fmt.Println("today is monday")
case 4:
	fmt.Println("today is tuesdayy")
case 5:
	fmt.Println("today is wednesday")
case 6:
	fmt.Println("today is thursday")
case 7:
	fmt.Println("today is friday")

default:
	fmt.Println("input numbewr is not approved")
*/


var op string
var a,b int
fmt.Println("enter a and b in a consecutive order : ")
fmt.Scanln(&a,&b)
fmt.Println("enter operator: ")
fmt.Scanln(&op)

switch op{
case "+":
	fmt.Println(a+b)
case  "-":
	fmt.Println(a-b)
case "*":
	fmt.Println(a*b)
case "/":
	fmt.Println(a/b)
default:
	fmt.Println("invalid input")

}








}






