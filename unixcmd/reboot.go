package unixcmd

import (
	"github.com/NubeIO/lib-command/command"
	log "github.com/sirupsen/logrus"
	"time"
)

func (inst *UnixCommand) HostReboot(options ...command.Options) (*command.Response, error) {
	log.Errorln("HOST WILL REBOOT IN 5 Sec good luck :)")
	time.Sleep(5 * time.Second)
	return cmd.Builder("sudo shutdown", "-r now").RunCommand(options...)
}
