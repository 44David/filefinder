package main

import (
	"errors"
	"github.com/charmbracelet/huh"
	"os/exec"
	"log"
)


func main() {
	var file string


	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}

	huh.NewInput().
		Title(string(out)).
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
