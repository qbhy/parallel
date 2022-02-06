package tests

import (
	"fmt"
	"github.com/qbhy/parallel"
	"testing"
	"time"
)

func TestWorkers(t *testing.T) {
	workers := parallel.NewWorkers(3)
	go func() {
		for i := 0; i < 100; i++ {
			(func(i int) {
				_ = workers.Handle(func() {
					time.Sleep(1 * time.Second)
					fmt.Println("worked", i)
				})
			})(i)
		}
	}()
	time.Sleep(3 * time.Second)
	workers.Stop()
}
