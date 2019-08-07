package main

import (
	"fmt"
)

var (
	i, j string
)

func register() {
	fmt.Println("im in register")
	options()
}

func options() {
	fmt.Println("1. Create new account")
	fmt.Println("2. view account")
	fmt.Println("3. modify account")
	fmt.Println("4. amount")
	fmt.Println("5. exit")
	fmt.Println("Please select option")
	fmt.Scanf("%s", &i)
}
func View() {
	fmt.Println("im in view")
	options()
}
func Edit() {
	fmt.Println("im in edit")
	options()
}
func Transfer() {
	fmt.Println("trnsfer")
	options()
}
func main() {
	fmt.Println("welcome to bank application")
	options()
	//fmt.Scanf("%s", &i)
	for {
		switch i {
		case "1":
			register()
			//options()
			//fallthrough
		case "2":
			View()
			//options()
			//fallthrough
		case "3":
			Edit()
			//options()
			//fallthrough
		case "4":
			Transfer()
			//options()
			//fallthrough
		case "5":
			break
		default:
			fmt.Printf("Please enter one of the options")
			options()
		}
	}
}
