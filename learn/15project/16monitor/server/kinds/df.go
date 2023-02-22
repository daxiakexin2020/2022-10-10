package kinds

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

type df struct {
	base       Base
	rows       [][]byte
	info       map[string]string
	isAlarming bool
	mu         sync.Mutex
}

const (
	DF                       = "磁盘"
	DEFAULT_DF_PERCENT       = 85
	DEFAULT_DF_TIME_INTERVAL = 5 * time.Second
	DEFAULT_DF_ALARM_COUNT   = 500
	SYMBOL                   = '*'
)

var DefaultDF = NewDF(DEFAULT_DF_PERCENT, DEFAULT_DF_TIME_INTERVAL)

func NewDF(percent int, timeInterval time.Duration) *df {
	if percent > 100 || percent < 0 {
		percent = DEFAULT_DF_PERCENT
	}
	return &df{base: Base{percent: percent, timeInterval: timeInterval}, info: map[string]string{}}
}

func (d *df) Name() string {
	return DF
}

func (d *df) Monitor(ctx context.Context) error {
	timer := time.NewTimer(d.base.timeInterval * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				cmd := exec.Command("df", "-h")
				out, err := cmd.CombinedOutput()
				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}
				fmt.Println(DF + "监控ok")

				//fmt.Printf("combined out:\n%s\n", string(out))
				row := []byte{}
				rows := [][]byte{}
				lastByte := byte(SYMBOL)
				for _, b := range out {
					if b == 10 {
						rows = append(rows, row)
						row = []byte{}
						continue
					}
					if b != 32 && b != 9 {
						lastByte = SYMBOL
						row = append(row, b)
					} else {
						if lastByte == SYMBOL {
							row = append(row, SYMBOL)
							lastByte = ' '
						}
					}
				}
				d.rows = rows
				d.Alarm()
			case <-ctx.Done():
				fmt.Println("上游协程退出了")
				timer.Stop()
				return
			}
			timer.Reset(d.base.timeInterval * time.Second)
		}
	}()
	return nil
}

func (d *df) Alarm() error {

	if len(d.rows) == 0 {
		return nil
	}
	var builder strings.Builder
	info := make(map[string]string, 0)
	for i, row := range d.rows {
		if i == 0 {
			continue
		}
		builder.Write(row)
		content := strings.Split(builder.String(), string(SYMBOL))
		if len(content) >= 8 {
			disk := content[0]
			used, err := strconv.Atoi(strings.Trim(content[7], "%"))
			if err == nil {
				if used >= d.base.percent && used <= 100 {
					info[disk] = fmt.Sprintf("磁盘使用率%v,请关注", strconv.Itoa(used)+"%")
				}
			}
		}
		builder.Reset()
	}
	d.info = info
	return nil
}

func (d *df) Html(info map[string]string) {
	var builder strings.Builder
	for disk, data := range info {
		builder.WriteString("<div style='color:red;font-size:20px'>" + disk + data + "</div>")
	}
}
