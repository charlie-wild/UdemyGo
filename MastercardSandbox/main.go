package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"github.com/gomodule/oauth1/oauth" //package for OAuth1.0a header complier
)

func main() {

	fmt.Println(httpRead())

}

func httpRead() string {
	resp, err := http.Get("https://google.com")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return (string(body))

}
