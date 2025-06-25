package workers

import (
	"fmt"
	"log"
	"sync"
)

type Request struct {
	Data    interface{} // It can be any data type, e.g., int, string, struct etc.
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	return Request{
		Data: s,
		Handler: func(i interface{}) {
			defer wg.Done()

			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
}
