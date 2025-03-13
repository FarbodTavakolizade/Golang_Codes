package main

import 
(
	"fmt"
    "math"
)

func main() {
	var a,b,c,root1,root2,delta float64

fmt.Println("enter a ,b,c of quadratic equationin a consecutive order: ")
fmt.Scanln(&a, &b, &c)

delta=(b*b)-(4*a*c)

if delta>0 {
    root1=(-b+math.Sqrt(delta))/(2*a)
    root2=(-b-math.Sqrt(delta))/(2*a)
    fmt.Println("two distinct roots exist:root1 =",root1,"and root2 is: ",root2)
}else if delta==0 {
    root1=(-b)/(2*a)
	root2=(-b)/(2*a)
    fmt.Println("two equal roots exist: root1 = root2 =",root1)
}else{
fmt.Println("delta is negative")
}
}
