package problem_repository

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
)

type problemRepository struct{}

var STATUS [3]string = [3]string{"requested", "rejected", "accepted"}

func NewRepository() ProblemRepository {
	return &problemRepository{}
}

func (repo *problemRepository) GetTableName() string {
	return "Candidate_Problem"
}

func (repo *problemRepository) GetDetailTableName() string {
	return "Detail_Problem"
}

func (repo *problemRepository) querySelectProblemCandidate() sq.SelectBuilder {
	builder := sq.Select(
		"id",
		"creator",
		"title",
		"type",
		"topic",
		"difficulty",
		"status",
	).From(repo.GetTableName())

	return builder
}

func (repo *problemRepository) queryUpdateProblemCandidate() sq.UpdateBuilder {
	builder := sq.Update(repo.GetTableName())

	return builder
}

func (repo *problemRepository) GetDetailById(ctx context.Context, db *sqlx.DB, id string) (*db_models.ProblemDetail, error) {
	out := new(db_models.ProblemDetail)
	query, args, err := sq.Select(
		"id",
		"detail",
	).From(repo.GetDetailTableName()).Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	err = db.GetContext(ctx, out, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return out, nil
}

func (repo *problemRepository) GetCandidateById(ctx context.Context, db *sqlx.DB, id string) (*db_models.ProblemCandidate, error) {
	out := new(db_models.ProblemCandidate)
	query, args, err := sq.Select(
		"id", "creator", "title", "type", "topic", "difficulty", "status",
	).From(repo.GetTableName()).Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	err = db.GetContext(ctx, out, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	detail, err := repo.GetDetailById(ctx, db, id)
	if err != nil {
		return nil, err
	}

	out.Detail = detail.Detail

	return out, nil
}

func (repo *problemRepository) InsertNewProblemData(ctx context.Context, db *sqlx.DB, values *db_models.ProblemCandidate) error {
	query, args, err := sq.Insert(repo.GetTableName()).
		Columns(
			"id", "creator", "title", "type", "topic", "difficulty", "status",
		).
		Values(
			values.ID,
			values.Creator,
			values.Title,
			values.Type,
			values.Topic,
			values.Difficulty,
			values.Status,
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

func (repo *problemRepository) InsertNewProblemDetail(ctx context.Context, db *sqlx.DB, values *db_models.ProblemDetail) error {
	query, args, err := sq.Insert(repo.GetDetailTableName()).
		Columns(
			"id", "detail",
		).
		Values(
			values.ID, values.Detail,
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

func (repo *problemRepository) InsertNewProblem(ctx context.Context, db *sqlx.DB, values *db_models.ProblemCandidate) error {
	err := repo.InsertNewProblemData(ctx, db, values)
	if err != nil {
		return err
	}

	detail := &db_models.ProblemDetail{
		ID:     values.ID,
		Detail: values.Detail,
	}

	err = repo.InsertNewProblemDetail(ctx, db, detail)

	return err
}

func (repo *problemRepository) GetProblemsByUserId(ctx context.Context, db *sqlx.DB, userId string) ([]*db_models.ProblemCandidate, error) {
	var out []*db_models.ProblemCandidate
	query, args, err := repo.querySelectProblemCandidate().
		Where(sq.Eq{"creator": userId}).ToSql()

	if err != nil {
		return nil, err
	}

	err = db.SelectContext(ctx, &out, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return out, nil
}

func (repo *problemRepository) GetCandidateProblemList(ctx context.Context, db *sqlx.DB, filter models.ProblemFilter) ([]*db_models.ProblemCandidate, error) {
	var problems []*db_models.ProblemCandidate
	queryString := repo.querySelectProblemCandidate().Where(sq.NotEq{"status": "accepted"})

	if len(filter.Category) > 0 {
		queryString = queryString.Where(sq.Eq{"topic": filter.Category})
	}

	if len(filter.Difficulty) > 0 {
		queryString = queryString.Where(sq.Eq{"difficulty": filter.Difficulty})
	}

	query, args, err := queryString.ToSql()
	if err != nil {
		return problems, err
	}

	err = db.SelectContext(ctx, &problems, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return problems, nil
		}
		return nil, err
	}

	return problems, nil
}

func (repo *problemRepository) UpdateProblemStatus(ctx context.Context, db *sqlx.DB, input *models.ProblemStatusUpdate) error {
	var (
		found bool = false
	)

	for _, stat := range STATUS {
		if input.Status == stat {
			found = true
		}
	}

	if !found {
		return errors.New("Invalid Status")
	}

	query, args, err := repo.queryUpdateProblemCandidate().
		Set("status", input.Status).
		Where(sq.Eq{"id": input.Id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *problemRepository) GetProblemList(ctx context.Context, db *sqlx.DB, filter models.ProblemFilter) ([]*db_models.ProblemCandidate, error) {
	var problems []*db_models.ProblemCandidate
	queryString := repo.querySelectProblemCandidate().Where(sq.Eq{"status": "accepted"})

	if len(filter.Category) > 0 {
		queryString = queryString.Where(sq.Eq{"topic": filter.Category})
	}

	if len(filter.Difficulty) > 0 {
		queryString = queryString.Where(sq.Eq{"difficulty": filter.Difficulty})
	}

	query, args, err := queryString.ToSql()
	if err != nil {
		return problems, err
	}

	err = db.SelectContext(ctx, &problems, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return problems, nil
		}
		return nil, err
	}

	return problems, nil
}
