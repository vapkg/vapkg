package core

import (
	"vapkg/internal/utils"
)

type Context struct {
	core     Core
	pwd      string
	ws       Workspace
	commands CommandRegistry
	logger   utils.ILogger
}

func NewContext(pwd string, cfg *Config) *Context {
	var err error
	var logger utils.ILogger

	if logger, err = utils.CreateActualLogger(cfg.LogFolder, cfg.LogLevel); err != nil {
		return nil
	}

	return &Context{
		core:     CreateCore(),
		pwd:      pwd,
		ws:       CreateWorkspace(pwd),
		commands: CreateCommandRegistry(),
		logger:   logger,
	}
}

func (ctx *Context) Commands() *CommandRegistry {
	return &ctx.commands
}

func (ctx *Context) Ws() *Workspace {
	return &ctx.ws
}

func (ctx *Context) Pwd() string {
	return ctx.pwd
}

func (ctx *Context) Core() ICore {
	return &ctx.core
}

func (ctx *Context) Logger() utils.ILogger {
	return ctx.logger
}

func (ctx *Context) Close() {
	if ctx.logger != nil {
		ctx.logger.Close()
	}
}
