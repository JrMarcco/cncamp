package homework

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func produce(ch chan<- int) {
	for {
		time.Sleep(time.Second)
		if n, ok := sendMessage(ch); ok {
			fmt.Printf("### send message [%d] to channel ###\n", n)
			continue
		}
		fmt.Println("### fail to send message, channel is full")
	}
}

func sendMessage(ch chan<- int) (int, bool) {
	n := rand.Intn(10)
	select {
	case ch <- n:
		return n, true
	default:
		return 0, false
	}
}

func consume(ch <-chan int) {
	for {
		if n, ok := receiveMessage(ch); ok {
			fmt.Printf("### receive messgae [%d] from channel ###\n", n)
			continue
		}
		fmt.Println("### fail to receive message, channel is empty")
		time.Sleep(15 * time.Second)
	}
}

func receiveMessage(ch <-chan int) (int, bool) {
	select {
	case n := <-ch:
		return n, true
	default:
		return 0, false
	}
}

func TestQueue(t *testing.T) {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}
