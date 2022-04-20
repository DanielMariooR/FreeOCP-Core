package user_repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
)

type userRepository struct{}

func NewRepository() UserRepository {
	return &userRepository{}
}

func (repo *userRepository) GetTableName() string {
	return "User"
}

func (repo *userRepository) queryGetUser() sq.SelectBuilder {
	builder := sq.Select(
		"id",
		"email",
		"username",
		"password",
		"isAdmin",
	).From(repo.GetTableName())

	return builder
}

func (repo *userRepository) queryInsertUser() sq.InsertBuilder {
	builder := sq.Insert(repo.GetTableName()).Columns(
		"id",
		"fullname",
		"email",
		"username",
		"password",
		"isAdmin",
	)

	return builder
}

func (repo *userRepository) GetUserByEmail(ctx context.Context, db *sqlx.DB, email string) (*db_models.User, error) {
	out := new(db_models.User)
	query, args, err := repo.queryGetUser().
		Where(sq.Eq{"email": email}).ToSql()
	if err != nil {
		return nil, err
	}

	err = db.GetContext(ctx, out, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return out, nil
}

func (repo *userRepository) GetUserById(ctx context.Context, db *sqlx.DB, userId string) (*db_models.User, error) {
	out := new(db_models.User)
	query, args, err := repo.queryGetUser().
		Where(sq.Eq{"id": userId}).ToSql()
	if err != nil {
		return nil, err
	}

	err = db.GetContext(ctx, out, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err

	}

	return out, nil
}

func (repo *userRepository) InsertNewUser(ctx context.Context, db *sqlx.DB, values *db_models.User) error {
	query, args, err := repo.queryInsertUser().Values(
		values.ID,
		values.Name,
		values.Email,
		values.Username,
		values.Password,
		values.Admin,
	).ToSql()

	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
