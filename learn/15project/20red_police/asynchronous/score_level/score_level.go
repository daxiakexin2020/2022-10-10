package score_level

import (
	"20red_police/asynchronous"
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)

type scoreLevel struct {
	itask     chan *bus
	capacity  int
	workerNum int
	boot      bool
	callback  func(string, int64) error
}

type bus struct {
	username string
	score    int64
}

var (
	_  asynchronous.Tasker = (*scoreLevel)(nil)
	ec                     = make(chan struct{}, 1)
)

var (
	gscoreLevel *scoreLevel
	slonce      sync.Once
)

const (
	limit          = 0.9
	max_worker_num = 1000
)

func newBus(username string, score int64) *bus {
	return &bus{username: username, score: score}
}

func ScoreLevel(capacity int, workerNum int, callback func(username string, score int64) error) *scoreLevel {
	slonce.Do(func() {
		gscoreLevel = newScoreLevel(capacity, workerNum, callback)
	})
	return gscoreLevel
}

func GScoreLevel() *scoreLevel {
	return gscoreLevel
}

func newScoreLevel(capacity int, workerNum int, callback func(username string, score int64) error) *scoreLevel {
	if workerNum > max_worker_num {
		workerNum = max_worker_num
	}
	return &scoreLevel{
		capacity:  capacity,
		workerNum: workerNum,
		itask:     make(chan *bus, capacity),
		callback:  callback,
	}
}

func (sl *scoreLevel) TaskName() string {
	return "SCORE_LEVEL"
}

func (sl *scoreLevel) Run() error {
	if sl.boot {
		return nil
	}
	sl.boot = true
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 0; i < sl.workerNum; i++ {
		go sl.worker(ctx)
	}
	log.Println("Score Level is runing.........................")
	for {
		if !sl.boot && len(sl.itask) == 0 {
			log.Println("all goroutine need stop,send signal:::::::::::::::::", sl.boot, len(sl.itask))
			cancelFunc()
			return nil
		}
	}
	return nil
}

func (sl *scoreLevel) worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("receive close signal, stoping......")
			return
		case task := <-sl.itask:
			if sl.callback != nil {
				sl.callback(task.username, task.score)
			}
		default:
		}
	}
}

func (sl *scoreLevel) Stop() error {
	if !sl.boot {
		return nil
	}
	sl.boot = false
	return nil
}

func (sl *scoreLevel) ExitSignal() chan struct{} {
	return ec
}

func (sl *scoreLevel) Add(username string, score int64) error {
	if !sl.boot {
		return errors.New("ScoreLevel is stoped")
	}
	if float32(len(sl.itask)) >= float32(sl.capacity)*limit {
		return errors.New("ScoreLevel is Beyond the limit")
	}
	b := newBus(username, score)
	sl.itask <- b
	return nil
}
