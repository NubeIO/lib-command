package unixcmd

import (
	"fmt"
	"github.com/NubeIO/lib-command/command"
	"testing"
)

func TestUnixCommand_Uptime(t *testing.T) {

	out, err := New(&command.Command{}).Uptime(command.DryRun)
	fmt.Println(out.CommandsJoined, err)

	out, err = New(&command.Command{}).Uptime(command.Run)
	fmt.Println(out, err)

	out, err = New(&command.Command{}).NodeGetVersion(command.Run)
	fmt.Println(out, err)
	if err != nil {
		return
	}
	fmt.Println(out.Out, err)

	pro, err := New(&command.Command{}).DetectArch()
	fmt.Println(pro, err)

}
