package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1]) //os.Open takes an argument, returns a variable of type file and an error

	if err != nil {
		fmt.Println("Error is: ", err)
		os.Exit(1) //exit the program if there is an error
	}

	io.Copy(os.Stdout, file) //can use io.Copy as the type File satisfies the reader interface

}
