package main

import (
	"errors"
	"fmt"
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
		//TODO check if string is a folder or file
		Validate(func(filename string) error {
			if !fileExist(filename) { 
				return errors.New("this file does not exist")
			} 

				return nil
		}).
		Value(&file).
		Run()

	
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
	}

	if fileInfo.IsDir() {
		cmdDir := exec.Command("code", file)
		cmdDirError := cmdDir.Run()

		if cmdDirError != nil {
			fmt.Println(cmdDirError.Error())
		}	

	} else {
		nano, nanoError := exec.LookPath("nano")
		if err != nil {
			log.Fatal(nanoError)
		} 
		
		//TODO fix this
		cmd := exec.Command("sudo", nano, file)
		cmdError := cmd.Run()

		if cmdError != nil {
			fmt.Println(cmdError.Error())
		}

	}



	
}
