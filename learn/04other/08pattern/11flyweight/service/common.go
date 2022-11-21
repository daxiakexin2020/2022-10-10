package service

import "sync"

type colorManage struct {
	color string
}

var (
	cm   *colorManage
	once sync.Once
)

func makeColorManage() {
	once.Do(func() {
		if cm == nil {
			cm = &colorManage{color: "黄皮肤"}
		}
	})
}

func GetColorManage() *colorManage {
	return cm
}
