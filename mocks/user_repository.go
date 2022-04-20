// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import db "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
import mock "github.com/stretchr/testify/mock"
import sqlx "github.com/jmoiron/sqlx"

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetTableName provides a mock function with given fields:
func (_m *UserRepository) GetTableName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetUserByEmail provides a mock function with given fields: ctx, _a1, email
func (_m *UserRepository) GetUserByEmail(ctx context.Context, _a1 *sqlx.DB, email string) (*db.User, error) {
	ret := _m.Called(ctx, _a1, email)

	var r0 *db.User
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, string) *db.User); ok {
		r0 = rf(ctx, _a1, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.DB, string) error); ok {
		r1 = rf(ctx, _a1, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserById provides a mock function with given fields: ctx, _a1, userId
func (_m *UserRepository) GetUserById(ctx context.Context, _a1 *sqlx.DB, userId string) (*db.User, error) {
	ret := _m.Called(ctx, _a1, userId)

	var r0 *db.User
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, string) *db.User); ok {
		r0 = rf(ctx, _a1, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.DB, string) error); ok {
		r1 = rf(ctx, _a1, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertNewUser provides a mock function with given fields: ctx, _a1, values
func (_m *UserRepository) InsertNewUser(ctx context.Context, _a1 *sqlx.DB, values *db.User) error {
	ret := _m.Called(ctx, _a1, values)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, *db.User) error); ok {
		r0 = rf(ctx, _a1, values)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}