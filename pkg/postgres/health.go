package postgres

func (db *DB) Close() error {
	return db.DB.Close()
}
