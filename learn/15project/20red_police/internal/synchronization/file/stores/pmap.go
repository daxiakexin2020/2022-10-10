package stores

import (
	"20red_police/common"
	"20red_police/internal/data"
	"20red_police/internal/data/memory"
	"20red_police/internal/model"
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

type Pmap struct {
	store
	mu sync.Mutex
}

const pmap_file_path = "pmap.txt"

func NewPmap() (*Pmap, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := pwd + "/internal/synchronization/file/txt/" + pmap_file_path
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	return &Pmap{store: store{file: file, reader: bufio.NewReader(file), writer: bufio.NewWriter(file)}}, nil
}

func (p *Pmap) Read() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	pick, err := data.MemoryPmap()
	if err != nil {
		return err
	}
	mpmap, ok := pick.(*memory.PMap)
	if !ok {
		return errors.New("memory class err:")
	}

	var pmaps []*model.PMap
	err = p.store.read(func(buf []byte) {
		var pmap model.PMap
		if err = json.Unmarshal(buf, &pmap); err != nil {
			log.Println("read json unmarshal err:", err)
		} else {
			pmaps = append(pmaps, &pmap)
		}
	})
	if err != nil {
		return err
	}
	for _, pmap := range pmaps {
		if _, err = mpmap.Create(pmap); err != nil {
			log.Println("file into memory register user err:", err)
		}
	}
	return nil
}

func (p *Pmap) Write(wdata interface{}) error {
	if wdata == nil {
		return p.bacthWrite()
	}
	return p.write(wdata)
}

func (p *Pmap) write(data interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	line, err := json.Marshal(data)
	if err != nil {
		log.Println("file marshal err:", err)
		return err
	}
	line = append(line, '\n')
	p.store.write(line)
	return p.store.flush()
}

func (p *Pmap) bacthWrite() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	pick, err := data.MemoryPmap()
	if err != nil {
		return err
	}
	mpmp, ok := pick.(*memory.PMap)
	if !ok {
		return errors.New("memory class err:")
	}

	var lines []byte
	list := mpmp.List()
	for _, pmap := range list {
		line, err := json.Marshal(pmap)
		if err != nil {
			log.Println("file marshal err:", err)
			continue
		}
		line = append(line, '\n')
		lines = append(lines, line...)
	}
	p.store.write(lines)
	return p.store.flush()
}

func (p *Pmap) Close() error {
	return p.store.close()
}

func (p *Pmap) Name() string {
	return common.REGISTER_FILE_PMAP
}

func PMapMemoryBatchSyncFile() error {
	pick, err := data.FilePMap()
	if err != nil {
		return err
	}
	pmap, ok := pick.(*Pmap)
	if !ok {
		return errors.New("file pmap class is err")
	}
	return pmap.Write(nil)
}

func ClosePmapSyncFile() error {
	pick, err := data.FilePMap()
	if err != nil {
		return err
	}
	pmap, ok := pick.(*Pmap)
	if !ok {
		return errors.New("file pmap class is err")
	}
	return pmap.store.close()
}

func PmapMemorySyncFile(model interface{}) error {
	pick, err := data.FilePMap()
	if err != nil {
		return err
	}
	pmap, ok := pick.(*Pmap)
	if !ok {
		return errors.New("file user class is err")
	}
	return pmap.Write(model)
}

func PmapFileSyncMemory() error {
	pick, err := data.FilePMap()
	if err != nil {
		return err
	}
	pmap, ok := pick.(*Pmap)
	if !ok {
		return errors.New("file user class is err")
	}

	return pmap.Read()
}
