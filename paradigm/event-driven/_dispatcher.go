package event_driven

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type (
	EventManager interface {
		AddEventHandler(priority int, handler any)
		RemoveEventHandlers(event any, selector func(listener Listener) bool) int
		GetEventHandler(event any, name string) *Listener
		GetEventHandlers(event any) Listeners

		Dispatch(event any) int64
		DispatchAsync(event any, parallel bool) <-chan int64
		fmt.Stringer
	}

	eventManager struct {
		lock           sync.RWMutex
		eventListeners map[string]Listeners
	}
)

func (e *eventManager) String() string {
	return fmt.Sprintf("%+v", e.eventListeners)
}

func (e *eventManager) GetEventHandlers(event any) Listeners {
	e.lock.RLock()
	defer e.lock.RUnlock()

	eventName := e.getStructName(event)
	if _, ok := e.eventListeners[eventName]; !ok {
		return nil
	}
	return e.eventListeners[eventName]
}

func NewManager() EventManager {
	return &eventManager{eventListeners: map[string]Listeners{}}
}

func (e *eventManager) AddEventHandler(priority int, handler any) {
	e.lock.Lock()
	defer e.lock.Unlock()

	listener := NewListener(priority, handler)
	listeners := e.eventListeners[listener.EventName]
	listeners = append(listeners, listener)
	SortListeners(listeners)
	e.eventListeners[listener.EventName] = listeners
}

func (e *eventManager) getStructName(event any) string {
	eventVal := reflect.TypeOf(event)
	if eventVal.Kind() == reflect.Ptr {
		eventVal = reflect.ValueOf(event).Elem().Type()
	} else if eventVal.Kind() == reflect.String {
		return event.(string)
	}
	if eventVal.Kind() != reflect.Struct {
		panic("event should be a struct")
	}
	return eventVal.Name()
}

func (e *eventManager) RemoveEventHandlers(event any, selector func(listener Listener) bool) int {
	e.lock.Lock()
	defer e.lock.Unlock()

	eventName := e.getStructName(event)
	if _, ok := e.eventListeners[eventName]; !ok {
		return 0
	}
	count := 0
	var newListeners []*Listener
	for _, each := range e.eventListeners[eventName] {
		if !selector(*each) {
			newListeners = append(newListeners, each)
		} else {
			count++
		}
	}
	if len(newListeners) == 0 {
		delete(e.eventListeners, eventName)
	} else {
		SortListeners(newListeners)
		e.eventListeners[eventName] = newListeners
	}
	return count
}

func (e *eventManager) GetEventHandler(event any, name string) *Listener {
	e.lock.RLock()
	defer e.lock.RUnlock()

	eventName := e.getStructName(event)
	if _, ok := e.eventListeners[eventName]; !ok {
		return nil
	}
	for _, each := range e.eventListeners[eventName] {
		if strings.EqualFold(each.ListenerName, name) {
			return each
		}
	}
	return nil
}

func (e *eventManager) getProperties(event any) (reflect.Value, reflect.Value) {
	var eventPointerValue reflect.Value
	var eventValue reflect.Value

	eventValue = reflect.ValueOf(event)
	if eventValue.Kind() == reflect.Ptr && eventValue.Elem().Kind() == reflect.Struct {
		eventValue = eventValue.Elem()
		eventPointerValue = reflect.ValueOf(event)
	} else {
		ptrEventStruct := reflect.New(eventValue.Type()).Interface() // *Event
		eventPointerValue = reflect.ValueOf(ptrEventStruct)          // reflect.Value(*Event)
		eventPointerValue.Elem().Set(eventValue)                     // assign value
		eventValue = eventPointerValue.Elem()                        // pointer and non pointer type must be the same struct
	}

	return eventValue, eventPointerValue
}

func (e *eventManager) Dispatch(event any) int64 {
	e.lock.RLock()
	defer e.lock.RUnlock()

	eventValue, eventPointerValue := e.getProperties(event)

	eventName := e.getStructName(eventValue.Interface())
	if _, ok := e.eventListeners[eventName]; !ok {
		return 0
	}

	return e.run(eventName, eventValue, eventPointerValue)
}

func (e *eventManager) run(eventName string, eventValue, eventPointerValue reflect.Value) int64 {
	handlers := int64(0)
	for _, handler := range e.eventListeners[eventName] {
		if !handler.Valid() {
			continue
		}

		handlers++
		handlerVal := reflect.ValueOf(handler.Fn)
		if handler.IsPointer {
			handlerVal.Call([]reflect.Value{eventPointerValue})
		} else {
			handlerVal.Call([]reflect.Value{eventValue})
		}
	}
	return handlers
}

func (e *eventManager) DispatchAsync(event any, parallel bool) <-chan int64 {
	result := make(chan int64, 1)
	go func() {
		e.lock.RLock()
		defer e.lock.RUnlock()

		eventValue, eventPointerValue := e.getProperties(event)

		eventName := e.getStructName(eventValue.Interface())
		if _, ok := e.eventListeners[eventName]; !ok {
			result <- 0
			return
		}

		handlers := int64(0)
		isPossible := true
		for _, handler := range e.eventListeners[eventName] {
			if handler.IsPointer {
				isPossible = false
				break
			}
		}

		if !isPossible || !parallel {
			// Async linear
			handlers = e.run(eventName, eventValue, eventPointerValue)
		} else {
			var wg sync.WaitGroup
			// Async parallel
			listeners := e.eventListeners[eventName]
			for _, handler := range listeners {
				if !handler.Valid() {
					continue
				}
				handlerVal := reflect.ValueOf(handler.Fn)
				wg.Add(1)
				handlers++
				go func() {
					handlerVal.Call([]reflect.Value{eventValue})
					wg.Done()
				}()
			}
			wg.Wait()
		}
		result <- handlers
		return
	}()
	return result
}
