package user_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	er "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/error"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/mocks"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user"
)

var (
	id       = uuid.New().String()
	name     = "Daniel Mario Reynaldi"
	email    = "danielmarioreynaldi@gmail.com"
	password = "somepassword"
	username = "danielmr"
	admin    = false
)

func TestUserService_GetUserData(t *testing.T) {
	type mockUserRepo struct {
		res *db_models.User
		err error
	}

	type args struct {
		ctx   context.Context
		email string
	}

	tests := []struct {
		name         string
		args         args
		mockUserRepo mockUserRepo
		want         *models.User
		wantErr      error
	}{
		{
			name: "Success to get User Data",
			args: args{
				context.TODO(),
				email,
			},
			mockUserRepo: mockUserRepo{
				res: &db_models.User{
					ID:       id,
					Name:     name,
					Email:    email,
					Username: username,
					Password: password,
					Admin:    admin,
				},
				err: nil,
			},
			want: &models.User{
				ID:       id,
				Username: username,
				Email:    email,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			userRepoMock := new(mocks.UserRepository)
			svc := user.NewService(sqlxDB)
			svc.InjectUserRepository(userRepoMock)

			userRepoMock.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockUserRepo.res, tt.mockUserRepo.err)

			got, err := svc.GetUserData(tt.args.ctx, tt.args.email)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}

}

func TestUserService_SignUp(t *testing.T) {

	type mockGetUser struct {
		res *db_models.User
		err error
	}

	type mockInsertUser struct {
		err error
	}

	type args struct {
		ctx   context.Context
		input *models.SignUpInput
	}

	tests := []struct {
		name           string
		args           args
		mockGetUser    mockGetUser
		mockInsertUser mockInsertUser
		want           *models.SignUpResponse
		wantErr        error
	}{
		{
			name: "Success to sign up",
			args: args{
				context.TODO(),
				&models.SignUpInput{
					Email:    email,
					Name:     name,
					Username: username,
					Password: password,
				},
			},
			mockGetUser: mockGetUser{
				res: nil,
				err: nil,
			},
			mockInsertUser: mockInsertUser{
				err: nil,
			},
			want: &models.SignUpResponse{
				Status:  "Success",
				Message: "User Created Succesfully",
			},
			wantErr: nil,
		},
		{
			name: "Fail to sign up (User Exists)",
			args: args{
				context.TODO(),
				&models.SignUpInput{
					Email:    email,
					Name:     name,
					Username: username,
					Password: password,
				},
			},
			mockGetUser: mockGetUser{
				res: &db_models.User{
					ID:       id,
					Email:    email,
					Name:     name,
					Username: username,
					Password: password,
					Admin:    admin,
				},
				err: nil,
			},
			mockInsertUser: mockInsertUser{
				err: nil,
			},
			want:    nil,
			wantErr: er.NewError(fmt.Errorf("%s", "User sudah terdaftar"), http.StatusUnauthorized, nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			userRepoMock := new(mocks.UserRepository)
			svc := user.NewService(sqlxDB)
			svc.InjectUserRepository(userRepoMock)

			userRepoMock.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetUser.res, tt.mockGetUser.err)
			userRepoMock.On("InsertNewUser", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockInsertUser.err)

			got, err := svc.SignUp(tt.args.ctx, tt.args.input)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}

}

func TestUserService_SignIn(t *testing.T) {
	type mockGetUser struct {
		res *db_models.User
		err error
	}

	type args struct {
		ctx   context.Context
		input *models.SignInInput
	}

	tests := []struct {
		name        string
		args        args
		mockGetUser mockGetUser
		want        *models.SignInResponse
		wantErr     error
	}{
		{
			name: "Success to sign up",
			args: args{
				context.TODO(),
				&models.SignInInput{
					Email:    email,
					Password: password,
				},
			},
			mockGetUser: mockGetUser{
				res: nil,
				err: nil,
			},
			want:    nil,
			wantErr: er.NewError(fmt.Errorf("%s", "User belum terdaftar"), http.StatusUnauthorized, nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			userRepoMock := new(mocks.UserRepository)
			svc := user.NewService(sqlxDB)
			svc.InjectUserRepository(userRepoMock)

			userRepoMock.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetUser.res, tt.mockGetUser.err)

			got, err := svc.SignIn(tt.args.ctx, tt.args.input)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}
}
