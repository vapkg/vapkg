package core

type ICommandRegistry interface {
	Register(sig string, command ICommand)
	Get(sig string) ICommand
	Unregister(sig string)
	Exists(sig string) bool
}

var _ ICommandRegistry = (*CommandRegistry)(nil)

type CommandRegistry struct {
	commands map[string]ICommand
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{make(map[string]ICommand)}
}

func CreateCommandRegistry() CommandRegistry {
	return CommandRegistry{make(map[string]ICommand)}
}

func (registry *CommandRegistry) Register(sig string, command ICommand) {
	registry.commands[sig] = command
}

func (registry *CommandRegistry) Get(sig string) ICommand {
	return registry.commands[sig]
}

func (registry *CommandRegistry) Unregister(sig string) {
	delete(registry.commands, sig)
}

func (registry *CommandRegistry) Exists(sig string) bool {
	_, ok := registry.commands[sig]
	return ok
}
