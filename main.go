package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id            string
	Password      string
	Name          string
	Age           int
	AccountNumber string
}

func SuccessfullLogin() {
	var option string

	fmt.Println("1.Withdraw money.\n2.Deposit money.\n3.Request balance.\n4.Quit the program.")
	fmt.Scan(&option)

	switch option {
	case "1":
		fmt.Println("Please withdraw money")
	case "2":
		fmt.Println("Please deposit money")
	case "3":
		fmt.Println("Please check balance")
	case "4":
		fmt.Println("exiting...")
		os.Exit(3)
	}

}

func login() {

	// Declaring some variables
	var id string
	var password string

	fmt.Println("Enter ID")
	fmt.Scan(&id)
	fmt.Println("Enter password")
	fmt.Scan(&password)

	// login(id, password)

	f, err := os.Open("userfile.json")

	defer f.Close()

	var obj1 User

	err = json.NewDecoder(f).Decode(&obj1)

	if err != nil {
		panic(err)
	}

	if id == obj1.Id && password == obj1.Password {
		fmt.Println("Login Successfull")

		SuccessfullLogin()

	} else {
		fmt.Println("ID or password is incorrect")
		main()
	}

}

func main() {
	var selectedOption string

	fmt.Println("Hi! Welcome to Mr. Mukesh ATM Machine! \n\nPlease select an option from the menu below: ")
	fmt.Println("\nl -> Login  \nq -> Quit")
	fmt.Scan(&selectedOption)
	fmt.Println("You have selected: ", selectedOption)

	switch selectedOption {
	case "l":
		login()
	case "q":
		os.Exit(3)
	default:
		fmt.Println("Please enter valid option")
		main()
	}

}
