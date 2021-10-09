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
		time.Sleep(time.Second * 2)
		return "执行了"
	})
	p.Add(func() interface{} {
		time.Sleep(time.Second * 2)
		return "执行了"
	})
	p.Add(func() interface{} {
		panic(errors.New("报错了"))
	})
	p.Add(func() interface{} {
		time.Sleep(time.Second * 5)
		return "执行了"
	})

	fmt.Println(p.Wait())
}

// 监听模式
func TestParallelListen(t *testing.T) {
	p := parallel.NewParallel(2)

	go func() {
		i := 0
		for ; i < 10; i++ {
			(func(i int) {
				result := p.Add(func() interface{} {
					time.Sleep(time.Second)
					fmt.Printf("每隔1秒执行一次 %d \n", i)
					return nil
				})

				fmt.Printf("添加结果: %v, i: %d \n", result, i)
			})(i)
		}

		ticker := time.NewTicker(time.Second)

		for {
			<-ticker.C
			i++
			(func(i int) {
				p.Add(func() interface{} {
					time.Sleep(time.Second)
					fmt.Println("每隔1秒执行一次", i)
					return nil
				})
			})(i)

			if i > 20 {
				p.Stop()
				break
			}
		}
	}()

	fmt.Println(p.Listen())
}

// 测试优雅退出
func TestParallelGracefulStop(t *testing.T) {
	p := parallel.NewParallel(2)

	go func() {
		i := 0
		for ; i < 10; i++ {
			(func(i int) {
				if i >= 5 {
					p.GracefulStop()
				}
				result := p.Add(func() interface{} {
					time.Sleep(time.Second)
					fmt.Printf("每隔1秒执行一次 %d \n", i)
					return nil
				})

				fmt.Printf("添加结果: %v, i: %d\n", result, i)
			})(i)
		}
	}()

	fmt.Println(p.Listen())
}
