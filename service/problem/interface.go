package problem

import (
	"context"

	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/problem/problem_repository"
)

type ProblemService interface {
	InjectRepository(problem_repository.ProblemRepository) error
	GetProblemCandidate(ctx context.Context, id string) (*models.ProblemCandidate, error)
	CreateNewProblem(ctx context.Context, problem *models.ProblemCreationInput) (*models.ProblemCreationResponse, error)
	GetProblemStatus(ctx context.Context, id string) (*models.ProblemStatusList, error)
	GetProblemDetail(ctx context.Context, id string) (*models.ProblemDetail, error)
	GetProblemCandidateList(ctx context.Context, filter models.ProblemFilter) (*models.ProblemCandidateList, error)
	AcceptProblem(ctx context.Context, id string) (*models.ProblemCreationResponse, error)
	RejectProblem(ctx context.Context, id string) (*models.ProblemCreationResponse, error)
	GetProblemList(ctx context.Context, filter models.ProblemFilter) (*models.ProblemCandidateList, error)
}
