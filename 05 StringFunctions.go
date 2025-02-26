package main

import (
	"fmt"
	"strings"
)

func main() {

// strings.Contains           [x is a string]    strings.contain(x,"phrase")   check if x contains phrase

Sentence:="hello i am learning golang golang "

fmt.Println(strings.Contains(Sentence, "reza"))


//strings.cut         [x is a string string]    strings.cut(x,"phrase")   delete phrase from x string 

fmt.Println(strings.Cut(Sentence, "reza"))



//strings.Count           [x is a string]  strings.count(x,"phrase")   count the number of times phrase appears in x string
fmt.Println(strings.Count(Sentence, "golang"))


//len    length of arguments    len(args)

fmt.Println(len(Sentence))


//strings.repeat       [x is a string]      strings.repeat(x,"phrase") 
fmt.Println(strings.Repeat("name ",10))


//strings.replace    [x is a string]     strings.replace(x,"old phrase","new phrase ",number of replacement )

fmt.Println(strings.Replace(Sentence, "hello","salam ",1))
fmt.Println(strings.Replace(Sentence,"golang","go",1))

//strings.replaceAll        [x is a string]    strings.replaceAll(x,"old phrase","new phrase")

fmt.Println(strings.ReplaceAll(Sentence,"golang","go"))


//strings.compare       [x is a string]     strings.compare("phrase 1","phrase 2")   for key sensitive comparison


fmt.Println(strings.Compare("ali","reza"))
fmt.Println(strings.Compare("reza","reza"))
fmt.Println(strings.Compare("ali","Reza"))


//strings.equalfold   [x is a string] strings.equalfold("phrase 1","phrase2")   key sensitive is not important

fmt.Println(strings.EqualFold("farbod","FARBOD"))


//strings.Index     strings.index("phrase "," chosen character ")   return chosen character index

fmt.Println(strings.Index("farbod","a"))




//upper case  & lower case  & title




fmt.Println(strings.ToUpper("farbod"))
fmt.Println(strings.ToLower("FARBOD"))
fmt.Println(strings.Title("hello i live in iran"))





//strings.trim      [x is a string]    strings.trim("phrase "," chosen part ") remove chosen part

fmt.Println(strings.Trim("   golang is good  "," "))
fmt.Println(strings.TrimLeft("IR626265462265"," IR"))
fmt.Println(strings.TrimRight("626265462265IR"," IR"))




//strings.split   [x is a string] strings.split(x,sptr)

fmt.Println(strings.Split(Sentence," "))


fmt.Println(len(strings.Split(Sentence," "))-1)


}