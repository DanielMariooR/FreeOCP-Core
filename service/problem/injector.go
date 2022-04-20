package problem

import (
	"errors"

	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/problem/problem_repository"
)

func (svc *problemService) InjectRepository(repo problem_repository.ProblemRepository) error {
	if repo != nil {
		svc.repository = repo
		return nil
	}
	return errors.New("problem repository not found")
}

