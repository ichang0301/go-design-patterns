package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	return nil
}

func (s *singleton) AddOne() int {
	return 0
}
