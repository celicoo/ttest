package internal

type NoopCommand struct {}

func (NoopCommand) Help() {}
func (NoopCommand) Error(error) {}
func (NoopCommand) Run() {}
