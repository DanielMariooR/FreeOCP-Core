package assignment

import (
	"errors"

	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/assignment/assignment_repository"
)

func (svc *assignmentService) InjectAssignmentRepository(repo assignment_repository.AssignmentRepository) error {
	if repo != nil {
		svc.repository = repo
		return nil
	}
	return errors.New("assignment repository not found")
}
