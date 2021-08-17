package tests

import (
	"errors"
	"fmt"
	"parallel"
	"testing"
)

func TestParallel(t *testing.T) {
	p := parallel.NewParallel(10)

	p.Add(func() interface{} {
		fmt.Println("执行了")
		return nil
	})
	p.Add(func() interface{} {
		panic(errors.New("报错了"))
	})

	fmt.Println(p.Wait())
}
