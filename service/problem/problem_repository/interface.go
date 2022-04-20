package problem_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
)

type ProblemRepository interface {
	GetTableName() string
	GetDetailTableName() string
	GetCandidateById(ctx context.Context, db *sqlx.DB, id string) (*db_models.ProblemCandidate, error)
	InsertNewProblem(ctx context.Context, db *sqlx.DB, values *db_models.ProblemCandidate) error
	GetDetailById(ctx context.Context, db *sqlx.DB, id string) (*db_models.ProblemDetail, error)
	InsertNewProblemDetail(ctx context.Context, db *sqlx.DB, values *db_models.ProblemDetail) error
	GetProblemsByUserId(ctx context.Context, db *sqlx.DB, userId string) ([]*db_models.ProblemCandidate, error)
	GetCandidateProblemList(ctx context.Context, db *sqlx.DB, filter models.ProblemFilter) ([]*db_models.ProblemCandidate, error)
	UpdateProblemStatus(ctx context.Context, db *sqlx.DB, input *models.ProblemStatusUpdate) error
	GetProblemList(ctx context.Context, db *sqlx.DB, filter models.ProblemFilter) ([]*db_models.ProblemCandidate, error)
}
