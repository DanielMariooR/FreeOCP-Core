package problem

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/problem/problem_repository"
)

type problemService struct {
	db         *sqlx.DB
	repository problem_repository.ProblemRepository
}

func NewService(db *sqlx.DB) ProblemService {
	return &problemService{
		db: db,
	}
}

func (svc *problemService) GetProblemCandidate(ctx context.Context, id string) (*models.ProblemCandidate, error) {
	problem, err := svc.repository.GetCandidateById(ctx, svc.db, id)
	if err != nil {
		return nil, err
	}

	resp := models.ProblemCandidate{
		ID:         problem.ID,
		Creator:    problem.Creator,
		Title:      problem.Title,
		Type:       problem.Type,
		Topic:      problem.Topic,
		Difficulty: problem.Difficulty,
		Status:     problem.Status,
		Detail:     problem.Detail,
	}

	return &resp, nil
}

func (svc *problemService) CreateNewProblem(ctx context.Context, problem *models.ProblemCreationInput) (*models.ProblemCreationResponse, error) {
	newId := uuid.New().String()
	problemData := &db_models.ProblemCandidate{
		ID:         newId,
		Creator:    problem.Creator,
		Title:      problem.Title,
		Type:       problem.Type,
		Topic:      problem.Topic,
		Difficulty: problem.Difficulty,
		Status:     "requested",
		Detail:     problem.Detail,
	}

	err := svc.repository.InsertNewProblem(ctx, svc.db, problemData)
	if err != nil {
		return nil, err
	}

	out := &models.ProblemCreationResponse{
		Status:  "Success",
		Message: "Problem Created Succesfully",
	}

	return out, nil
}

func (svc *problemService) GetProblemStatus(ctx context.Context, id string) (*models.ProblemStatusList, error) {
	var problems []*models.ProblemStatus

	db_problems, err := svc.repository.GetProblemsByUserId(ctx, svc.db, id)
	if err != nil {
		return nil, err
	}

	for _, problem := range db_problems {
		temp := models.ProblemStatus{
			ID:         problem.ID,
			Title:      problem.Title,
			Topic:      problem.Topic,
			Difficulty: problem.Difficulty,
			Status:     problem.Status,
		}

		problems = append(problems, &temp)
	}

	resp := models.ProblemStatusList{
		Problems: problems,
	}

	return &resp, nil
}

func (svc *problemService) GetProblemCandidateList(ctx context.Context, filter models.ProblemFilter) (*models.ProblemCandidateList, error) {
	var problems []*models.ProblemCandidateTable

	db_problems, err := svc.repository.GetCandidateProblemList(ctx, svc.db, filter)
	if err != nil {
		return nil, err
	}

	for _, problem := range db_problems {
		temp := models.ProblemCandidateTable{
			ID:         problem.ID,
			Title:      problem.Title,
			Topic:      problem.Topic,
			Difficulty: problem.Difficulty,
		}

		problems = append(problems, &temp)
	}

	resp := models.ProblemCandidateList{
		Problems: problems,
	}

	return &resp, nil
}

func (svc *problemService) GetProblemDetail(ctx context.Context, id string) (*models.ProblemDetail, error) {
	problem, err := svc.repository.GetCandidateById(ctx, svc.db, id)
	if err != nil {
		return nil, err
	}

	resp := models.ProblemDetail{
		ID:     problem.ID,
		Detail: problem.Detail,
	}

	return &resp, nil
}

func (svc *problemService) AcceptProblem(ctx context.Context, id string) (*models.ProblemCreationResponse, error) {
	value := &models.ProblemStatusUpdate{
		Id:     id,
		Status: "accepted",
	}

	err := svc.repository.UpdateProblemStatus(ctx, svc.db, value)
	if err != nil {
		return nil, err
	}

	out := &models.ProblemCreationResponse{
		Status:  "Success",
		Message: "Problem Updated Succesfully",
	}

	return out, err
}

func (svc *problemService) RejectProblem(ctx context.Context, id string) (*models.ProblemCreationResponse, error) {
	value := &models.ProblemStatusUpdate{
		Id:     id,
		Status: "rejected",
	}

	err := svc.repository.UpdateProblemStatus(ctx, svc.db, value)
	if err != nil {
		return nil, err
	}

	out := &models.ProblemCreationResponse{
		Status:  "Success",
		Message: "Problem Updated Succesfully",
	}

	return out, err
}

func (svc *problemService) GetProblemList(ctx context.Context, filter models.ProblemFilter) (*models.ProblemCandidateList, error) {
	var problems []*models.ProblemCandidateTable

	db_problems, err := svc.repository.GetProblemList(ctx, svc.db, filter)
	if err != nil {
		return nil, err
	}

	for _, problem := range db_problems {
		temp := models.ProblemCandidateTable{
			ID:         problem.ID,
			Title:      problem.Title,
			Topic:      problem.Topic,
			Difficulty: problem.Difficulty,
		}

		problems = append(problems, &temp)
	}

	resp := models.ProblemCandidateList{
		Problems: problems,
	}

	return &resp, nil
}
