package app

var registry []Command

// Register adds a subcommand (call from main init() or a command package init).
func Register(c Command) {
	registry = append(registry, c)
}

// All returns registered commands in Register order.
func All() []Command {
	return registry
}

// FindCommand matches Name() or Aliases() when implemented.
func FindCommand(name string) Command {
	for _, c := range registry {
		if c.Name() == name {
			return c
		}
		if a, ok := c.(interface{ Aliases() []string }); ok {
			for _, al := range a.Aliases() {
				if al == name {
					return c
				}
			}
		}
	}
	return nil
}
