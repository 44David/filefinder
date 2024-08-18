package main

import "github.com/charmbracelet/huh"


func main() {
	var file string

	huh.NewInput().
		Title("Start searching").
		Prompt("? ").
		Value(&file).
		Run()
}