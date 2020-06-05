package caye

import "github.com/kminehart/caye/languages/golang"

// V1Languages ...
type V1Languages struct {
}

func (l *V1Languages) Go() Language {
	return &golang.Language{}
}
