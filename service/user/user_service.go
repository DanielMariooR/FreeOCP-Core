package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/config"
	er "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/error"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user/user_repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	db             *sqlx.DB
	userRepository user_repository.UserRepository
}

func NewService(db *sqlx.DB) UserService {
	return &userService{
		db: db,
	}
}

func (svc *userService) GetUserData(ctx context.Context, email string) (*models.User, error) {
	user, err := svc.userRepository.GetUserByEmail(ctx, svc.db, email)
	if err != nil {
		return nil, err
	}

	resp := models.User{
		Username: user.Username,
		Email:    user.Email,
		ID:       user.ID,
	}

	return &resp, nil
}

func (svc *userService) SignIn(ctx context.Context, input *models.SignInInput) (*models.SignInResponse, error) {
	user, err := svc.userRepository.GetUserByEmail(ctx, svc.db, input.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, er.NewError(fmt.Errorf("%s", "User belum terdaftar"), http.StatusUnauthorized, nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, er.NewError(fmt.Errorf("%s", "Email atau password salah!"), http.StatusUnauthorized, nil)
	}

	secret := config.GetConfig().JWTSecret
	token, expiration, err := generateToken(user, time.Now().Add(1*time.Hour), []byte(secret))
	if err != nil {
		return nil, err
	}

	expAt := expiration.Format("2006-01-02 15:04:05")
	resp := models.SignInResponse{
		Email:     input.Email,
		Token:     token,
		ExpiresAt: expAt,
	}

	return &resp, nil
}

func generateToken(user *db_models.User, expiration time.Time, secret []byte) (string, time.Time, error) {
	claims := &models.Claims{
		ID:       user.ID,
		Admin:    user.Admin,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expiration, nil
}

func (svc *userService) SignUp(ctx context.Context, input *models.SignUpInput) (*models.SignUpResponse, error) {
	user, err := svc.userRepository.GetUserByEmail(ctx, svc.db, input.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, er.NewError(fmt.Errorf("%s", "User sudah terdaftar"), http.StatusUnauthorized, nil)
	}

	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	values := &db_models.User{
		ID:       uuid.New().String(),
		Name:     input.Name,
		Email:    input.Email,
		Username: input.Username,
		Password: string(bcryptPassword),
		Admin:    false,
	}

	err = svc.userRepository.InsertNewUser(ctx, svc.db, values)
	if err != nil {
		return nil, err
	}

	out := &models.SignUpResponse{
		Status:  "Success",
		Message: "User Created Succesfully",
	}

	return out, nil
}
