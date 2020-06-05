package git

import (
	"context"
	"strings"

	"github.com/kminehart/caye"
)

type CLI struct{}

func (c *CLI) GetCommitSHA(ctx context.Context) (string, error) {
	result, err := caye.Run(ctx, "git", "rev-parse", "HEAD")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(result.GetOutput()), nil
}

func NewCLI() *CLI {
	return &CLI{}
}
