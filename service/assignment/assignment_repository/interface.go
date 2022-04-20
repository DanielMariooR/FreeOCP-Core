package assignment_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
)

type AssignmentRepository interface {
	GetTableName() string
	GetProblemTableName() string
	GetAssignmentById(ctx context.Context, db *sqlx.DB, id string) (*db_models.Assignment, error)
	InsertAssignment(ctx context.Context, db *sqlx.DB, value *db_models.AssignmentCreation) error
	GetAssignmentProblemsById(ctx context.Context, db *sqlx.DB, id string) ([]*db_models.ProblemTypeDetail, error)
	InsertAssignmentDesc(ctx context.Context, db *sqlx.DB, value *db_models.Assignment) error
	InsertAssignmentProblem(ctx context.Context, db *sqlx.DB, values *[]db_models.AssignmentProblem) error
}
