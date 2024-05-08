package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/leetcode-golang-classroom/go-with-twirp-sample/rpc/haberdasher"
)

func main() {
	client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

	hat, err := client.MakeHat(context.Background(), &haberdasher.Size{Inches: 16})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Printf("I have a nice new hat: %+v\n", hat)
}
