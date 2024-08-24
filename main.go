package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"	
)


func fileExist(file string) bool {
	_, error := os.Stat(file)

	return !errors.Is(error, os.ErrNotExist)
}


func main() {
	var input string


	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}	
	
	fmt.Println(string(out))

	fmt.Scan(&input)
	
	
	fileInfo, err := os.Stat(input)
	if err != nil {
		fmt.Println(err)
	}

	if fileInfo.IsDir() {
		cmdDir := exec.Command("code", input)
		fmt.Println("Opening directory...")
		cmdDirError := cmdDir.Run()
		fmt.Println("Opened.")

		if cmdDirError != nil {
			fmt.Println(cmdDirError.Error())
		}	

	} else {
		nano, nanoError := exec.LookPath("nano")
		if err != nil {
			log.Fatal(nanoError)
		} 
		
		cmd := exec.Command("sudo", nano, input)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmdError := cmd.Run()

		if cmdError != nil {
			fmt.Println(cmdError.Error())
		}

	}
}
