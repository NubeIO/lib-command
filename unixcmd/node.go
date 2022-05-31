package unixcmd

import (
	"github.com/NubeIO/lib-command/command"
	"strings"
)

func (inst *UnixCommand) NodeGetVersion(options ...command.Options) (*command.Response, error) {
	res, err := cmd.Builder("nodejs", "-v").RunCommand(options...)
	cmdOut := res.Out
	if strings.Contains(cmdOut, "v") {
		res.Out = cmdOut
		res.Ok = true
		return res, err
	} else {
		res.Out = "not installed"
		res.Ok = false
		return res, err
	}
}
