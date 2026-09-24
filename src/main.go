package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var user *User
	reader := bufio.NewReader(os.Stdin)
	for {
		for {
			fmt.Print("Enter your role: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "admin" || input == "employee" || input == "guest" {
				user = CurrUser(input)
				break
			} else {
				fmt.Print("Not a role! Try again\n\n")
			}
		}
		for {
			fmt.Print("\nEnter your action\nor\nQ to Quit\nor\nR to change role\n:")
			action, _ := reader.ReadString('\n')
			action = strings.TrimSpace(action)

			if action == "Q" || action == "q" {
				fmt.Println("Goodbye!\n")
				return
			}

			if action == "R" || action == "r" {
				fmt.Println("Switching Role\n")
				break
			}

			switch action {
			case "delete":
				if user.Type == "admin" {
					fmt.Print("You can delete!\n\n")
				} else {
					fmt.Print("Access Denied\n\n")
				}
			case "write":
				if user.Type == "admin" || user.Type == "employee" {
					fmt.Print("You can write!\n\n")
				} else {
					fmt.Print("Access Denied\n\n")
				}
			case "read":
				fmt.Print("You can read\n\n")
			default:
				fmt.Print("Not a valid action, try again\n\n")
			}
		}

	}
}
