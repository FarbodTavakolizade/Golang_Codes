package main
import "fmt"
 //<-
//var channelName chan type              channelName:=make (chan type)   

func main(){
	x:=make(chan string)

	go func (){// goroutine
		x<-"salam" //sent to cha

	}()
	salam:=<-x//recieve from channel
		fmt.Println(salam)
}

