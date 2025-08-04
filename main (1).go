
package main

import "fmt"

type Box[T any] struct{
  value T
}
//constraint generic
type Number interface{
  ~int | ~float64
  
}

type Calculator[T Number] struct{
  a, b  T
}
//method
func (c Calculator[T]) sum() T{
  return c.a + c.b
}

func (b Box[T]) Get() T{
  return b.value
}


//multi type generic

func Combine[A any , B any](a A , b B) string{
  return fmt.Sprintf("%v - %v",a,b)
}

func main(){

  intBox :=Box[int]{value:42}
  fmt.Println(intBox.Get())
  
  strBox :=Box[string]{value:"hello world"}
    fmt.Println(strBox.Get())
    
    
    c1:= Calculator[int]{a: 10, b: 20}
    fmt.Println(c1.sum())
    
    fmt.Println(Combine(10 ,"salam")) //10 - salam
}
 

