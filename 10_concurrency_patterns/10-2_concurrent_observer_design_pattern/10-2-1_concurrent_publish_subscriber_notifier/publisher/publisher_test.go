package publisher

import (
	"sync"
	"testing"
)

type mockSubscriber struct {
	notifyTestingFunc func(msg any) error
	closeTestingFunc  func()
}

func (m *mockSubscriber) Notify(msg any) error {
	return m.notifyTestingFunc(msg)
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}

func TestPublisher(t *testing.T) {
	var wg sync.WaitGroup

	sub := &mockSubscriber{
		notifyTestingFunc: func(msg any) error {
			defer wg.Done()

			s, ok := msg.(string)
			if !ok {
				t.Fatal("Could not assert result")
			}

			if s != msg {
				t.Errorf("Incorrect string: %s. Expected: %s", s, msg)
			}

			return nil // skip to check errors
		},
		closeTestingFunc: func() {
			wg.Done()
		},
	}

	// Test adding subscriber
	p := NewPublisher()
	p.AddSubscriberCh() <- sub
	wg.Add(1)

	// Test publishing a message
	msg := "Hello"
	p.PublishingCh() <- msg
	wg.Wait()

	pubCon := p.(*publisher)
	if len(pubCon.subscribers) != 1 {
		t.Errorf("Expected 1 subscriber, but got %d", len(pubCon.subscribers))
	}

	// Test removing subscriber
	wg.Add(1)
	p.RemoveSubscriberCh() <- sub
	wg.Wait()

	if len(pubCon.subscribers) != 0 {
		t.Errorf("Expected 0 subscribers, but got %d", len(pubCon.subscribers))
	}

	p.Stop()
}
