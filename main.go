package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	open := flag.Bool("o", false, "Open the vault")
	close := flag.Bool("c", false, "Close the vault")
	name := flag.String("n", "", "Name of the vault")
	pass := flag.String("p", "", "Password for the vault")

	flag.Parse()

	path := os.Getenv("VAULT")
	fmt.Println(path)
	fmt.Println(*open)
	if *open {
		command1 := `unzip`
		command2 := path + `/.saves/` + *name + `.zip`
		command3 := *pass
		fmt.Println(command1, "-P", command3, command2)
		cmd := exec.Command(command1, "-P", command3, command2)
		fmt.Println(command1, command2, command3)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else if *close {
		cmd := exec.Command("zip", "-re", "-P", *pass, *name+".zip", *name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		cmd = exec.Command("mv", *name+".zip", path+"/.saves/")
		cmd.Run()
	}

}
