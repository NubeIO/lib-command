package product

type ProductType int

//go:generate stringer -type=ProductType
const (
	RubixCompute ProductType = iota
	RubixComputeIO
	RubixCompute5
	Edge28
	Nuc
	Server
	AllLinux
	Mac
	None
)
