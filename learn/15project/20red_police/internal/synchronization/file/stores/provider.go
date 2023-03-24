package stores

import (
	"20red_police/internal/synchronization/file"
	"log"
)

func SynchronizationRegister() error {
	user, err := NewUser()
	if err != nil {
		return err
	}
	pmap, err := NewPmap()
	if err != nil {
		return err
	}
	return file.Register(user, pmap)
}

func GoSynchronizationRunBuilder() error {
	if err := SynchronizationRegister(); err != nil {
		return err
	}
	if err := UserFileSyncMemory(); err != nil {
		return err
	}
	if err := PmapFileSyncMemory(); err != nil {
		return err
	}
	log.Println("GoSynchronizationRunBuilder ok..........................................................")
	return nil
}

func GoSynchronizationStopBuilder() error {
	if err := UserMemoryBatchSyncFile(); err != nil {
		return err
	}
	if err := CloseUserSyncFile(); err != nil {
		return err
	}
	if err := PMapMemoryBatchSyncFile(); err != nil {
		return err
	}
	if err := ClosePmapSyncFile(); err != nil {
		return err
	}
	return nil
}
