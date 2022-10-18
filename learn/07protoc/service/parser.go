package service

type Parser interface {
	Decode(filepath string) error
}
