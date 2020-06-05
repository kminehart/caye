package tools

import "context"

type Git interface {
	GetCommitSHA(context.Context) (string, error)
}
