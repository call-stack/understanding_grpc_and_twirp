package main

import (
	"fmt"
	"net/http"

	"github.com/kalpitpant/understanding-twirp/protoc/fdm"
	"github.com/kalpitpant/understanding-twirp/server/fdmserver"
)

func main() {
	server := &fdmserver.FDMServer{}
	twirpHandler := fdm.NewFileDownloadManagerServer(server)

	fmt.Println("Starting the server....")
	http.ListenAndServe(":8080", twirpHandler)
}
