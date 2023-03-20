package data

import (
	"20red_police/common"
)

func FileUser() (Class, error) {
	pick, err := GclassTree().Pick(common.REGISTER_FILE_DATA_USER)
	if err != nil {
		return nil, err
	}
	return pick, nil
}

func MemoryUser() (Class, error) {
	pick, err := GclassTree().Pick(common.REGISTER_DATA_USER)
	if err != nil {
		return nil, err
	}
	return pick, err
}

func MemoryRoom() (Class, error) {
	pick, err := GclassTree().Pick(common.REGISTER_DATA_ROOM)
	if err != nil {
		return nil, err
	}
	return pick, nil
}
