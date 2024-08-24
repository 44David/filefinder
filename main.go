package main

import (
	"errors"
	"fmt"
	"log"

	//"fmt"
	//"log"
	"os"
	"os/exec"
	"strings"

	"github.com/peterh/liner"
)


func fileExist(file string) bool {
	_, error := os.Stat(file)

	return !errors.Is(error, os.ErrNotExist)
}

	var (
		names []string
	)


func main() {

	var input string

	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}	

	names = strings.Split(strings.ToLower(string(out)), "\n")

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range names {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	fmt.Println(string(out))

	if file, err := line.Prompt("? "); err == nil {
		log.Print("This is the name: ", file)
		input = file
	} else if err == liner.ErrPromptAborted {
		log.Print("Process stopped")
	} else {
		log.Print("Error reading line.")
	}


	
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
