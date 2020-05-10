package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)


func main()  {
	 resp, err := http.Get("http://google.co.uk")
	 if err != nil {
		 fmt.Println("Error:", err)
		 os.Exit(1)
	 }

	/* // create a byte slice that holds the data written into by the read function
	 bs := make([]byte, 99999) //give me an empty byte slice with space for 99999 elements

	 resp.Body.Read(bs)
	 fmt.Println(string(bs)) */

	 io.Copy(os.Stdout, resp.Body)
}