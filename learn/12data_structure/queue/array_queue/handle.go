package array_queue

import "sync"

//先进先出   后进后出

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		array: make([]string, 0),
		lock:  sync.Mutex{},
	}
}

func (as *ArrayStack) Push(v string) {
	as.lock.Lock()
	defer as.lock.Unlock()
	as.array = append(as.array, v)
	as.size = as.size + 1
}

func (as *ArrayStack) Pop() string {
	as.lock.Lock()
	defer as.lock.Unlock()
	if as.IsEmpty() {
		panic("empty")
	}
	v := as.array[0]
	newArray := make([]string, as.size-1, as.size-1)
	for i := 1; i < as.size-1; i++ {
		newArray[i-1] = as.array[i]
	}
	as.array = newArray
	as.size = as.size - 1
	return v
}

func (as *ArrayStack) Top() string {
	if as.IsEmpty() {
		panic("empty")
	}
	return as.array[0]
}

func (as *ArrayStack) Size() int {
	return as.size
}

func (as *ArrayStack) IsEmpty() bool {
	return as.size == 0
}
