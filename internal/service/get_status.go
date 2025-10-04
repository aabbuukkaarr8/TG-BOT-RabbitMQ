package service

func (s *Service) Status(id int64) (string, error) {
	status, err := s.repo.Status(id)
	if err != nil {
		return "", err
	}
	return status, nil
}
