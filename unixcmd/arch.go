package unixcmd

import (
	"errors"
	"github.com/NubeIO/lib-command/command"
	"runtime"
	"strings"
)

type Arch struct {
	ArchModel    string `json:"arch_model"`
	IsBeagleBone bool   `json:"is_beagle_bone,omitempty"`
	IsRaspberry  bool   `json:"is_raspberry,omitempty"`
	IsArm        bool   `json:"is_arm,omitempty"`
	IsAMD64      bool   `json:"is_amd64,omitempty"`
	IsAMD32      bool   `json:"is_amd32,omitempty"`
	IsARMf       bool   `json:"is_armf,omitempty"`
	IsARMv7l     bool   `json:"is_armv7l,omitempty"`
	IsLinux      bool   `json:"is_linux"`
	Err          error
}

type ArchCheck struct {
	Windows bool
	Linux   bool
	Darwin  bool
}

func (inst *UnixCommand) ArchCheck() (arch ArchCheck) {
	s := runtime.GOOS
	switch s {
	case "linux":
		arch.Linux = true
		return arch
	case "windows":
		arch.Windows = true
		return arch
	case "darwin":
		arch.Darwin = true
		return arch
	}
	return arch
}

func (inst *UnixCommand) ArchIsLinux() bool {
	s := runtime.GOOS
	switch s {
	case "linux":
		return true
	}
	return false
}

//DetectArch can detect hardware type is in ARM or AMD
func (inst *UnixCommand) DetectArch(options ...command.Options) (arch *Arch, err error) {
	arch = &Arch{}
	res, err := cmd.Builder("dpkg", "--print-architecture").RunCommand(options...)
	if err != nil {
		return nil, err
	}
	cmdOut := res.Out
	if strings.Contains(cmdOut, "amd64") {
		arch.ArchModel = cmdOut
		arch.IsAMD64 = true
		return arch, nil
	} else if strings.Contains(cmdOut, "amd32") {
		arch.ArchModel = cmdOut
		arch.IsAMD32 = true
		return arch, nil
	} else if strings.Contains(cmdOut, "armhf") {
		arch.ArchModel = cmdOut
		arch.IsARMf = true
		arch.IsArm = true
		return arch, nil
	} else if strings.Contains(cmdOut, "armv7l") {
		arch.ArchModel = cmdOut
		arch.IsARMv7l = true
		arch.IsArm = true
		return arch, nil
	}
	return arch, errors.New("could not find correct arch type")
}

type NubeProduct struct {
	IsRC, IsEdge bool
	Name         string
}

//DetectNubeProduct can detect hardware type is in ARM or AMD and also if hardware is for example a Raspberry PI
func (inst *UnixCommand) DetectNubeProduct(options ...command.Options) (*NubeProduct, error) {
	out := &NubeProduct{}
	res, err := cmd.Builder("cat", "/proc/device-tree/model").RunCommand(options...)
	if err != nil {
		arch, _ := inst.DetectArch()
		if arch == nil {
			return nil, errors.New("unable to check arch type")
		}
		out.Name = arch.ArchModel
		return out, nil
	}
	cmdOut := res.Out
	if strings.Contains(cmdOut, "Raspberry Pi") {
		out.IsRC = true
		out.Name = "rubix-compute"
		return out, err
	} else if strings.Contains(cmdOut, "BeagleBone Black") {
		out.IsEdge = true
		out.Name = "edge-28"
		return out, err
	}
	return out, err
}

func (inst *UnixCommand) CheckEdge28() error {
	out, _ := inst.DetectNubeProduct()
	if out.IsEdge {
	} else {
		return errors.New("the host product is not type edge-28")
	}
	return nil

}
