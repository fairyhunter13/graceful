package graceful

import (
	"os"
	"os/signal"
	"syscall"
)

type SignalChan struct {
	channel chan os.Signal
}

func New() *SignalChan {
	chanServer := make(chan os.Signal, 1)
	signal.Notify(chanServer, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	return &SignalChan{
		channel: chanServer,
	}
}

func (sc *SignalChan) Wait() {
	defer close(sc.channel)

	<-sc.channel
	signal.Stop(sc.channel)
}
