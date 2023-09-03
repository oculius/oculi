package event_driven

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
)

type (
	Listeners []*Listener
	Listener  struct {
		ListenerName string
		EventName    string
		Priority     int
		Fn           any
		IsPointer    bool
		eventId      int
	}
)

var counter = 1

func SortListeners(listener Listeners) {
	sort.Slice(listener, func(i, j int) bool {
		if listener[i].Priority == listener[j].Priority {
			return listener[i].eventId < listener[j].eventId
		}
		return listener[i].Priority > listener[j].Priority
	})
}

func (l *Listener) Valid() bool {
	return l.eventId != 0
}

func (l *Listener) String() string {
	return fmt.Sprintf("%+v", *l)
}

func (l Listeners) String() string {
	copied := make([]Listener, len(l))
	for i, each := range l {
		copied[i] = *each
	}
	return fmt.Sprintf("%+v", copied)
}

func NewListener(priority int, handler any) *Listener {
	defer func() {
		counter++
	}()
	fnValue := reflect.ValueOf(handler)

	// Check if fn is a function
	if fnValue.Kind() != reflect.Func {
		panic("provided value is not a function, listener should be a function")
	}

	// Check number of parameters
	numIn := fnValue.Type().NumIn()
	if numIn != 1 {
		panic("listener function should have exactly 1 parameter")
	}

	// Check return values
	numOut := fnValue.Type().NumOut()
	if numOut != 0 {
		panic("listener function should not have any return values")
	}

	// Get the struct name from the first parameter
	isPointer := false
	firstParamType := fnValue.Type().In(0)
	if firstParamType.Kind() == reflect.Ptr && firstParamType.Elem().Kind() == reflect.Struct {
		firstParamType = firstParamType.Elem()
		isPointer = true
	}
	if firstParamType.Kind() != reflect.Struct {
		panic("listener function's parameter should be an event struct")
	}

	return &Listener{
		ListenerName: runtime.FuncForPC(fnValue.Pointer()).Name(),
		EventName:    firstParamType.Name(),
		IsPointer:    isPointer,
		Priority:     priority,
		Fn:           handler,
		eventId:      counter,
	}
}
