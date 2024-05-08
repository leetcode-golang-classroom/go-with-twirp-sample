package main

import (
	"net/http"

	"github.com/leetcode-golang-classroom/go-with-twirp-sample/internal/haberdasherserver"
	"github.com/leetcode-golang-classroom/go-with-twirp-sample/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
