package core

import "vapkg/internal/utils"

type Context struct {
	core     Core
	pwd      string
	ws       Workspace
	commands CommandRegistry
	logger   utils.Logger
}

func CreateContext(pwd string) Context {
	return *NewContext(pwd)
}

func NewContext(pwd string) *Context {
	return &Context{
		core:     CreateCore(),
		pwd:      pwd,
		ws:       CreateWorkspace(pwd),
		commands: CreateCommandRegistry(),
		logger:   utils.CreateLogger(utils.InfoLog | utils.DebugLog | utils.WarnLog),
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
	return &ctx.logger
}
