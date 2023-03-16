package room_timeout

import (
	"20red_police/asynchronous"
	"fmt"
	"testing"
	"time"
)

func TestNewListNode(t *testing.T) {
}
func TestRemoveHeadNode(t *testing.T) {
}

func TestAddNodeToTail(t *testing.T) {

}

func TestRun(t *testing.T) {
}

func TestTask(t *testing.T) {
	out := NewRoomTimeOut(100, nil)
	task := asynchronous.NewManager()
	task.Register(out)
	go task.Run()

	time.Sleep(1 * time.Second)
	out.ExitSignal() <- struct{}{}
	fmt.Println("sssssss")
	select {}
}
