package core

import (
	"fmt"
)

// ctx + opts with value
type CommandHandlerFn func(*Context, map[string]string) error

type Command struct {
	Usage       string
	Description string
	Handler     CommandHandlerFn
	Options     map[string]bool
}

func (c *Command) Execute(ctx *Context, opts map[string]string) error {

	if ctx == nil {
		return fmt.Errorf("context is nil")
	}

	for k := range opts {
		if _, ok := c.Options[k]; !ok {
			return fmt.Errorf("invalid option %s%s", OptionPrefix, k)
		}
	}

	for k := range c.Options {
		if !c.Options[k] {
			continue
		}

		if _, ok := opts[k]; !ok {
			return fmt.Errorf("missing required option %s", k)
		}
	}

	return c.Handler(ctx, opts)
}
