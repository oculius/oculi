package di

import (
	"sync"
)

var (
	once     sync.Once
	instance Storage
)

func Instance() Storage {
	once.Do(func() {
		if instance != nil {
			panic("dependency injection storage: instance is not nil")
		}
		instance = newStorage()
	})
	return instance
}
