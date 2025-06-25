package publisher

import "github.com/ichang0301/go-design-patterns/10_concurrency_patterns/10-2_concurrent_observer_design_pattern/10-2-1_concurrent_publish_subscriber_notifier/subscriber"

type Publisher interface {
	start()
	AddSubscriberCh() chan<- subscriber.Subscriber
	RemoveSubscriberCh() chan<- subscriber.Subscriber
	PublishingCh() chan<- any
	Stop()
}

// The publisher has proxying functions such as AddSubscriberCh, RemoveSubscriberCh, PublishingCh, and Stop to separate the complexity of the concurrent structure
type publisher struct {
	subscribers        []subscriber.Subscriber
	addSubscriberCh    chan subscriber.Subscriber
	removeSubscriberCh chan subscriber.Subscriber
	in                 chan any
	stop               chan struct{}
}

func (p *publisher) AddSubscriberCh() chan<- subscriber.Subscriber {
	return p.addSubscriberCh
}

func (p *publisher) RemoveSubscriberCh() chan<- subscriber.Subscriber {
	return p.removeSubscriberCh
}

func (p *publisher) PublishingCh() chan<- any {
	return p.in
}

func (p *publisher) Stop() {
	close(p.stop)
}

func (p *publisher) start() {
	for {
		select {
		case msg := <-p.in:
			for _, sub := range p.subscribers {
				sub.Notify(msg)
				// go sub.Notify(msg)	// It can be a race condition if the channel is closed
			}
		case sub := <-p.addSubscriberCh:
			p.subscribers = append(p.subscribers, sub)
		case sub := <-p.removeSubscriberCh:
			for i, candidate := range p.subscribers { // O(N) complexity
				if candidate == sub {
					p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
					candidate.Close()
					break
				}
			}
		case <-p.stop:
			for _, sub := range p.subscribers {
				sub.Close()
			}
			close(p.addSubscriberCh)
			close(p.in)
			close(p.removeSubscriberCh)

			return
		}
	}
}

func NewPublisher() Publisher {
	p := &publisher{
		addSubscriberCh:    make(chan subscriber.Subscriber),
		removeSubscriberCh: make(chan subscriber.Subscriber),
		in:                 make(chan any),
		stop:               make(chan struct{}),
	}
	go p.start()
	return p
}
