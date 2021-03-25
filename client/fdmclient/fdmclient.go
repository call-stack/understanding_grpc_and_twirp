package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kalpitpant/understanding-twirp/protoc/fdm"
)

func main() {
	client := fdm.NewFileDownloadManagerProtobufClient("http://localhost:8080", &http.Client{})

	msg, er := client.Health(context.Background(), &fdm.StatusMessage{Body: "Show me the message"})

	if er != nil {
		fmt.Println("There is some errro in the code.")
	}

	fmt.Println(msg.Body)

}
