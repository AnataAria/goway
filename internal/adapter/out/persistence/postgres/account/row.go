package account

import (
	"time"
)

type AccountRow struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
