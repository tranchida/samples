package main

import (
	"github.com/fatih/color"
)

func main() {
	
	red := color.New(color.FgRed)
	red.Println("Hi there!")

	blue := color.New(color.FgBlue)
	blue.Println("Hi there!")

}