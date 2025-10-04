package repository

func (r *Repository) Status(id int64) (string, error) {
	query := `SELECT status FROM notifications WHERE id = $1`
	var status string
	err := r.store.GetConn().QueryRow(query, id).Scan(&status)
	return status, err
}
