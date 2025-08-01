package core

import (
	"fmt"
)

const OptionPrefix = "--"

// ctx + opts with value
type CommandHandlerFn func(c IContext, m map[string]string) error

type ICommand interface {
	Execute(ctx IContext, opts map[string]string) error
}

var _ ICommand = (*Command)(nil)

type Command struct {
	Usage       string
	Description string
	Handler     CommandHandlerFn
	Options     map[string]bool
}

func (c *Command) Execute(ctx IContext, opts map[string]string) error {

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
			return fmt.Errorf("option '%s%s' is required", OptionPrefix, k)
		}
	}

	return c.Handler(ctx, opts)
}
