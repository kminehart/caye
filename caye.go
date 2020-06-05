package caye

import context "context"

// Language defines standard operations that are language dependent
type Language interface {
	Build() error
	Test() error
}

type Languages interface {
	Go() Language
}

type Logger interface {
	Errorln()
	Errorf()

	Println()
	Printf()

	Fatalln()
	Fatalf()
}

// The Caye interface defines the general type used to interact with the Caye service.
type Caye interface {
	Logger
	Languages() Languages
	Run(ctx context.Context, args ...string) (*Result, error)
	Done()
}

// Result ...
type Result struct {
	Output string
}

// GetOutput ...
func (r *Result) GetOutput() string {
	return r.Output
}
