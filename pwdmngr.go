package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func retrieve(platform string) {
	file, err := os.Open(pwd_db)
	if err != nil {
		fmt.Println(err)
		return
	}

	input := bufio.NewScanner(file)
	// . bufio is an inbuilt package which does buffered IO operations. It’s very convenient to use. It’s scanner type which we are getting by bufio.NewScanner(f) can read input and break into lines which we are using in our code.

	for input.Scan() {
		entry := strings.Split(input.Text(), ",")
		if entry[0] == platform {
			fmt.Println(entry[1], entry[2])
			return
		}
	}
	fmt.Printf("Platform %s not known\n", platform)

}

func main() {
	fmt.Println("Password manager")

	var args []string
	args = os.Args
	// os.Args is an array storing all the command line parameters in order they were passed to program.
	// fmt.Println(args)
	if args[1] == "add" {
		store(args[2], args[3], args[4])
	} else if args[1] == "get" {
		retrieve(args[2])
	} else {
		fmt.Println("Invalid operation ", args[1])
	}
}
