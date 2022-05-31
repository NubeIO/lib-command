package unixcmd

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-command/command"

	"strings"
)

type EdgeNetworking struct {
	IPAddress  string `json:"ip_address" post:"true"`
	SubnetMask string `json:"subnet_mask" post:"true"`
	Gateway    string `json:"gateway" post:"true"`
	SetDHCP    bool   `json:"set_dhcp" post:"true"`
	Password   string `json:"password"`
}

func (inst *UnixCommand) EdgeSetIP(net *EdgeNetworking, options ...command.Options) (ok bool, err error) {
	if net == nil {
		return false, errors.New("no values where valid")
	}
	arch, err := inst.DetectArch()
	if arch.IsBeagleBone {
		return false, errors.New("error incorrect arch type")
	}
	iface, err := inst.edge28Iface()
	if err != nil || iface == "" {
		return false, errors.New("error on get network interface name")
	}
	cmdOptions := ""
	if !net.SetDHCP {
		//_, err = validation.IsIPAddr(net.IPAddress)  //TODO add back in see https://github.com/NubeIO/lib-networking/blob/f38ed5db7a8ce8b8feae3a94ab8f19c6c6abe81b/ip/checks.go#L10
		//if err != nil {
		//	return false, errors.New(fmt.Sprintf(" %s couldn't be parsed as an IPAddress", net.IPAddress))
		//}
		//_, err = validation.IsIPAddr(net.SubnetMask)
		//if err != nil {
		//	return false, errors.New(fmt.Sprintf(" %s couldn't be parsed as an SubnetMask", net.SubnetMask))
		//}
		//_, err = validation.IsIPAddr(net.Gateway)
		//if err != nil {
		//	return false, errors.New(fmt.Sprintf(" %s couldn't be parsed as an Gateway", net.Gateway))
		//}
		cmdOptions = fmt.Sprintf("echo N00B2828 | sudo connmanctl config %s --ipv4 manual %s %s %s", iface, net.IPAddress, net.SubnetMask, net.Gateway)
	} else {
		cmdOptions = fmt.Sprintf("sudo connmanctl config %s --ipv4 dhcp", iface)
	}
	_, err = cmd.Builder(cmdOptions).RunCommand(options...)
	if err != nil {
		return true, err
	}
	//inst.CMD.Commands = command.Builder(cmd)
	//res := inst.CMD.RunCommand()

	return

}

func (inst *UnixCommand) edge28Iface(options ...command.Options) (interfaceName string, err error) {

	res, err := cmd.Builder("connmanctl services").RunCommand(options...)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", errors.New("failed to get interface")
	} else {
		if strings.Contains(res.Out, "*AO") {
			interfaceName = strings.ReplaceAll(res.Out, "*AO Wired", "")
			return StandardizeSpaces(interfaceName), nil
		}
		if strings.Contains(res.Out, "*AR") {
			interfaceName = strings.ReplaceAll(res.Out, "*AR Wired", "")
			return StandardizeSpaces(interfaceName), nil
		} else {
			return "", errors.New("failed to parse interface")
		}

	}
}

//StandardizeSpaces will remove all extra white spaces in text but will leave one white space between a word or letter
func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
