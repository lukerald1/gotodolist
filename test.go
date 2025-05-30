package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var command string
	slice := make([]string, 0, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the TO DO List CLI app!")
	for {
		fmt.Println("Enter your command (create, read, update, delete or exit):")
		fmt.Scanln(&command)
		switch command {
		case "create":
			fmt.Println("Enter task name:")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input != "" {
				slice = append(slice, input)
			} else {
				fmt.Println("Please input task name")
			}
		case "read":
			fmt.Println("Your tasks:")
			for i, task := range slice {
				task = strings.TrimSpace(task)
				if task != "" {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "update":
			fmt.Println("Enter task name to update:")
			uptask, _ := reader.ReadString('\n')
			uptask = strings.TrimSpace(uptask)
			for i, _ := range slice {
				if slice[i] == uptask {
					fmt.Println("Enter new task name:")
					uptask1, _ := reader.ReadString('\n')
					slice[i] = uptask1
				}
			}
		case "delete":
			fmt.Println("Enter task name to remove:")
			detask, _ := reader.ReadString('\n')
			detask = strings.TrimSpace(detask)
			for i, task := range slice {
				if task == detask {
					slice = append(slice[:i], slice[i+1:]...)
					fmt.Println("Removed task #", i, "with name", task, "successfully!")
				}
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command! Please, try again!")
		}
	}
}
