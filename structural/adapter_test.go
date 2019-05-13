package structural

import "testing"

func TestPrinterAdapter(t *testing.T) {
	msg := "Hello world!"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnedMessage := adapter.PrintStored()
	if returnedMessage != "Legacy Printer: Hello world!" {
		t.Errorf("Legacy Message didn't match: %s\n", returnedMessage)
	}

	adapter2 := PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMessage2 := adapter2.PrintStored()
	if returnedMessage2 != "Hello world!" {
		t.Errorf("Message didn't match: %s\n", returnedMessage2)
	}

}
