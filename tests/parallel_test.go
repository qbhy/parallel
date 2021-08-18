package tests

import (
	"errors"
	"fmt"
	"github.com/qbhy/parallel"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	p := parallel.NewParallel(2)

	p.Add(func() interface{} {
		time.Sleep(time.Second*2)
		return "执行了"
	})
	p.Add(func() interface{} {
		time.Sleep(time.Second*2)
		return "执行了"
	})
	p.Add(func() interface{} {
		panic(errors.New("报错了"))
	})
	p.Add(func() interface{} {
		time.Sleep(time.Second*5)
		return "执行了"
	})

	fmt.Println(p.Wait())
}
