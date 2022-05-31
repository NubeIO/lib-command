package product

import (
	"errors"
	"github.com/NubeIO/lib-command/command"
	"github.com/NubeIO/lib-command/unixcmd"
	"github.com/NubeIO/lib-dirs/dirs/jparse"
	log "github.com/sirupsen/logrus"
)

type Product struct {
	Type     string `json:"type"`
	Version  string `json:"version"`
	Hardware string `json:"hardware"`
	Arch     string `json:"arch"` //amd64
}

func Get(fileAndPath ...string) (*Product, error) {
	return read(fileAndPath...)
}

const (
	FilePath = "/data/product.json"
)

var cmd = unixcmd.New(&command.Command{})

func read(fileAndPath ...string) (*Product, error) {
	path := FilePath
	if len(fileAndPath) > 0 {
		path = fileAndPath[0]
		if path == "" {
			return nil, errors.New("path can not be nil")
		}
	}
	p := &Product{}
	j := jparse.New()
	var err error
	if readErr := j.ParseToData(path, p); readErr != nil {
		log.Errorln("read-product: read from json err", readErr.Error())
		err = readErr
		return nil, readErr
	}
	res, err := cmd.DetectNubeProduct()
	if err != nil {
		return nil, err
	}
	if res.Name == "" {
		archType := cmd.ArchCheck()
		if archType.Linux {
			res.Name = "Linux"
		}
		if archType.Darwin {
			res.Name = "Darwin"
		}
	}
	resp, err := cmd.DetectArch()
	if err != nil {
		return nil, err
	}
	p.Hardware = res.Name
	p.Arch = resp.ArchModel
	return p, err
}

func CheckProduct(s string) (ProductType, error) {
	switch s {
	case RubixCompute.String():
		return RubixCompute, nil
	case RubixComputeIO.String():
		return RubixComputeIO, nil
	case Edge28.String():
		return Edge28, nil
	case AllLinux.String():
		return AllLinux, nil
	case RubixCompute5.String():
		return RubixCompute5, nil
	case Nuc.String():
		return Nuc, nil
	case Mac.String():
		return Mac, nil
	}

	return None, errors.New("invalid product type, try RubixCompute")

}
