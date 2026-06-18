package login

import (
	"context"
	"database/sql"
)

type repositoryLogin struct {
	db *sql.DB
}

// GetByEmail implements [repoContract].
func (r repositoryLogin) GetByEmail(ctx context.Context, email string) (auth Auth, err error) {
	query := `
		SELECT
			public_id
			, email
			, password
			, is_active
			, role
		FROM auth
		WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email)
	err = row.Scan(
		&auth.PublicID,
		&auth.Email,
		&auth.Password,
		&auth.IsActive,
		&auth.Role,
	)

	if err != nil {
		return
	}

	return
}

func NewRepositoryLogin(db *sql.DB) repositoryLogin {
	return repositoryLogin{db: db}
}
