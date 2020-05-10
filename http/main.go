package main

import (
	"fmt"
	"net/http"
	"os"
)


func main()  {
	 resp, err := http.Get("http://google.co.uk")
	 if err != nil {
		 fmt.Println("Error:", err)
		 os.Exit(1)
	 }

	 // create a byte slice that holds the data written into by the read function
}