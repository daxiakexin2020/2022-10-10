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
	out := Timeout(100, 10, nil)
	manager := asynchronous.Manager()
	manager.Register(out)
	go manager.Run()

	time.Sleep(1 * time.Second)
	//out.ExitSignal() <- struct{}{}
	//manager.StopTask(out.TaskName())
	fmt.Println("sssssss")
	select {}
}

func TestLivetime(t *testing.T) {
	addtime := time.Now().Unix()
	livetime := time.Second * 1
	fmt.Println("time start", time.Unix(time.Now().Unix(), 0), time.Duration(1))
	time.Sleep(time.Second * 2)
	if time.Unix(addtime, 0).Add(livetime).After(time.Now()) {
		fmt.Println("time ok", time.Unix(addtime, 0).Add(livetime).Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println("time out", time.Unix(addtime, 0).Add(livetime).Format("2006-01-02 15:04:05"))
	}
}
