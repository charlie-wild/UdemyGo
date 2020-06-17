package main

import (
	"fmt"
	"testing"
)

func TestHTTP(t *testing.T) {
	_, err := fmt.Println(httpRead())
	if err != nil {
		t.Errorf("Response was incorrect, got %d, want a response.", err)
	}

}
