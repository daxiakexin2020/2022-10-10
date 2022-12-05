package server_govern

func (s *server) Register() error {
	rs, err := s.getRegisteServer()
	if err != nil {
		return err
	}
	return rs.Put(s.makeKey(), s.makeValue())
}

func (s *server) UnRegister() error {
	rs, err := s.getRegisteServer()
	if err != nil {
		return err
	}
	return rs.Delete(s.makeKey())
}
