package service

import (
	"sync"
	"time"
)

/**
todo 互斥锁如何实现公平？？

互斥锁有两种状态：正常状态和饥饿状态。

在正常状态下，所有等待锁的 goroutine 按照FIFO顺序等待。唤醒的 goroutine 不会直接拥有锁，而是会和新请求锁的 goroutine 竞争锁的拥有。新请求锁的 goroutine 具有优势：它正在 CPU 上执行，而且可能有好几个，所以刚刚唤醒的 goroutine 有很大可能在锁竞争中失败。在这种情况下，这个被唤醒的 goroutine 会加入到等待队列的前面。 如果一个等待的 goroutine 超过 1ms 没有获取锁，那么它将会把锁转变为饥饿模式。

在饥饿模式下，锁的所有权将从 unlock 的 goroutine 直接交给交给等待队列中的第一个。新来的 goroutine 将不会尝试去获得锁，即使锁看起来是 unlock 状态, 也不会去尝试自旋操作，而是放在等待队列的尾部。

如果一个等待的 goroutine 获取了锁，并且满足一以下其中的任何一个条件：(1)它是队列中的最后一个；(2)它等待的时候小于1ms。它会将锁的状态转换为正常状态。

正常状态有很好的性能表现，饥饿模式也是非常重要的，因为它能阻止尾部延迟的现象。
*/

type RW interface {
	Write()
	Read()
}

const cost = time.Microsecond //1 微秒(百万分之一秒)

type Lock struct {
	count int
	mu    sync.Mutex //互斥锁，完全互斥，  读与读，写与读，写与写，完全互斥
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost) //模拟耗时  1微妙
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}

type RWLock struct {
	count int
	mu    sync.RWMutex //读写锁，分为读锁与写锁，   读锁，可以多个协程重复加，读锁与读锁不互斥，读锁与写锁互斥；写锁，多个协程不可以重复加，写锁与写锁互斥，并且写锁与读锁互斥
}

func (rwl *RWLock) Write() {
	rwl.mu.Lock() //加写锁
	rwl.count++
	time.Sleep(cost)
	rwl.mu.Unlock()
}

func (rwl *RWLock) Read() {
	rwl.mu.RLock() //加读锁
	_ = rwl.count
	time.Sleep(cost)
	rwl.mu.RUnlock()
}
