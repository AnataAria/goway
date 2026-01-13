package account

import (
	"database/sql"

	domainAccount "github.com/AnataAria/goway/internal/domain/account"
	persistAccount "github.com/AnataAria/goway/internal/port/out/persistence/account"
	"github.com/jmoiron/sqlx"
)

type PostgresAccountRepository struct {
	db *sqlx.DB
}

func NewPostgresAccountRepository(db *sqlx.DB) persistAccount.AccountRepository {
	return &PostgresAccountRepository{db: db}
}

func (r *PostgresAccountRepository) Save(account *domainAccount.Account) error {
	row := ToRow(account)
	_, err := r.db.NamedExec(`
		INSERT INTO accounts (id, email, password, name, created_at)
		VALUES (:id, :email, :password, :name, :created_at)
		ON CONFLICT (id) DO UPDATE SET
			email = :email,
			password = :password,
			name = :name,
			created_at = :created_at
	`, row)
	return err
}

func (r *PostgresAccountRepository) FindByID(id string) (*domainAccount.Account, error) {
	var row AccountRow
	err := r.db.Get(&row, "SELECT * FROM accounts WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return ToDomain(row)
}

func (r *PostgresAccountRepository) FindByEmail(email string) (*domainAccount.Account, error) {
	var row AccountRow
	err := r.db.Get(&row, "SELECT * FROM accounts WHERE email = $1", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return ToDomain(row)
}

func (r *PostgresAccountRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Get(&count, "SELECT count(*) FROM accounts WHERE email = $1", email)
	return count > 0, err
}
