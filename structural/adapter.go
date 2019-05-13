package structural

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

// by defining Print in MyLegacyPrinter , MyLegacyPrinter implements LegacyPrinter
type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s", s)
	println(newMsg)
	return newMsg
}

// adapter definition
type NewPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

// new Print method
func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = p.OldPrinter.Print(p.Msg)
	} else {
		newMsg = p.Msg
	}
	return
}
