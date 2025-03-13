package main

import"fmt"


type Person struct {
	name string
	age int 
	salary int
	job string
}


func main(){
var pers1 Person


pers1.name = "reza"
pers1.age =25
pers1.salary =1000
pers1.job = "mechanic"

fmt.Println("fisrt person Name: ", pers1.name)
fmt.Println("fisrt person age: ", pers1.age)
fmt.Println("fisrt person salary: ", pers1.salary)
fmt.Println("fisrt person job: ", pers1.job)
fmt.Println(".........................................")

printPersonData(pers1)




}


func printPersonData(pers Person){

	fmt.Println(" person Name: ", pers.name)
	fmt.Println(" person age: ", pers.age)
	fmt.Println(" person salary: ", pers.salary)
	fmt.Println(" person job: ", pers.job)
	}
	
	