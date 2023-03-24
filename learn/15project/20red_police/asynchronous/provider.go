package asynchronous

import (
	"log"
)

func GoAsynchronous(task ...Tasker) error {
	m := Manager()
	if err := m.Register(task...); err != nil {
		return err
	}
	go m.Run()
	log.Println("Asynchronous is runing...........................")
	return nil
}

func STOP() error {
	return Manager().Stop()
}
