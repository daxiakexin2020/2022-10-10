package draw

import (
	"12lucky_draw/model/draw/weightedrand"
	"errors"
	"log"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type draw_level int

type time_key string

const (
	DRAW_LEVEL_1 draw_level = iota + 1
	DRAW_LEVEL_2
	DRAW_LEVEL_3
	DRAW_LEVEL_4
)

const (
	stage_01 time_key = "00-08"
	stage_02 time_key = "08-12"
	stage_03 time_key = "12-24"
)

var total_draw_poll = map[draw_level]int32{
	DRAW_LEVEL_1: 10,
	DRAW_LEVEL_2: 100,
	DRAW_LEVEL_3: 1000,
	DRAW_LEVEL_4: math.MaxInt32,
}

var time_draw_poll = map[draw_level]int32{
	DRAW_LEVEL_1: 0,
	DRAW_LEVEL_2: 0,
	DRAW_LEVEL_3: 0,
	DRAW_LEVEL_4: math.MaxInt32,
}

var currentTimeKey time_key

var drawTimeSet = map[time_key]map[draw_level]int32{
	stage_01: {
		DRAW_LEVEL_1: 2,
		DRAW_LEVEL_2: 10,
		DRAW_LEVEL_3: 100,
		DRAW_LEVEL_4: math.MaxInt32,
	},
	stage_02: {
		DRAW_LEVEL_1: 3,
		DRAW_LEVEL_2: 30,
		DRAW_LEVEL_3: 300,
		DRAW_LEVEL_4: math.MaxInt32,
	},
	stage_03: {
		DRAW_LEVEL_1: 5,
		DRAW_LEVEL_2: 60,
		DRAW_LEVEL_3: 600,
		DRAW_LEVEL_4: math.MaxInt32},
}

var draw_level_title = map[draw_level]string{
	DRAW_LEVEL_1: "一",
	DRAW_LEVEL_2: "二",
	DRAW_LEVEL_3: "三",
	DRAW_LEVEL_4: "谢",
}

var draw_count = map[draw_level]int64{
	DRAW_LEVEL_1: 0,
	DRAW_LEVEL_2: 0,
	DRAW_LEVEL_3: 0,
	DRAW_LEVEL_4: 0,
}

var (
	glock sync.Mutex
	start bool
)

type draw struct {
	wd *weightedrand.Weightedrand
}

func Pick(weight int32) (draw_level, error) {
	if !start {
		return 0, errors.New("draw not starting")
	}
	glock.Lock()
	defer glock.Unlock()
	poll := make(map[any]int32)
	for k, v := range time_draw_poll {
		if k != DRAW_LEVEL_4 {
			poll[k] = v * weight
		} else {
			poll[k] = v
		}
	}
	d := &draw{wd: weightedrand.NewWeightedrand(poll)}
	return d.pick(), nil
}

func (d *draw) pick() draw_level {
	result := d.wd.Pick()
	level := result.(draw_level)
	if level != DRAW_LEVEL_4 {
		d.sub(level)
	}
	d.count(level)
	return level
}

func (d *draw) sub(level draw_level) {
	c := time_draw_poll[level]
	if c >= 0 {
		atomic.AddInt32(&c, -1)
		time_draw_poll[level] = c
	}
}

func (d *draw) count(level draw_level) {
	c := draw_count[level]
	atomic.AddInt64(&c, 1)
	draw_count[level] = c
}

func CountResult() map[draw_level]int64 {
	return draw_count
}

func GetLevelTitle(level draw_level) string {
	return draw_level_title[level]
}

func Start() error {
	glock.Lock()
	defer glock.Unlock()
	if start {
		return errors.New("draw already starting")
	}

	h := time.Now().Hour()
	initTimeDrawPoll(h)

	start = true
	return nil
}

func ResetTimeDrawPoll(h int) {
	glock.Lock()
	defer glock.Unlock()
	log.Println("check draw........................................")
	t := checkTimeKey(h)
	if t == currentTimeKey {
		return
	}
	currentTimeKey = t
	time_draw_poll = drawTimeSet[t]
}

func initTimeDrawPoll(h int) error {
	t := checkTimeKey(h)
	currentTimeKey = t
	time_draw_poll = drawTimeSet[t]
	return nil
}

func checkTimeKey(h int) time_key {
	var t time_key
	if 0 <= h && h < 8 {
		t = stage_01
	} else if 8 <= h && h < 12 {
		t = stage_02
	} else {
		t = stage_03
	}
	return t
}

func ShowTimeDrawPoll() map[draw_level]int32 {
	return time_draw_poll
}

func Stop() error {
	glock.Lock()
	defer glock.Unlock()
	if !start {
		return errors.New("draw already inactive")
	}
	start = false
	return nil
}
