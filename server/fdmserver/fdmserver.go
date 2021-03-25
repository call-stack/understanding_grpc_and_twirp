package fdmserver

import (
	"context"
	"fmt"

	"github.com/kalpitpant/understanding-twirp/protoc/fdm"
)

//FDMServer is the struct which will handle server side things.
type FDMServer struct{}

//Health is a menthods belongs to FDMServer
func (f *FDMServer) Health(c context.Context, m *fdm.StatusMessage) (*fdm.StatusMessage, error) {

	return &fdm.StatusMessage{Body: "Server health is ok."}, nil
}

//Download is a menthods which will get the image urls to download images.
func (f *FDMServer) Download(c context.Context, m *fdm.URLs) (*fdm.StatusMessage, error) {
	for _, url := range m.Url {
		fmt.Println("Get the image from url--> ", url)
	}
	return &fdm.StatusMessage{Body: "We got the urls"}, nil
}
