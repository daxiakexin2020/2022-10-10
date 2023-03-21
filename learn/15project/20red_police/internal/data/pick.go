package data

import (
	"20red_police/common"
)

func FileUser() (Class, error) {
	pick, err := GclassTree().Pick(common.REGISTER_FILE_USER)
	if err != nil {
		return nil, err
	}
	return pick, nil
}

func MemoryUser() (Class, error) {
	pick, err := GclassTree().Pick(common.REGISTER_MEMORY_USER)
	if err != nil {
		return nil, err
	}
	return pick, err
}

func MemoryRoom() (Class, error) {
	pick, err := GclassTree().Pick(common.REGISTER_MEMORY_ROOM)
	if err != nil {
		return nil, err
	}
	return pick, nil
}

func MemoryPmap() (Class, error) {
	pick, err := GclassTree().Pick(common.REGISTER_MEMORY_PMAP)
	if err != nil {
		return nil, err
	}
	return pick, nil
}

func FilePMap() (Class, error) {
	pick, err := GclassTree().Pick(common.REGISTER_FILE_PMAP)
	if err != nil {
		return nil, err
	}
	return pick, nil
}
