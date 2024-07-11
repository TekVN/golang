package os

import (
	"os"
	"os/signal"
)

func RegisterSignalHandler(fn func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}

	go WaitSignalHandler(fn, signals...)
}

func WaitSignalHandler(fn func(), signals ...os.Signal) {
	if len(signals) == 0 {
		return
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	<-ch

	fn()
}
