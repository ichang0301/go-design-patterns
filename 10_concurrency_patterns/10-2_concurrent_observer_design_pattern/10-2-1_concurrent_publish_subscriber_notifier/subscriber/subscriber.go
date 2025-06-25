package subscriber

type Subscriber interface {
	Notify(any) error
	Close()
}
