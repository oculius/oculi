package di

import (
	"go.uber.org/fx"
	"sync"
)

type (
	storage struct {
		opts []fx.Option
		sync.RWMutex
	}

	Storage interface {
		Add(opts []fx.Option)
		Clear()
		Build() []fx.Option
	}
)

func (s *storage) Add(opts []fx.Option) {
	s.Lock()
	defer s.Unlock()

	if len(opts) > 0 {
		s.opts = append(s.opts, opts...)
	}
}

func (s *storage) Clear() {
	s.Lock()
	defer s.Unlock()

	s.opts = nil
}

func (s *storage) Build() []fx.Option {
	s.RLock()
	defer s.RUnlock()

	X := len(s.opts)
	result := make([]fx.Option, X)
	for i, each := range s.opts {
		result[i] = each
	}
	return result
}
