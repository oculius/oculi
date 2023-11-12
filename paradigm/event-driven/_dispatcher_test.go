package event_driven

import (
	"fmt"
	"strings"
	"testing"
)

type PlayerJoinEvent struct {
	PlayerName string
}

type PlayerLeaveEvent struct {
	PlayerName string
}

type EventASD interface {
}

func TestNewDispatcher(t *testing.T) {
	t.Run("test", func(tt *testing.T) {
		dispatcher := NewManager()
		dispatcher.AddEventHandler(1, func(event EventASD) {

		})
	})
	t.Run("general", func(tt *testing.T) {
		dispatcher := NewManager()

		// Adding event handlers
		dispatcher.AddEventHandler(0, func(event EventASD) {
			fmt.Println("r1, Player joined:", event)
		})
		// Adding event handlers
		dispatcher.AddEventHandler(0, func(event *PlayerJoinEvent) {
			fmt.Println("r2, Player joined:", event.PlayerName)
		})
		// Adding event handlers
		dispatcher.AddEventHandler(0, func(event PlayerJoinEvent) {
			fmt.Println("r3, pointered, Player joined:", event.PlayerName)
			event.PlayerName = "Johny"
		})

		dispatcher.AddEventHandler(1, func(event PlayerLeaveEvent) {
			fmt.Println("Player left:", event.PlayerName)
		})
		fmt.Println("--")
		fmt.Println(dispatcher.GetEventHandlers(EventASD(PlayerJoinEvent{})))
		// Adding event handlers
		dispatcher.RemoveEventHandlers("PlayerLeaveEvent", func(listener Listener) bool {
			return strings.HasSuffix(listener.ListenerName, "func1.4")
		})
		dispatcher.AddEventHandler(0, func(event PlayerJoinEvent) {
			fmt.Println("r4, pointered, Player joined:", event.PlayerName)
		})

		// Dispatching events
		//playerJoinEvent := &PlayerJoinEvent{PlayerName: "John"}
		//playerLeaveEvent := &PlayerLeaveEvent{PlayerName: "Jane"}
		itf := &PlayerJoinEvent{"HAHA"}

		dispatcher.Dispatch(EventASD(itf))
		//dispatcher.DispatchAsync(playerJoinEvent, false)
		//fmt.Println("TESTASD")
		//dispatcher.DispatchAsync(playerLeaveEvent, false)
		//fmt.Println("TESTASD")
		//playerJoinEvent2 := PlayerJoinEvent{"HAHA"}
		//dispatcher.DispatchAsync(playerJoinEvent2, false)
		//fmt.Println(playerJoinEvent, playerJoinEvent2)
		//fmt.Println("TESTASD")
		//time.Sleep(1 * time.Second)
	})
}
