package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"github.com/charmbracelet/huh"
	"github.com/peterh/liner"
)

var (
	Cyan = "\033[36m"
	Green = "\033[32m"
	Red = "\033[31m"
	defaultColor = "\033[0m"
)

var names []string
var editor string

func config() {

	huh.NewSelect[string]().
		Title("Choose your preferred text editor").
		Options(
			huh.NewOption("Vim", "vim"),
			huh.NewOption("Nano", "nano"),
		).
		Value(&editor).
		Run()

	configFile, err := os.Create("config.txt") 
	if err != nil {
		fmt.Println(err)
		return
	} 

	configFile.WriteString(editor)
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

		names = strings.Split((string(out)), "\n")

		line := liner.NewLiner()
		defer line.Close()

		line.SetCtrlCAborts(true)

		line.SetCompleter(func(line string) (c []string) {
			for _, n := range names {
				if strings.HasPrefix(strings.ToLower(n), line) {
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
			fmt.Println(Red + "*file was not found*" + defaultColor)
		}

		if fileInfo.IsDir() {
			cmdDir := exec.Command("code", input)
			fmt.Println("Opening directory...")
			cmdDirError := cmdDir.Run()
			fmt.Println("Opened.")

			if cmdDirError != nil {
				fmt.Println(Red + "*could not open directory*" + defaultColor)
			}	

		} else {
			textEditor, err := os.ReadFile("config.txt")
			if err != nil {
				fmt.Println(Red + "*an error occurred with your preferred text editor*" + defaultColor)
			}

		
			cmd := exec.Command("sudo", string(textEditor), input)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			cmdError := cmd.Run()

			if cmdError != nil {
				if string(textEditor) == "" {
					fmt.Println(Red + "*could not find a preferred text editor")
				} else {
					fmt.Println(Red + "*could not open file in " + string(textEditor) + "*" + defaultColor)
				}
				
			}

		}
	}
}
