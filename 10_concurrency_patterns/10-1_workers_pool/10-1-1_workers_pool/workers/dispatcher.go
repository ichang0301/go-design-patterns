package workers

import "time"

// the Dispatcher interface is to hide launching implementation details from the user
// and it is merely acting as a Facade design pattern.
type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(wl WorkerLauncher) {
	wl.LaunchWorker(d.inCh)
}

func (d *dispatcher) MakeRequest(req Request) {
	select {
	case d.inCh <- req: // sending operation: it will block until the request is sent successfully
	case <-time.After(time.Second * 5): // receiving operation: it will be triggered after 5 seconds if the request can't be sent successfully, and the function will return
		return
	}
}

func (d *dispatcher) Stop() {
	close(d.inCh) // When closing the incoming channel, it will provoke a chain reaction so that all workers will stop
}

func NewDispatcher(bufferSize int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, bufferSize), // buffered channel to hold requests
	}
}
