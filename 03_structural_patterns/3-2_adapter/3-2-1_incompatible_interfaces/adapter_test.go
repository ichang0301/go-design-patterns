package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	// Table Driven Test
	cases := []struct {
		name        string
		oldPrinter  LegacyPrinter
		expectedMsg string
	}{
		{
			name:        "should_return_message_with_prefix_when_using_MyLegacyPrinter",
			oldPrinter:  &MyLegacyPrinter{},
			expectedMsg: "Legacy Printer: Adapter: Hello World!\n",
		},
		{
			name:        "should_return_pure_message_without_LegacyPrinter",
			oldPrinter:  nil,
			expectedMsg: "Hello World!",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			adapter := PrintAdapter{OldPrinter: c.oldPrinter, Msg: "Hello World!"}
			returnedMsg := adapter.PrintStored()

			if returnedMsg != c.expectedMsg {
				t.Errorf("Message didn't match: %s\n", returnedMsg)
			}
		})
	}
}
