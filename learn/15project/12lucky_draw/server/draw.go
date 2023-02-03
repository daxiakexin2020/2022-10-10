package server

import (
	mdraw "12lucky_draw/model/draw"
	"github.com/robfig/cron/v3"
	"time"
)

type drawService struct{}

func NewDrawService() *drawService {
	return &drawService{}
}

func (ds *drawService) Draw(weight int32) (int, error) {
	res, err := mdraw.Pick(weight)
	return int(res), err
}

func (ds *drawService) resetTimeDrawPoll(h int) {
	mdraw.ResetTimeDrawPoll(h)
}

func (ds *drawService) SelectTimeDrawPoll() {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		ds.resetTimeDrawPoll(time.Now().Hour())
	})
	c.Start()
}

func (ds *drawService) Start() error {
	return mdraw.Start()
}

func (ds *drawService) ShowTimeDrawPoll() map[int]int32 {
	data := mdraw.ShowTimeDrawPoll()
	format := make(map[int]int32, len(data))
	for k, v := range data {
		format[int(k)] = v
	}
	return format
}

func (ds *drawService) CountResult() map[int]int64 {
	data := mdraw.CountResult()
	format := make(map[int]int64, len(data))
	for k, v := range data {
		format[int(k)] = v
	}
	return format
}
