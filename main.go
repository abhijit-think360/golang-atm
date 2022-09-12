package main

import (
	"encoding/json"
	"os"

	// "log"
	"fmt"
)

type user struct {
	Username  string
	Pin       int
	Balance   int
	Statement []string
}

func EncodeUser(username string, pin int, balance int) string {
	newUser := user{username, pin, balance, nil}
	newUserJson, err := json.MarshalIndent(newUser, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(newUserJson))
	return string(newUserJson)
}

func register() {
	var username string
	var pin int
	var balance int
	fmt.Println("Enter a username:")
	fmt.Scanln(&username)
	fmt.Println("Enter a pin:")
	fmt.Scanln(&pin)
	fmt.Println("Enter the initial balance:")
	fmt.Scanln(&balance)
	// var newUser = user{username, pin, balance, nil}
	// fmt.Println(newUser)
	// fmt.Println("********")
	// EncodeUser(username, pin, balance)
	// fmt.Println("********")
	file, _ := json.Marshal(EncodeUser(username, pin, balance))
	// _ = ioutil.WriteFile("users.json", file, fs.ModeAppend)
	// // _ = ioutil.WriteFile()
	f, err := os.OpenFile("users.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(string(file)); err != nil {
		panic(err)
	}
}

func main() {
	var userType int
	fmt.Println("Choose one of the below options:")
	fmt.Println("1. Registered User")
	fmt.Println("2. New User")
	fmt.Scanln(&userType)
	if userType == 1 {
		register()
	} else {
		fmt.Println("Invalid User")
	}

}
