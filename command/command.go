package command

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type Command struct {
	ShellToUse string
	SetPath    string //set working dir
	Commands   []string
}

func New(cmd *Command) *Command {
	return cmd
}

type Response struct {
	Ok             bool
	Out            string
	OutByte        []byte
	Commands       []string
	CommandsJoined string
}

func (inst *Command) Builder(args ...string) *Command {
	inst.Commands = args
	return inst
}

type Options int

const (
	Debug  Options = iota //log outputs and command
	DryRun Options = iota //used for just getting the commands that have been pass in and don't run the command
	Run    Options = iota // will just run by default
)

func (inst *Command) RunCommand(args ...Options) (res *Response, err error) {
	res = &Response{}
	if len(inst.Commands) <= 0 {
		err = fmt.Errorf("no command provided")
		return res, err
	}
	shell := inst.ShellToUse //bash -c, "/usr/bin/ls"
	if shell == "" {
		shell = inst.Commands[0]
	}
	debug := false
	if len(args) > 0 {
		for _, arg := range args {
			switch arg {
			case Debug:
				debug = true
			case DryRun:
				res.Commands = inst.Commands
				debug = true
				res.CommandsJoined = strings.Join(res.Commands[:], " ")
				return res, nil
			default:
			}
		}
	}

	if debug {
		log.Infoln("command to run:", exec.Command(shell, inst.Commands[1:]...).String())
	}
	cmd := exec.Command(shell, inst.Commands[1:]...)
	output, err := cmd.CombinedOutput()
	outAsString := strings.TrimRight(string(output), "\n")
	if err != nil {
		log.Errorln("command err:", fmt.Sprint(err)+": "+outAsString)
		return
	}
	cmd.Dir = inst.SetPath
	res.Out = outAsString
	res.Ok = true
	res.OutByte = output
	return res, nil
}
