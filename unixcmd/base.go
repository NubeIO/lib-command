package unixcmd

import "github.com/NubeIO/lib-command/command"

var cmd *command.Command

type UnixCommand struct {
	CMD *command.Command
}

func New(c *command.Command) *UnixCommand {
	cmd = c
	return &UnixCommand{}
}

func (inst *UnixCommand) Uptime(options ...command.Options) (*command.Response, error) {
	return cmd.Builder("uptime").RunCommand(options...)
}
