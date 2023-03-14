package network

type MiddleFunc func(request *Request) error

type middle struct {
	f []MiddleFunc
}

func newMiddle() *middle {
	return &middle{f: make([]MiddleFunc, 0)}
}

func (m *middle) register(mf ...MiddleFunc) {
	m.f = append(m.f, mf...)
}

func (m *middle) run(request *Request) error {
	for _, mf := range m.f {
		err := mf(request)
		if err != nil {
			return err
		}
	}
	return nil
}
