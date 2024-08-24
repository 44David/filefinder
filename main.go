package main

import (
	//"errors"
	"fmt"
	"log"
	//"log"
	"os"
	"os/exec"
	"strings"
	"github.com/peterh/liner"
)

var Cyan = "\033[36m"
var Green = "\033[32m"
var defaultColor = "\033[0m"

var names []string


func config() {
	line := liner.NewLiner()
	defer line.Close()

	
	line.SetCtrlCAborts(true)

	var editors = []string{"vim", "nano"}

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range editors {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	fmt.Printf(Green + "vim \nnano\n" + defaultColor)
	if input, err := line.Prompt("Type option: "); err == nil {

		configFile, err := os.Create("config.txt") 
		if err != nil {
			fmt.Println(err)
			return
		} 

		configFile.WriteString(input)

	}
}

func main() {

	if (len(os.Args) == 2) && (os.Args[1] == "config") {
		config()

	} else {

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


		fmt.Println(Cyan + string(out) + defaultColor)

		if file, err := line.Prompt("? "); err == nil {
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
			
			cmd := exec.Command("sudo", "nano", input)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			cmdError := cmd.Run()

			if cmdError != nil {
				fmt.Println(cmdError.Error())
			}

		}
	}
}
