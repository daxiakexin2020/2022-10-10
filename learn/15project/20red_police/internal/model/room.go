package model

type status int

const (
	STATUS_WAITING status = iota + 1
	STATUS_PLAYING
	STATUS_DISSOLVE
	STATUS_OVER
)

type Room struct {
	Id           string
	Name         string
	MapName      string
	MapUserCount int
	Status       status
	CreateTime   string
	Players      map[string]*Player
	Owner        string
}
