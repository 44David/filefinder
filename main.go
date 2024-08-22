package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"	
	"github.com/gdamore/tcell/v2"
)


func fileExist(file string) bool {
	_, error := os.Stat(file)

	return !errors.Is(error, os.ErrNotExist)
}


func main() {
	var input string
	
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	screen, screenErr := tcell.NewScreen()
	if screenErr != nil {
		log.Fatal("Screen error.")
	}
	if screenErr := screen.Init(); screenErr != nil {
		log.Fatal("Screen init error")
	}

	screen.SetStyle(defStyle)
	screen.Clear()


	for {
		screen.Show()
		ev := screen.PollEvent()

		switch ev := ev.(type) {
			case *tcell.EventResize:
				screen.Sync()

			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					return
				}
				fmt.Println(ev.Key())
		}
	}

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
