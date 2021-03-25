package fdmserver

import (
	"context"

	"github.com/kalpitpant/understanding-twirp/protoc/fdm"
)

//FDMServer is the struct which will handle server side things.
type FDMServer struct{}

//Health is a menthods belongs to FDMServer
func (f *FDMServer) Health(c context.Context, m *fdm.StatusMessage) (*fdm.StatusMessage, error) {

	return &fdm.StatusMessage{Body: "Server Health is ok."}, nil
}
