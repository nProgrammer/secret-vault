package controllers

import (
	"os"
	"os/exec"
)

func Open(path string, name *string, pass *string) {
	command1 := `unzip`
	command2 := path + `/.saves/` + *name + `.zip`
	command3 := *pass
	cmd := exec.Command(command1, "-P", command3, command2)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func Close(path string, name *string, pass *string) {
	cmd := exec.Command("zip", "-re", "-P", *pass, *name+".zip", *name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = exec.Command("mv", *name+".zip", path+"/.saves/")
	cmd.Run()
}
