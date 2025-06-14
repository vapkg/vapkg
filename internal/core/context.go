package core

import "os"

type Context struct {
	config    IConfig
	core      Core
	pwd       string
	ws        Workspace
	commands  CommandRegistry
	providers ProviderRegistry
	logger    ILogger
}

func NewContext(logger ILogger, cfg IConfig) *Context {

	if pwd, err := os.Getwd(); err == nil {

		return &Context{
			config:    cfg,
			core:      CreateCore(),
			pwd:       pwd,
			ws:        CreateWorkspace(pwd),
			commands:  CreateCommandRegistry(),
			providers: CreateProviderRegistry(),
			logger:    logger,
		}
	}

	return nil
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

func (ctx *Context) Logger() ILogger {
	return ctx.logger
}

func (ctx *Context) Providers() *ProviderRegistry {
	return &ctx.providers
}

func (ctx *Context) Config() IConfig {
	return ctx.config
}

func (ctx *Context) Close() {
	if ctx.logger != nil {
		ctx.logger.Close()
	}
}
