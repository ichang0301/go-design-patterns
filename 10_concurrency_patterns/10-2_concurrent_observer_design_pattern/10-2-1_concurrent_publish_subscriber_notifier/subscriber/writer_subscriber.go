package subscriber

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type writerSubscriber struct {
	id     int
	in     chan any
	Writer io.Writer
}

func (s *writerSubscriber) Notify(msg any) error {
	var err error

	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()

	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = errors.New("Timeout")
	}

	return err
}

func (s *writerSubscriber) Close() {
	close(s.in)
}

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}

	s := &writerSubscriber{
		id:     id,
		in:     make(chan any),
		Writer: out,
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()

	return s
}
