package service

func (s *Service) Delete(id int64) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
