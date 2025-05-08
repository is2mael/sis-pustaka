package utils

import "sync"

type OnceSingleton struct {
	once sync.Once
	obj  interface{}
	fn   func() interface{}
}

func NewOnceSingleton(fn func() interface{}) *OnceSingleton {
	return &OnceSingleton{fn: fn}
}

func (s *OnceSingleton) Get() interface{} {
	s.once.Do(func() {
		s.obj = s.fn()
	})
	return s.obj
}