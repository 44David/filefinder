package main

import (
	"errors"

	"github.com/charmbracelet/huh"
)


func main() {
	var file string

	huh.NewInput().
		Title("Start searching").
		Prompt("? ").
		Validate(func(filename string) error {
			if filename == "Test" { 
				return errors.New("this file does not exist")
			} 

			return nil
		}).
		Value(&file).
		Run()
}
