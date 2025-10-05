package service

import "context"

func (s *Service) Status(ctx context.Context, id int64) (string, error) {
	status, err := s.repo.Status(ctx, id)
	if err != nil {
		return "", err
	}
	return status, nil
}
