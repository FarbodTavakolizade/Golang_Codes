package main
import(
	"fmt"
	"os/exec"
	"strings"
)

func main(){
	target:="192.168.1.1"

	cmd:= exec.Command("nmap",target)

	output ,err:=cmd.CombinedOutput()  

	if err!=nil{
		fmt.Println("error:",err)
		return 
	}
	fmt.Println("results: ")
	fmt.Println(strings.TrimSpace(string (output)))
}