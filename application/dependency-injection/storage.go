package di

import (
	"sync"

	"go.uber.org/fx"
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

func newStorage() Storage {
	return &storage{
		RWMutex: sync.RWMutex{},
		opts:    make([]fx.Option, 0, 25),
	}
}

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
	i := 0
	for _, each := range s.opts {
		if each == nil {
			continue
		}
		result[i] = each
		i++
	}
	return result
}
