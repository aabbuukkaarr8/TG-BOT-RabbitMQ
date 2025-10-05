package service

import "context"

func (s *Service) Sent(ctx context.Context, id int64) error {
	err := s.repo.Sent(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
