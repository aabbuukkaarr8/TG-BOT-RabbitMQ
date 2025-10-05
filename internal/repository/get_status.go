package repository

import "context"

func (r *Repository) Status(ctx context.Context, id int64) (string, error) {
	query := `SELECT status FROM notifications WHERE id = $1`
	var status string
	err := r.store.DB.QueryRowContext(ctx, query, id).Scan(&status)
	return status, err
}
