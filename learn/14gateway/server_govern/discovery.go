package server_govern

func (s *server) Discovery() ([]string, error) {
	rs, err := s.getRegisteServer()
	if err != nil {
		return nil, err
	}
	list, err := rs.Get(s.makeKey())
	if err != nil {
		return nil, err
	}
	return list, nil
}
