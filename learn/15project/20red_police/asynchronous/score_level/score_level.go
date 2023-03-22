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
}

type bus struct {
	username string
	score    int64
}

func newBus(username string, score int64) *bus {
	return &bus{username: username, score: score}
}

var (
	_  asynchronous.Tasker = (*scoreLevel)(nil)
	ec                     = make(chan struct{}, 1)
)

var (
	gscoreLevel *scoreLevel
	slonce      sync.Once
)

func ScoreLevel(capacity int, workerNum int) *scoreLevel {
	slonce.Do(func() {
		gscoreLevel = newScoreLevel(capacity, workerNum)
	})
	return gscoreLevel
}

func GScoreLevel() *scoreLevel {
	return gscoreLevel
}

func newScoreLevel(capacity int, workerNum int) *scoreLevel {
	return &scoreLevel{
		capacity:  capacity,
		workerNum: workerNum,
		itask:     make(chan *bus, capacity),
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
		go sl.worker(i, ctx)
	}
	log.Println("score level is runing")
	for {
		if !sl.boot && len(sl.itask) == 0 {
			log.Println("all goroutine need stop,send signal:::::::::::::::::", sl.boot, len(sl.itask))
			cancelFunc()
			return nil
		}
	}
	return nil
}

func (sl *scoreLevel) worker(index int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("receive close signal...")
			return
		case task := <-sl.itask:
			//handle bus
			log.Printf("index:%d handle task:%s,%d\n", index, task.username, task.score)
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
	b := newBus(username, score)
	sl.itask <- b
	return nil
}
