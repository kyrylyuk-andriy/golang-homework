package hw12

import (
	"strings"
)

type TextDecorator interface {
	Decorate(text string) string
}

type RemoveSpaceDecorator struct{}

func (r *RemoveSpaceDecorator) Decorate(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

type ReplaceCharachterDecorator struct {
	OldChar string
	NewChar string
}

func (r *ReplaceCharachterDecorator) Decorate(text string) string {
	return strings.ReplaceAll(text, r.OldChar, r.NewChar)
}
