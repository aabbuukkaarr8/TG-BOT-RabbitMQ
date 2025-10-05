package repository

import "context"

func (r *Repository) Delete(ctx context.Context, id int64) error {
	query := `UPDATE notifications SET status = 'deleted' WHERE id = $1`
	_, err := r.store.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
