package main

import (
	"fmt"
	"sync"

	"github.com/ichang0301/go-design-patterns/10_concurrency_patterns/10-1_workers_pool/10-1-1_workers_pool/workers"
)

func main() {
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
	defer wg.Wait()

	for i := range requests {
		req := workers.NewStringRequest(fmt.Sprintf("(Msg_id: %d) -> Hello", i), i, &wg)
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()
}
