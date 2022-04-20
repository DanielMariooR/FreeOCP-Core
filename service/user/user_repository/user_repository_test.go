package user_repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user/user_repository"
)

var (
	id       = uuid.New().String()
	name     = "Daniel Mario Reynaldi"
	email    = "danielmarioreynaldi@gmail.com"
	password = "somepassword"
	username = "danielmr"
	admin    = false
)

func TestUserRepository_GetTableName(t *testing.T) {
	r := user_repository.NewRepository()

	expectedTableName := "User"
	assert.Equal(t, expectedTableName, r.GetTableName())
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}

	type mockSelect struct {
		userByEmail *db_models.User
		err         error
	}

	tests := []struct {
		name       string
		args       args
		mockSelect mockSelect
		want       *db_models.User
		wantErr    error
	}{
		{
			name: "Success to get user by email",
			args: args{
				context.TODO(),
				email,
			},
			mockSelect: mockSelect{
				userByEmail: &db_models.User{
					ID:       id,
					Email:    email,
					Username: username,
					Name:     name,
					Password: password,
					Admin:    admin,
				},
			},
			want: &db_models.User{
				ID:       id,
				Email:    email,
				Username: username,
				Password: password,
				Admin:    admin,
			},
			wantErr: nil,
		},
		{
			name: "failed to get user by email",
			args: args{
				context.TODO(),
				email,
			},
			mockSelect: mockSelect{
				err: sql.ErrNoRows,
			},
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "sqlmock")

			if tt.mockSelect.userByEmail != nil {
				data := tt.mockSelect.userByEmail
				rows := sqlmock.NewRows([]string{
					"id",
					"email",
					"username",
					"password",
					"isAdmin",
				})
				rows.AddRow(data.ID, data.Email, data.Username, data.Password, data.Admin)
				mock.ExpectQuery(`SELECT id, email, username, password, isAdmin`).WillReturnRows(rows)
				mock.ExpectCommit()

			}

			if tt.mockSelect.err != nil {
				mock.ExpectQuery(`SELECT id, email, username, password, isAdmin`).WillReturnError(tt.mockSelect.err)
			}

			r := user_repository.NewRepository()
			got, err := r.GetUserByEmail(tt.args.ctx, sqlxDB, tt.args.email)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}

}

func TestUserRepository_GetUserByID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
	}

	type mockSelect struct {
		userByEmail *db_models.User
		err         error
	}

	tests := []struct {
		name       string
		args       args
		mockSelect mockSelect
		want       *db_models.User
		wantErr    error
	}{
		{
			name: "Success to get user by email",
			args: args{
				context.TODO(),
				id,
			},
			mockSelect: mockSelect{
				userByEmail: &db_models.User{
					ID:       id,
					Email:    email,
					Username: username,
					Name:     name,
					Password: password,
					Admin:    admin,
				},
			},
			want: &db_models.User{
				ID:       id,
				Email:    email,
				Username: username,
				Password: password,
				Admin:    admin,
			},
			wantErr: nil,
		},
		{
			name: "failed to get user by email",
			args: args{
				context.TODO(),
				email,
			},
			mockSelect: mockSelect{
				err: sql.ErrNoRows,
			},
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "sqlmock")

			if tt.mockSelect.userByEmail != nil {
				data := tt.mockSelect.userByEmail
				rows := sqlmock.NewRows([]string{
					"id",
					"email",
					"username",
					"password",
					"isAdmin",
				})
				rows.AddRow(data.ID, data.Email, data.Username, data.Password, data.Admin)
				mock.ExpectQuery(`SELECT id, email, username, password, isAdmin`).WillReturnRows(rows)
				mock.ExpectCommit()

			}

			if tt.mockSelect.err != nil {
				mock.ExpectQuery(`SELECT id, email, username, password, isAdmin`).WillReturnError(tt.mockSelect.err)
			}

			r := user_repository.NewRepository()
			got, err := r.GetUserById(tt.args.ctx, sqlxDB, tt.args.userId)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}

}

func TestUserRepository_InsertNewUser(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *db_models.User
	}

	type mockExec struct {
		userData *db_models.User
		err      error
	}

	tests := []struct {
		name     string
		args     args
		mockExec mockExec
		wantErr  error
	}{
		{
			name: "success to insert new user",
			args: args{
				context.TODO(),
				&db_models.User{
					ID:       id,
					Name:     name,
					Username: username,
					Email:    email,
					Password: password,
					Admin:    admin,
				},
			},
			mockExec: mockExec{
				userData: &db_models.User{
					ID:       id,
					Name:     name,
					Username: username,
					Email:    email,
					Password: password,
					Admin:    admin,
				},
				err: nil,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "sqlmock")

			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO User (id,fullname,email,username,password,isAdmin) VALUES (?,?,?,?,?,?)`)).WillReturnResult(sqlmock.NewResult(1, 1))

			r := user_repository.NewRepository()
			err = r.InsertNewUser(tt.args.ctx, sqlxDB, tt.args.input)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}
