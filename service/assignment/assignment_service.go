package assignment

import (
	"context"
	"encoding/json"
	"math"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/assignment/assignment_repository"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/course/course_repository"
)

type assignmentService struct {
	db         *sqlx.DB
	repository assignment_repository.AssignmentRepository
}

func NewService(db *sqlx.DB) AssignmentService {
	return &assignmentService{
		db: db,
	}
}

func (svc *assignmentService) GetAssignment(ctx context.Context, id string) (*models.AssignmentResponse, error) {
	assignment, err := svc.repository.GetAssignmentById(ctx, svc.db, id)
	if err != nil {
		return nil, err
	}

	out := models.AssignmentResponse{
		ID:         assignment.ID,
		Creator:    assignment.Creator,
		Title:      assignment.Title,
		Duration:   assignment.Duration,
		Topic:      assignment.Topic,
		Difficulty: assignment.Difficulty,
	}

	problems, err := svc.repository.GetAssignmentProblemsById(ctx, svc.db, id)
	if err != nil {
		return nil, err
	}

	var resp_problems []*models.ProblemTypeDetail
	for _, problem := range problems {
		var data map[string]interface{}

		err := json.Unmarshal([]byte(problem.Detail), &data)
		if err != nil {
			panic(err)
		}

		temp := models.ProblemTypeDetail{
			ID:       problem.ID,
			Type:     problem.Type,
			Choice:   data["choice"],
			Question: data["question"],
		}

		resp_problems = append(resp_problems, &temp)
	}

	out.Problems = resp_problems
	return &out, nil
}

func (svc *assignmentService) CreateAssignment(ctx context.Context, input *models.AssignmentCreation) (*models.AssignmentCreationResponse, error) {
	newId := uuid.New().String()
	assignment := db_models.AssignmentCreation{}

	assignment.Desc = db_models.Assignment{
		ID:         newId,
		Creator:    input.Desc.Creator,
		Title:      input.Desc.Title,
		Duration:   input.Desc.Duration,
		Topic:      input.Desc.Topic,
		Difficulty: input.Desc.Difficulty,
	}

	for _, problem := range input.Problems {
		assignment.Problems = append(assignment.Problems, db_models.AssignmentProblem{
			AssignmentID: newId,
			ProblemID:    problem.ProblemID,
		})
	}

	err := svc.repository.InsertAssignment(ctx, svc.db, &assignment)
	if err != nil {
		return nil, err
	}

	resp := &models.AssignmentCreationResponse{
		Status:  "Success",
		Message: "Section Created Succesfully",
		ID:      newId,
	}

	return resp, nil
}

func (svc *assignmentService) CalculateScore(ctx context.Context, answers *models.AssignmentSubmission) (int, error) {
	db_problems, err := svc.repository.GetAssignmentProblemsById(ctx, svc.db, answers.ID)
	if err != nil {
		return -1, err
	}

	correct := 0.0
	total := float64(len(db_problems))

	for _, problem := range db_problems {
		var data map[string]interface{}

		err := json.Unmarshal([]byte(problem.Detail), &data)
		if err != nil {
			panic(err)
		}

		solution := models.ProblemTypeDetailAnswer{
			ID:     problem.ID,
			Type:   problem.Type,
			Choice: data["choice"],
			Answer: data["answer"],
		}

		for _, answer := range answers.Answers {
			if answer.ID != solution.ID {
				continue
			}

			ans_ans := reflect.ValueOf(answer.Answer)
			sol_ans := reflect.ValueOf(solution.Answer)
			sol_cho := reflect.ValueOf(solution.Choice)

			if answer.Type == "pilgan" || answer.Type == "checkbox" {
				check := true
				for i := 0; i < ans_ans.Len(); i++ {
					ans_idx := ans_ans.Index(i).Interface()
					sol_idx := sol_ans.Index(i).Interface()
					if ans_idx != sol_idx {
						check = false
						break
					}
				}
				if check {
					correct++
				}
			} else if answer.Type == "isian" {
				ans_string := ans_ans.Index(0).Interface().(string)
				cho_string := sol_cho.Index(0).Interface().(string)

				if sol_ans.Index(0).Interface().(float64) == 0 {
					ans_string = strings.ToLower(ans_string)
					cho_string = strings.ToLower(cho_string)
				}

				if ans_string == cho_string {
					correct++
				}
			} else if answer.Type == "plist" {
				check := true
				for i := 0; i < ans_ans.Len(); i++ {
					ans_idx := ans_ans.Index(i).Interface().(string)
					cho_idx := sol_cho.Index(i).Interface().(string)
					if ans_idx != cho_idx {
						check = false
						break
					}
				}
				if check {
					correct++
				}
			}
			break
		}
	}

	score := int(math.Round(correct / total * 100))

	return score, nil
}

func (svc *assignmentService) GetScore(ctx context.Context, userId string, answers *models.AssignmentSubmission) (*models.AssignmentScore, error) {
	score, err := svc.CalculateScore(ctx, answers)
	if err != nil {
		return nil, err
	}

	materialId := answers.ID
	courseRepo := course_repository.NewRepository()

	courseId, err := courseRepo.GetCourseIDByMaterialID(ctx, svc.db, materialId)
	if err != nil {
		return nil, err
	}

	err = courseRepo.StoreUserProgress(ctx, svc.db, answers.ID, courseId, userId, score)
	if err != nil {
		return nil, err
	}

	resp := &models.AssignmentScore{
		Score: score,
	}

	return resp, nil
}
