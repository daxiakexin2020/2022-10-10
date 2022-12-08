package server_govern

import "errors"

func (s *server) Register() error {
	rs, err := s.getRegisteServer()
	if err != nil {
		return err
	}
	if len(s.Addr) == 0 {
		return errors.New("Register服务，不允许服务地址为空")
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
