package client

import (
	"sync"
	"testing"
)

func TestNotifySignalChanAfterDoneDoesNotPanic(t *testing.T) {
	doneChan := make(chan struct{})
	signalChan := make(chan struct{}, 1)

	close(doneChan)

	for i := 0; i < 10; i++ {
		notifySignalChan(doneChan, signalChan)
	}
}

func TestNotifySignalChanConcurrentAfterDoneDoesNotPanic(t *testing.T) {
	doneChan := make(chan struct{})
	signalChan := make(chan struct{}, 1)

	close(doneChan)

	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			notifySignalChan(doneChan, signalChan)
		}()
	}

	wg.Wait()
}

func TestForwardMessageAfterDoneReturnsFalse(t *testing.T) {
	doneChan := make(chan struct{})
	readChan := make(chan []byte)

	close(doneChan)

	if ok := forwardMessage(doneChan, readChan, []byte("payload")); ok {
		t.Fatalf("expected forwardMessage to stop after done channel closed")
	}
}
