package utils

import  "sync"

type Singleton struct {
	once sync.Once
	obj interface{}
	fn func() interface{}
}

func NewOneSingleton(fn func() interface{}) *Singleton {
	return &Singleton{
		fn: fn,
	}
}

func (s *Singleton) Get() interface{} {
	s.once.Do(func() {
		s.obj = s.fn()
	})
	return s.obj
}