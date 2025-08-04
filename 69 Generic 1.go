// after version 1.18 go     Generic
package main

import "fmt"

func main(){
  PrintState1(5)
  PrintState2(5.25)
  PrintState3("jns")
  PrintState4(true)
  
  fmt.Println("-------------------")
  
  PrintStateGeneric(78)
  PrintStateGeneric(15.86)
  PrintStateGeneric(false)
  PrintStateGeneric("sndjnd")
  
  PrintStateGenericList([]int{1,2,3})
  PrintStateGenericList([]string{"abc","def","djkd"})
}

func PrintStateGenericList[T any](list[] T){
  for _ ,x := range list{
      fmt.Println(x) 
    }
  }



func PrintStateGeneric[T any](item T){
  fmt.Println(item)
}

func PrintState1(a int){
  fmt.Println(a)
}
func PrintState2(b float64){
  fmt.Println(b)
}

func PrintState3(c string){
  fmt.Println(c)
}

func PrintState4(d bool){
  fmt.Println(d)
}
