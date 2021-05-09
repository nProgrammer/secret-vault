package controllers

import "flag"

func Flags() (*bool, *bool, *string, *string) {
	open := flag.Bool("o", false, "Open the vault")
	close := flag.Bool("c", false, "Close the vault")
	name := flag.String("n", "", "Name of the vault")
	pass := flag.String("p", "", "Password for the vault")

	flag.Parse()

	return open, close, name, pass
}
