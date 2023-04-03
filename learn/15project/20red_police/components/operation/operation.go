package operation

type operation struct {
	list []chan struct{}
}

func Oeration() *operation {
	return defaultOeration
}

var defaultOeration = newoOeration()

func newoOeration() *operation {
	return &operation{list: make([]chan struct{}, 0)}
}

func (o *operation) Register(c ...chan struct{}) {
	o.list = append(o.list, c...)
}

func (o *operation) Each() error {
	for _, c := range o.list {
		if len(c) == 0 && cap(c) == 1 {
			c <- struct{}{}
		}
	}
	return nil
}
