package wrappers

import (
	"context"

	"github.com/kminehart/caye"
)

type StatusWrapper struct {
	Caye   caye.Caye
	Builds caye.BuildsClient
}

func (c *StatusWrapper) Run(ctx context.Context, args ...string) (*caye.Result, error) {
	id := ctx.Value(caye.ContextKeyBuildID)
	_, err := c.Builds.StartRun(ctx, &caye.StartRunRequest{
		BuildId: id.(string),
	})
	if err != nil {
		c.Builds.EndRun(ctx, &caye.EndRunRequest{
			BuildId: id.(string),
			Error:   true,
		})
		return nil, err
	}

	result, err := c.Caye.Run(ctx, args...)

	c.Builds.EndRun(ctx, &caye.EndRunRequest{
		BuildId: id.(string),
		Error:   err != nil,
	})

	return result, err
}
