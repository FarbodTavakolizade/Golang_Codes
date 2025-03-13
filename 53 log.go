package main

import(
    "log"
    "os"
)
func main(){

    file ,err:=os.OpenFile("app.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
    if  err!=nil{
        log.Fatalln("Error opening log file ")
    }

    log.SetOutput(file)

    log.Println("this is a log message.")
    log.Printf("logging a new variable%s\t","reza")
    log.Fatalf("this  is a fatal log massage")
    

}