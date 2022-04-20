package assignment

import (
	"context"

	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/assignment/assignment_repository"
)

type AssignmentService interface {
	InjectAssignmentRepository(assignment_repository.AssignmentRepository) error
	GetAssignment(ctx context.Context, id string) (*models.AssignmentResponse, error)
	CreateAssignment(ctx context.Context, input *models.AssignmentCreation) (*models.AssignmentCreationResponse, error)
	CalculateScore(ctx context.Context, answers *models.AssignmentSubmission) (int, error)
	GetScore(ctx context.Context, userId string, answers *models.AssignmentSubmission) (*models.AssignmentScore, error)
}
