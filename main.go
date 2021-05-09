package main

import (
	"os"
	"secret-vault/controllers"
)

func main() {
	open, close, name, pass := controllers.Flags()

	path := os.Getenv("VAULT")
	if *open {
		controllers.Open(path, name, pass)
	} else if *close {
		controllers.Close(path, name, pass)
	}

}
