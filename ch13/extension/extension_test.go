package extension_test

import (
	"fmt"
	"testing"
)

type Code string

type Programer interface {
	WriteHelloWorld() Code
}

type GoProgramer struct {
}

func (g *GoProgramer) WriteHelloWorld() Code {
	return "fmt.Print(\"Go Hello World\")"
}

type PhpProgramer struct {
}

func (g *PhpProgramer) WriteHelloWorld() Code {
	return "fmt.Print(\"PHP Hello World\")"
}

func WriteFirstProgramer(p Programer) {
	fmt.Print(p.WriteHelloWorld())
}

func TestInterface(t *testing.T) {
	var g = new(GoProgramer)
	var p = new(PhpProgramer)
	WriteFirstProgramer(g)
	WriteFirstProgramer(p)
}
