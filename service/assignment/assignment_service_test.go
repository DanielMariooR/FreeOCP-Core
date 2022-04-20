package assignment_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/mocks"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/assignment"
)

var (
	id         = uuid.New().String()
	creator    = "DanielMr"
	title      = "intro to python"
	duration   = 2000
	topic      = "programming"
	difficulty = "mudah"
	problem1   = uuid.New().String()
	problem2   = uuid.New().String()
	problem3   = uuid.New().String()
)

func TestAssignmentService_GetAssignment(t *testing.T) {
	type mockGetAssignment struct {
		res *db_models.Assignment
		err error
	}

	type mockGetAssignmentProblems struct {
		res []*db_models.ProblemTypeDetail
		err error
	}

	type args struct {
		ctx          context.Context
		assignmentId string
	}

	tests := []struct {
		name                      string
		args                      args
		mockGetAssignment         mockGetAssignment
		mockGetAssignmentProblems mockGetAssignmentProblems
		want                      *models.AssignmentResponse
		wantErr                   error
	}{
		{
			name: "Success to Get Assignment",
			args: args{
				context.TODO(),
				id,
			},
			mockGetAssignment: mockGetAssignment{
				res: &db_models.Assignment{
					ID:         id,
					Creator:    creator,
					Topic:      topic,
					Title:      title,
					Duration:   duration,
					Difficulty: difficulty,
				},
				err: nil,
			},
			mockGetAssignmentProblems: mockGetAssignmentProblems{
				res: []*db_models.ProblemTypeDetail{
					{
						ID:     problem1,
						Detail: "{}",
						Type:   "isian",
					},
					{
						ID:     problem2,
						Detail: "{}",
						Type:   "checkbox",
					},
					{
						ID:     problem3,
						Detail: "{}",
						Type:   "pilgan",
					},
				},
				err: nil,
			},
			want: &models.AssignmentResponse{
				ID:         id,
				Creator:    creator,
				Title:      title,
				Duration:   duration,
				Difficulty: difficulty,
				Topic:      topic,
				Problems: []*models.ProblemTypeDetail{
					{
						ID:   problem1,
						Type: "isian",
					},
					{
						ID:   problem2,
						Type: "checkbox",
					},
					{
						ID:   problem3,
						Type: "pilgan",
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			assignmentRepoMock := new(mocks.AssignmentRepository)
			svc := assignment.NewService(sqlxDB)
			svc.InjectAssignmentRepository(assignmentRepoMock)

			assignmentRepoMock.On("GetAssignmentById", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetAssignment.res, tt.mockGetAssignment.err)
			assignmentRepoMock.On("GetAssignmentProblemsById", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetAssignmentProblems.res, tt.mockGetAssignmentProblems.err)

			got, err := svc.GetAssignment(tt.args.ctx, tt.args.assignmentId)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestAssignmentService_CreateAssignment(t *testing.T) {
	type mockInsertAssignment struct {
		err error
	}

	type args struct {
		ctx   context.Context
		input *models.AssignmentCreation
	}

	tests := []struct {
		name                 string
		args                 args
		mockInsertAssignment mockInsertAssignment
		want                 *models.AssignmentCreationResponse
		wantErr              error
	}{
		{
			name: "Success to insert assignment",
			args: args{
				context.TODO(),
				&models.AssignmentCreation{
					Desc: models.Assignment{
						ID:         id,
						Title:      title,
						Topic:      topic,
						Duration:   duration,
						Creator:    creator,
						Difficulty: difficulty,
					},
					Problems: []models.AssignmentProblem{
						{
							ProblemID: problem1,
						},
						{
							ProblemID: problem2,
						},
						{
							ProblemID: problem3,
						},
					},
				},
			},
			mockInsertAssignment: mockInsertAssignment{
				err: nil,
			},
			want: &models.AssignmentCreationResponse{
				Status:  "Success",
				Message: "Section Created Succesfully",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			assignmentRepoMock := new(mocks.AssignmentRepository)
			svc := assignment.NewService(sqlxDB)
			svc.InjectAssignmentRepository(assignmentRepoMock)
			assignmentRepoMock.On("InsertAssignment", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockInsertAssignment.err)

			got, err := svc.CreateAssignment(tt.args.ctx, tt.args.input)

			assert.Equal(t, tt.want.Status, got.Status, tt.name)
			assert.Equal(t, tt.want.Message, got.Message, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}
