package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	commands := GetCommands()
	for{
		fmt.Print("pokedex > ")
		for scanner.Scan(){
			text := scanner.Text()
			if command, ok := commands[text]; ok{
				command.callback()
			} else{
				fmt.Print("Command not found!")
			}
			fmt.Print("\n")
			break
		}
	}
}