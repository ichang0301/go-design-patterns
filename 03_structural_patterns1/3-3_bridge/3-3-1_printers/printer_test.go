package bridge

import (
	"errors"
	"strings"
	"testing"
)

const (
	errorIncorrectErrorMessageFormat = `Error message was not correct.
Actual: %s
Expected: %s`
	errorDidNotPassWriterFormat = `API2 did not write correctly on the io.Writer.
Actual: %s
Expected: %s`
	errorDoesNotMatchErrorMessageFormat = `The expected message on the io.Writer doesn't match actual.
Actual: %s
Expected: %s`
)

func TestPrintAPI1(t *testing.T) {
	api1 := &PrinterImpl1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: %v\n", err)
	}
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}

	err = errors.New("Content received on Writer was empty")
	return
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}

	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrorMessage := "you need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf(errorIncorrectErrorMessageFormat, err.Error(), expectedErrorMessage)
		}
	}

	testWriter := TestWriter{}
	api2.Writer = &testWriter

	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %v\n", err)
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf(errorDidNotPassWriterFormat, testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	normal := NormalPrinter{
		Message: expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Error(err)
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		Message: expectedMessage,
		Printer: &PrinterImpl2{Writer: &testWriter},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err)
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf(errorDoesNotMatchErrorMessageFormat, testWriter.Msg, expectedMessage)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: Hello io.Writer"

	packt := PacktPrinter{
		Message: passedMessage,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Error(err)
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{
		Message: passedMessage,
		Printer: &PrinterImpl2{Writer: &testWriter},
	}

	err = packt.Print()
	if err != nil {
		t.Error(err)
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf(errorDoesNotMatchErrorMessageFormat, testWriter.Msg, expectedMessage)
	}
}
