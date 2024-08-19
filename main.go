package main

import (
	"errors"
	//"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
)


func fileExist(file string) bool {
	_, error := os.Stat(file)

	return !errors.Is(error, os.ErrNotExist)
}

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
			if !fileExist(filename) { 
				return errors.New("this file does not exist")
			} 

				return nil
		}).
		Value(&file).
		Run()

	
}
