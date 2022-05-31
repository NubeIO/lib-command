package command

import (
	"fmt"
	"strings"
	"testing"
)

func TestCMD(t *testing.T) {
	cmd := New(&Command{SetPath: "/home/aidan", Commands: []string{"ls"}})
	str := []string{"/home/aidan", "ls"}
	fmt.Println(strings.Join(str[:], " "))
	out, err := cmd.RunCommand()
	fmt.Println(out.Out, err)
	out, err = cmd.Builder("uptime").RunCommand()
	fmt.Println(out.Out, err)

	out, err = cmd.Builder("uptime").RunCommand()
	fmt.Println(out.Out, err)

}
