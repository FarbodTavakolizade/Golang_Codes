//ifconfig .... down
//ifconfig .... hw ether xx:xx:xx:xx:xx:xx
//ifconfig .... up

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCommand(cmd string, args []string, password string) error {
	command := exec.Command("sudo", append([]string{cmd}, args...)...)
	command.Stdin = strings.NewReader(password + "\n")
	err := command.Run()
	if err != nil {
		return fmt.Errorf("error executing command : %v", err)
	}
	return nil
}

func changeMac(interfaceName, newMac, password string) error {
	err := runCommand("ifconfig", []string{interfaceName, "down"}, password)
	if err != nil {
		return err
	}

	err = runCommand("ifconfig", []string{interfaceName, "hw", "ether", newMac}, password)
	if err != nil {
		return err
	}
	err = runCommand("ifconfig", []string{interfaceName, "up"}, password)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	fmt.Print("\n enter network interface: ")

	interfacName, _ := reader.ReadString('\n')
	interfacName = strings.TrimSpace(interfacName)

	fmt.Println("enter new mac address: ")
	newMac, _ := reader.ReadString('\n')
	newMac = strings.TrimSpace(newMac)

	err := changeMac(interfacName, newMac, password)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("mac address changed successfully")
	}
}
