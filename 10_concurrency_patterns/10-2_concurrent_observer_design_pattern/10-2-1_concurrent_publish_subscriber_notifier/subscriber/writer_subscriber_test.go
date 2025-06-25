package subscriber

import (
	"strings"
	"sync"
	"testing"
)

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

func TestWriterSubscriber(t *testing.T) {
	msg := "Hello"

	var wg sync.WaitGroup
	wg.Add(1)

	sub := NewWriterSubscriber(0, nil)
	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Errorf("incorrect string: %s", res)
			}
			wg.Done()
		},
	}

	err := sub.Notify(msg)
	if err != nil {
		wg.Done()
		t.Error(err)
	}

	wg.Wait()
	sub.Close()
}
