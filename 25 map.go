package main

import "fmt"

func main() {

// var name =map[key type] value type{key1 : value1, key2: value2,...key n :value n}
//name:=map[key type] value type{key1 : value1, key2: value2,...key n :value n}

/*
a:=map[string]int{"number one":1, "number two":2, "number three":3}

fmt.Printf("a:\t %v",a)
*/
// name :=make(map[keytype]value type)



var name =map[int]string{1:"farbod",2:"reza",3:"amir"}

b:=make(map[string]int)

b["name"]=1
b["number"]=2
//use delete function for removing in map 
delete(b, "name")



fmt.Printf("name:\t %v",name)

fmt.Printf("b:\t %v",b)






}