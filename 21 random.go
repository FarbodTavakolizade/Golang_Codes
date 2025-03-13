package main
import (
	"fmt"
	"time"
	"math/rand"
)

//seed
func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber:=rand.Intn(100)+1
	fmt.Println(randomNumber)

}