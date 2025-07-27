package core

import "vapkg/internal/workspace"

type IContext interface {
	Commands() ICommandRegistry
	Workspace() workspace.IWorkspaceManager
	Core() ICore
	Logger() ILogger
	Providers() *ProviderRegistry
	Config() IConfig
	Close()
}

var _ IContext = (*Context)(nil)

type Context struct {
	config    IConfig
	core      Core
	commands  CommandRegistry
	providers ProviderRegistry
	workspace workspace.IWorkspaceManager
	logger    ILogger
}

func NewContext(logger ILogger, cfg IConfig) (*Context, error) {
	wsm := workspace.NewManager()
	if wsm.IsExist() {
		if err := wsm.LoadWorkspace(); err != nil {
			return nil, err
		}
	}

	return &Context{
		config:    cfg,
		core:      CreateCore(),
		commands:  CreateCommandRegistry(),
		workspace: wsm,
		providers: CreateProviderRegistry(),
		logger:    logger,
	}, nil
}

func (ctx *Context) Commands() ICommandRegistry {
	return &ctx.commands
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

func (ctx *Context) Workspace() workspace.IWorkspaceManager {
	return ctx.workspace
}

func (ctx *Context) Close() {
	if ctx.logger != nil {
		ctx.logger.Close()
	}
}
