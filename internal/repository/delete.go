package repository

func (r *Repository) Delete(id int64) error {
	query := `UPDATE notifications SET status = 'deleted' WHERE id = $1`
	_, err := r.store.GetConn().Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
