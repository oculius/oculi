package di

import (
	"sync"
)

var (
	once     sync.Once
	instance *storage
)

func getInstance() *storage {
	once.Do(func() {
		if instance != nil {
			panic("dependency injection storage: instance is not nil. perhaps race condition occur?")
		}
		instance = newStorage()
	})
	return instance
}
