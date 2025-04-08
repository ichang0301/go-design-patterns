package adapter

import "fmt"

type PrintAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrintAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}

	return
}
