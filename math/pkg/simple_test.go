package pkg

import (
	"fmt"
	"testing"
)

func TestGenerateSimpleMathIntAddition(t *testing.T) {
	generateSimpleMathIntAddition(100)
}

func TestGenerateSimpleMathDecimal(t *testing.T) {
	f := generateSimpleMathDecimal(0)
	fmt.Println("generate float: ", f)
}
