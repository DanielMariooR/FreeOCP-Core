package user_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
)

type UserRepository interface {
	GetTableName() string
	GetUserByEmail(ctx context.Context, db *sqlx.DB, email string) (*db_models.User, error)
	GetUserById(ctx context.Context, db *sqlx.DB, userId string) (*db_models.User, error)
	InsertNewUser(ctx context.Context, db *sqlx.DB, values *db_models.User) error
}
