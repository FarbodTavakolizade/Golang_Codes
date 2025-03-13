package main
import "fmt"


type employee struct {
	firstname,lastname string
	age,salary int

}


func main() {
	
emp :=&employee{"reza", "amiri",21,100}

fmt.Println("first name is ",(*emp).firstname)




}