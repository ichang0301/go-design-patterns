package main

import (
	"fmt"
	"regexp"
	"sync"
	"testing"

	"github.com/ichang0301/go-design-patterns/10_concurrency_patterns/10-1_workers_pool/10-1-1_workers_pool/workers"
)

func TestMain(t *testing.T) {
	bufferSize := 100
	var dispatcher workers.Dispatcher = workers.NewDispatcher(bufferSize)

	workersCount := 3
	for i := range workersCount {
		var wl workers.WorkerLauncher = &workers.PrefixSuffixWorker{
			Id:      i,
			PrefixS: fmt.Sprintf("WorkerID: %d", i),
			SuffixS: " World",
		}
		dispatcher.LaunchWorker(wl)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)

	for i := range requests {
		req := workers.Request{
			Data: fmt.Sprintf("(Msg_id: %d) -> Hello", i),
			Handler: func(i interface{}) {
				defer wg.Done()

				s, ok := i.(string)
				if !ok {
					t.Fail()
				}

				ok, err := regexp.Match(`WorkerID\: \d* -\> \(MSG_ID: \d*\) -> [A-Z]*\sWorld`, []byte(s))
				if !ok || err != nil {
					t.Fail()
				}
			},
		}
		dispatcher.MakeRequest(req)
	}
}
