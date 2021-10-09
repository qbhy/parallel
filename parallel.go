package parallel

import (
	"errors"
	"sync"
)

type Parallel struct {
	Callbacks []func() interface{}
	channel   chan int
	status    ParallelStatus
}

type ParallelStatus int

const (
	NORMAL = iota
	LISTENING
	STOPPED
	GRACEFUL_STOP
)

var (
	StoppedError = errors.New("该 parallel 已经停止!")
)

func NewParallel(concurrent int) *Parallel {
	return &Parallel{
		Callbacks: make([]func() interface{}, 0),
		channel:   make(chan int, concurrent),
		status:    NORMAL,
	}
}

func (this *Parallel) Add(callback func() interface{}) error {
	if this.IsStopped() {
		return StoppedError
	} else {
		this.Callbacks = append(this.Callbacks, callback)
		return nil
	}
}

func (this *Parallel) IsStopped() bool {
	return this.status == STOPPED || this.status == GRACEFUL_STOP
}

func (this *Parallel) Wait() (results map[int]interface{}) {
	queues := this.Callbacks
	this.Clear()

	wg := sync.WaitGroup{}
	wg.Add(len(queues))
	resultMutex := sync.RWMutex{}

	results = map[int]interface{}{}
	for key, callback := range queues {
		this.channel <- 0
		go func(key int, callback func() interface{}) {
			// 捕捉异常
			defer func() {
				if err := recover(); err != nil {
					resultMutex.Lock()
					results[key] = err
					resultMutex.Unlock()
				}

				<-this.channel
				wg.Done()
			}()

			result := callback()

			resultMutex.Lock()
			results[key] = result
			resultMutex.Unlock()
		}(key, callback)
	}

	wg.Wait()

	return
}

func (this *Parallel) Run() map[int]interface{} {
	return this.Wait()
}

func (this *Parallel) Stop() {
	this.status = STOPPED
}

func (this *Parallel) GracefulStop() {
	this.status = GRACEFUL_STOP
}

func (this *Parallel) Listen() (err error) {
	this.status = LISTENING

	defer func() {
		if this.status == GRACEFUL_STOP {
			this.Wait()
		}
	}()

	for {
		if this.status == LISTENING {
			this.Wait()
		} else {
			break
		}
	}

	return err
}

func (this *Parallel) Clear() {
	this.Callbacks = make([]func() interface{}, 0)
}
