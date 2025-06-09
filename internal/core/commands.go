package core

type CommandRegistry struct {
	commands map[string]*Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{make(map[string]*Command)}
}

func CreateCommandRegistry() CommandRegistry {
	return CommandRegistry{make(map[string]*Command)}
}

func (registry *CommandRegistry) Register(sig string, command *Command) {
	registry.commands[sig] = command
}

func (registry *CommandRegistry) Get(sig string) *Command {
	return registry.commands[sig]
}

func (registry *CommandRegistry) Unregister(sig string) {
	delete(registry.commands, sig)
}

func (registry *CommandRegistry) Exists(sig string) bool {
	_, ok := registry.commands[sig]
	return ok
}
