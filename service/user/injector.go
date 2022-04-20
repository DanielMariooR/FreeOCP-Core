package user

import (
	"errors"

	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user/user_repository"
)

func (svc *userService) InjectUserRepository(repo user_repository.UserRepository) error {
	if repo != nil {
		svc.userRepository = repo
		return nil
	}
	return errors.New("user repository not found")
}
