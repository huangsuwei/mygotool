package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func WaitForSignals() {
	sigChan := make(chan os.Signal)

	signal.Notify(
		sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGINT:
			fmt.Println("received SIGINT...")
			return
		case syscall.SIGTERM:
			fmt.Println("received SIGTERM...")
			return
		}
	}
}
