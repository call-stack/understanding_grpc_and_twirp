package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kalpitpant/understanding-twirp/protoc/fdm"
)

func main() {
	client := fdm.NewFileDownloadManagerProtobufClient("http://localhost:8080", &http.Client{})

	// msg, er := client.Health(context.Background(), &fdm.StatusMessage{})

	// if er != nil {
	// 	fmt.Println("There is some errro in the code.")
	// }

	imageURLs := []string{"url1", "url2"}

	msg, er := client.Download(context.Background(), &fdm.URLs{Url: imageURLs})

	if er != nil {
		fmt.Println("There is some errro in the code.")
	}

	fmt.Println(msg.Body)

}
