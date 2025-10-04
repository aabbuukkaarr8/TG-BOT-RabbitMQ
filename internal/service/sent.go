package service

func (s *Service) Sent(id int64) error {
	err := s.repo.Sent(id)
	if err != nil {
		return err
	}
	return nil
}
