package product

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	got, err := Get()
	fmt.Println(got, err)

	got, err = Get("/data/product.jso")
	fmt.Println(got, err)
}
