package user

import (
	"database/sql"

	domainUser "github.com/AnataAria/goway/internal/domain/user"
	persistUser "github.com/AnataAria/goway/internal/port/out/persistence/user"
	"github.com/jmoiron/sqlx"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) persistUser.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(user *domainUser.User) error {
	row := ToRow(user)
	_, err := r.db.NamedExec(`
		INSERT INTO users (id, email, password, name, created_at)
		VALUES (:id, :email, :password, :name, :created_at)
		ON CONFLICT (id) DO UPDATE SET
			email = :email,
			password = :password,
			name = :name,
			created_at = :created_at
	`, row)
	return err
}

func (r *PostgresUserRepository) FindByID(id string) (*domainUser.User, error) {
	var row UserRow
	err := r.db.Get(&row, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return ToDomain(row)
}

func (r *PostgresUserRepository) FindByEmail(email string) (*domainUser.User, error) {
	var row UserRow
	err := r.db.Get(&row, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return ToDomain(row)
}

func (r *PostgresUserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Get(&count, "SELECT count(*) FROM users WHERE email = $1", email)
	return count > 0, err
}
