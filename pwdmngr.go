package main

import (
	"fmt"
	"os"
)

const pwd_db = "password.db"

func store(platform string, username string, password string) {
	entry := platform + "," + username + "," + password + "\n"
	f, err := os.OpenFile(pwd_db, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(entry)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	fmt.Println("Password manager")

	var args []string
	args = os.Args
	// os.Args is an array storing all the command line parameters in order they were passed to program.
	// fmt.Println(args)
	if len(args) < 5 { // check if there are enough arguments
		fmt.Println("Usage: go run main.go add platform username password")
		return
	}

	if args[1] == "add" {
		store(args[2], args[3], args[4])
	}
}
