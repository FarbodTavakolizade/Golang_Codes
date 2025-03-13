package main

import "fmt"

/*
type interface_name interface{

function_name(list of argument types)(list of return types)
}
*/

type bot interface {
	getGreeting() string
}


type englishBot struct {}
type spanishBot struct {}

func (englishBot) getGreeting() string {
	return "hello welcome to bot"
}

func (spanishBot) getGreeting() string {
	return "hola"
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}





func main(){


	englishPart:=englishBot{}
	spanishPart:=spanishBot{}

	printGreeting(spanishPart)
	printGreeting(englishPart)
}