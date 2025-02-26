package main

import "fmt"

func main(){

for i:=10; i<=100;  i+=10{

	fmt.Println(i)
}

//continue      break      in loops


for i:=1; i<=5 ; i++{
	if i==3{
		continue
	}
	fmt.Println(i)

}


for i:=1 ; i<=5 ; i++{

if i==3{
	break
}
fmt.Println(i)

}


sum:=0

for i:=1; i<=100; i++{
	if i==50{
		break
	}

    sum+=i

	
}

fmt.Println(sum)



















}