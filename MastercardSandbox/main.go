package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"github.com/gomodule/oauth1/oauth" //package for OAuth1.0a header complier
)

func main() {
	//lambda.Start(HandleRequest)
	fmt.Println(httpRead())

}

func HandleRequest(ctx context.Context) (string, error) {
	resp := httpRead()
	return resp, nil
}

func httpRead() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://google.com", nil)
	req.Header.Add("Header", "Value")
	resp, err := client.Do(req)

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
