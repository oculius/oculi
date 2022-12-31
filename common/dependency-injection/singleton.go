package di

import "sync"

var (
	creator  sync.Once
	instance Storage
)

func Instance() Storage {
	creator.Do(func() {
		if instance != nil {
			panic("dependency injection storage: instance is not nil")
		}
		instance = &storage{}
	})
	return instance
}
