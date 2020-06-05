package caye

import (
	context "context"

	"github.com/kminehart/caye/config"
	grpc "google.golang.org/grpc"
)

// Client encompasses all of the functionality needed to interact with the Caye server from a CI environment
type Client struct {
	Context context.Context
	Builds  BuildsClient
}

// New creates a new Caye client, and panics if that shit ain't happening, panic
func New() *Client {
	cfg := config.GetConfig()
	conn, err := grpc.Dial(cfg.Endpoint, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	c := &Client{
		Builds:  NewBuildsClient(conn),
		Context: context.Background(),
	}

	// res, err := c.Builds.StartBuild(c.Context, &StartBuildRequest{})
	// if err != nil {
	// 	log.Fatalln("Failed to initialize build", err)
	// }

	// c.Context = context.WithValue(c.Context, ContextKeyBuildID, res.GetId())
	return c
}
