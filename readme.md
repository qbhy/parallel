# Parallel

## 安装 - installation
```bash
go get github.com/qbhy/parallel
```

## 使用 - usage
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
> 也可以参考 `tests/parallel_test.go` 的代码

https://github.com/qbhy/parallel  
qbhy0715@qq.com
