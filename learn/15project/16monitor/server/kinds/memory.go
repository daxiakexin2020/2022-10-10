package kinds

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type memory struct {
	percent       int
	timeInterval  time.Duration
	lastAlarmTime int64
}

const (
	MEMORY                 = "内存"
	MEMORY_DEFAULT_PERCENT = 95
)

func NewMemory(percent int, timeInterval time.Duration) *memory {
	if percent > 100 || percent < 0 {
		percent = MEMORY_DEFAULT_PERCENT
	}
	return &memory{percent: percent, timeInterval: timeInterval}
}

func (m *memory) Name() string {
	return MEMORY
}

func (m *memory) Monitor(ctx context.Context) error {

	timer := time.NewTimer(m.timeInterval * time.Second)
	go func() {
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				cmd := exec.Command("ls", "-lah")
				out, err := cmd.CombinedOutput()
				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}
				fmt.Printf("combined out:\n%s\n", string(out))
			}
			timer.Reset(m.timeInterval * time.Second)
		}
	}()
	return nil
}

func (m *memory) Alarm() error {
	return nil
}
