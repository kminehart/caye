package caye

import (
	"bytes"
	context "context"
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

type defaultCaye struct {
}

func Default() Caye {
	return &defaultCaye{}
}

func (d *defaultCaye) Languages() Languages {
	return &V1Languages{}
}

func (d *defaultCaye) Run(ctx context.Context, args ...string) (*Result, error) {
	return Run(ctx, args...)
}

func (d *defaultCaye) Done() {
}

// Run runs a command
func Run(ctx context.Context, args ...string) (*Result, error) {
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	output := bytes.NewBuffer(nil)
	cmd.Stdout = io.MultiWriter(os.Stdout, output)
	cmd.Stderr = io.MultiWriter(os.Stderr, output)
	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, output.String())
	}

	return &Result{
		Output: output.String(),
	}, nil
}
