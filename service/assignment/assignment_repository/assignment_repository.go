package assignment_repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	er "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/error"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
)

type assignmentRepository struct{}

func NewRepository() AssignmentRepository {
	return &assignmentRepository{}
}

func (repo *assignmentRepository) GetTableName() string {
	return "assignment"
}

func (repo *assignmentRepository) GetProblemTableName() string {
	return "assignment_problem"
}

func (repo *assignmentRepository) querySelectAssignment() sq.SelectBuilder {
	builder := sq.Select(
		"id",
		"creator",
		"title",
		"duration",
		"topic",
		"difficulty",
	).From(repo.GetTableName())

	return builder
}

func (repo *assignmentRepository) queryInsertAssignment() sq.InsertBuilder {
	builder := sq.Insert(repo.GetTableName()).Columns(
		"id",
		"creator",
		"title",
		"duration",
		"topic",
		"difficulty",
	)

	return builder
}

func (repo *assignmentRepository) queryInsertAssignmentProblem() sq.InsertBuilder {
	builder := sq.Insert(repo.GetProblemTableName()).Columns(
		"assignment_id",
		"problem_id",
	)

	return builder
}

func (repo *assignmentRepository) GetAssignmentById(ctx context.Context, db *sqlx.DB, id string) (*db_models.Assignment, error) {
	data := new(db_models.Assignment)
	queryString := repo.querySelectAssignment().Where(sq.Eq{"id": id})

	query, args, err := queryString.ToSql()
	if err != nil {
		return data, err
	}

	err = db.GetContext(ctx, data, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, er.NewError(fmt.Errorf("%s", "Assignment Not Found!"), http.StatusBadRequest, nil)
		}
		return nil, err
	}

	return data, nil
}

func (repo *assignmentRepository) InsertAssignmentDesc(ctx context.Context, db *sqlx.DB, value *db_models.Assignment) error {
	query, args, err := repo.queryInsertAssignment().Values(
		value.ID,
		value.Creator,
		value.Title,
		value.Duration,
		value.Topic,
		value.Difficulty,
	).ToSql()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *assignmentRepository) InsertAssignmentProblem(ctx context.Context, db *sqlx.DB, values *[]db_models.AssignmentProblem) error {
	if len(*values) <= 0 {
		return nil
	}

	queryBuilder := repo.queryInsertAssignmentProblem()

	for _, problem := range *values {
		queryBuilder = queryBuilder.Values(
			problem.AssignmentID,
			problem.ProblemID,
		)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *assignmentRepository) InsertAssignment(ctx context.Context, db *sqlx.DB, value *db_models.AssignmentCreation) error {
	err := repo.InsertAssignmentDesc(ctx, db, &value.Desc)
	if err != nil {
		return err
	}

	err = repo.InsertAssignmentProblem(ctx, db, &value.Problems)
	if err != nil {
		return err
	}

	return nil
}

func (repo *assignmentRepository) GetAssignmentProblemsById(ctx context.Context, db *sqlx.DB, id string) ([]*db_models.ProblemTypeDetail, error) {
	var problems []*db_models.ProblemTypeDetail

	query, args, err := sq.Select(
		"p.id as id",
		"p.detail as detail",
		"cp.type as type",
	).From("Detail_Problem p").InnerJoin("assignment_problems ap on p.id = ap.problem_id").InnerJoin("Candidate_Problem cp on cp.id = p.id").Where(sq.Eq{"ap.assignment_id": id}).ToSql()

	if err != nil {
		return problems, err
	}

	err = db.SelectContext(ctx, &problems, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return problems, nil
		}

		return problems, err
	}

	return problems, nil
}
