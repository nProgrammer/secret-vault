package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Name of vault: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Password for vault: ")
	pass, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	pass = strings.TrimSpace(pass)
	path := os.Getenv("VAULT")
	command1 := `unzip`
	command2 := path + `/.saves/` + name + `.zip`
	command3 := pass
	cmd := exec.Command(command1, "-P", command3, command2)
	fmt.Println(command1, command2, command3)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
