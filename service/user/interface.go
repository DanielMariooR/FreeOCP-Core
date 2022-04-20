package user

import (
	"context"

	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user/user_repository"
)

type UserService interface {
	InjectUserRepository(user_repository.UserRepository) error
	GetUserData(ctx context.Context, email string) (*models.User, error)
	SignIn(ctx context.Context, input *models.SignInInput) (*models.SignInResponse, error)
	SignUp(ctx context.Context, input *models.SignUpInput) (*models.SignUpResponse, error)
}
