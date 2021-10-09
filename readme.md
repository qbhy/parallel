# Parallel
golang 协程并行库，可以指定并发数量

## 安装 - installation
```bash
go get github.com/qbhy/parallel
```

## 使用 - usage
1. 普通用法
    ```go
    package tests
    
    import (
        "errors"
        "fmt"
        "github.com/qbhy/parallel"
        "testing"
    )
    
    func TestParallel(t *testing.T) {
        // 最多 10 个协程同时运行
        p := parallel.NewParallel(10)
    
        p.Add(func() interface{} {
            return "执行了"
        })
        
        p.Add(func() interface{} {
            panic(errors.New("报错了"))
        })
    
        fmt.Println(p.Wait())
        //会输出 map[0:执行了 1:报错了]
    }
    ```
2. 简单的队列
    ```go
    package tests
    
    import (
        "fmt"
        "github.com/qbhy/parallel"
        "testing"
        "time"
    )
    
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
    ```
    > 也可以参考 `tests/parallel_test.go` 的代码

https://github.com/qbhy/parallel  
qbhy0715@qq.com
