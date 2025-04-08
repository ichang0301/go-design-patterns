package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(message string) error {
	fmt.Print(message)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(message string) error {
	if p.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}

	fmt.Fprint(p.Writer, message)
	return nil
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Message string
	Printer PrinterAPI
}

func (p *NormalPrinter) Print() error {
	p.Printer.PrintMessage(p.Message)
	return nil
}

type PacktPrinter struct {
	Message string
	Printer PrinterAPI
}

func (p *PacktPrinter) Print() error {
	p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", p.Message))
	return nil
}
