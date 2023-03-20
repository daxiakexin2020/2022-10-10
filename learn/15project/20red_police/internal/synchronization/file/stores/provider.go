package stores

import "20red_police/internal/synchronization/file"

func GoSynchronization() error {
	user, err := NewUser()
	if err != nil {
		return err
	}
	return file.Register(user)
}
