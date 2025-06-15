package future

import "fmt"

type ExecuteStringFunc func() (string, error)
type SuccessFunc func(string)
type FailFunc func(error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (s *MaybeString) Success(fn SuccessFunc) *MaybeString {
	s.successFunc = fn
	return s
}

func (s *MaybeString) Fail(fn FailFunc) *MaybeString {
	s.failFunc = fn
	return s
}

func (s *MaybeString) Execute(fn ExecuteStringFunc) *MaybeString {
	go func(s *MaybeString) {
		str, err := fn()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
	return s
}

func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure\n", msg)
	return func() (string, error) {
		return msg, nil
	}
}
