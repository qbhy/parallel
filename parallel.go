package parallel

import "sync"

type Parallel struct {
	Callbacks []func() interface{}
	channel   chan int
}

func NewParallel(concurrent int) *Parallel {
	return &Parallel{
		Callbacks: make([]func() interface{}, 0),
		channel:   make(chan int, concurrent),
	}
}

func (p *Parallel) Add(callback func() interface{}) {
	p.Callbacks = append(p.Callbacks, callback)
}

func (p *Parallel) Wait() (results map[int]interface{}) {
	wg := sync.WaitGroup{}
	wg.Add(len(p.Callbacks))
	resultMutex := sync.RWMutex{}

	results = map[int]interface{}{}
	for key, callback := range p.Callbacks {
		p.channel <- 0
		go func(key int, callback func() interface{}) {
			// 捕捉异常
			defer func() {
				if err := recover(); err != nil {
					results[key] = err
					resultMutex.Unlock()
				}

				<-p.channel
				wg.Done()
			}()

			resultMutex.Lock()
			results[key] = callback()
			resultMutex.Unlock()
		}(key, callback)
	}

	wg.Wait()
	p.Clear()

	return
}

func (p *Parallel) Clear() {
	p.Callbacks = make([]func() interface{}, 0)
}
