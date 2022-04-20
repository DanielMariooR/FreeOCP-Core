package problem_test

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
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/problem"
)

var (
	status     = "accepted"
	id         = uuid.New().String()
	creator    = "DanielMr"
	title      = "intro to python"
	topic      = "programming"
	problem1   = uuid.New().String()
	problem2   = uuid.New().String()
	problem3   = uuid.New().String()
	difficulty = "mudah"
)

func TestProblemService_GetProblemCandidate(t *testing.T) {
	var (
		id = uuid.New().String()
	)

	type mockGetProblemCandidate struct {
		res *db_models.ProblemCandidate
		err error
	}

	type args struct {
		ctx       context.Context
		problemId string
	}

	tests := []struct {
		name                    string
		args                    args
		mockGetProblemCandidate mockGetProblemCandidate
		want                    *models.ProblemCandidate
		wantErr                 error
	}{
		{
			name: "Success to get problem candidate",
			args: args{
				context.TODO(),
				id,
			},
			mockGetProblemCandidate: mockGetProblemCandidate{
				res: &db_models.ProblemCandidate{
					ID:         id,
					Creator:    creator,
					Difficulty: difficulty,
					Title:      title,
					Topic:      topic,
					Type:       "pilgan",
					Status:     status,
					Detail:     "",
				},
			},
			want: &models.ProblemCandidate{
				ID:         id,
				Creator:    creator,
				Title:      title,
				Type:       "pilgan",
				Topic:      topic,
				Difficulty: difficulty,
				Status:     status,
				Detail:     "",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			problemRepoMock := new(mocks.ProblemRepository)
			svc := problem.NewService(sqlxDB)
			svc.InjectRepository(problemRepoMock)
			problemRepoMock.On("GetCandidateById", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetProblemCandidate.res, tt.mockGetProblemCandidate.err)

			got, err := svc.GetProblemCandidate(tt.args.ctx, tt.args.problemId)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestProblemService_CreateNewProblem(t *testing.T) {
	type mockInsertNewProblem struct {
		err error
	}

	type args struct {
		ctx   context.Context
		input *models.ProblemCreationInput
	}

	tests := []struct {
		name                 string
		args                 args
		mockInsertNewProblem mockInsertNewProblem
		want                 *models.ProblemCreationResponse
		wantErr              error
	}{
		{
			name: "Success to create new problem",
			args: args{
				context.TODO(),
				&models.ProblemCreationInput{
					Creator:    creator,
					Title:      title,
					Type:       "pilgan",
					Topic:      topic,
					Difficulty: difficulty,
					Detail:     "{}",
				},
			},
			mockInsertNewProblem: mockInsertNewProblem{
				err: nil,
			},
			want: &models.ProblemCreationResponse{
				Status:  "Success",
				Message: "Problem Created Succesfully",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			problemRepoMock := new(mocks.ProblemRepository)
			svc := problem.NewService(sqlxDB)
			svc.InjectRepository(problemRepoMock)
			problemRepoMock.On("InsertNewProblem", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockInsertNewProblem.err)

			got, err := svc.CreateNewProblem(tt.args.ctx, tt.args.input)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}

func TestProblemService_AcceptProblem(t *testing.T) {
	type args struct {
		ctx       context.Context
		problemId string
	}

	type mockUpdateProblemStatus struct {
		err error
	}

	tests := []struct {
		name                    string
		args                    args
		mockUpdateProblemStatus mockUpdateProblemStatus
		want                    *models.ProblemCreationResponse
		wantErr                 error
	}{
		{
			name: "Success to accept problems",
			args: args{
				context.TODO(),
				id,
			},
			mockUpdateProblemStatus: mockUpdateProblemStatus{
				err: nil,
			},
			want: &models.ProblemCreationResponse{
				Status:  "Success",
				Message: "Problem Updated Succesfully",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			problemRepoMock := new(mocks.ProblemRepository)
			svc := problem.NewService(sqlxDB)
			svc.InjectRepository(problemRepoMock)
			problemRepoMock.On("UpdateProblemStatus", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockUpdateProblemStatus.err)

			got, err := svc.AcceptProblem(tt.args.ctx, tt.args.problemId)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}

func TestProblemService_RejectProblem(t *testing.T) {
	type args struct {
		ctx       context.Context
		problemId string
	}

	type mockUpdateProblemStatus struct {
		err error
	}

	tests := []struct {
		name                    string
		args                    args
		mockUpdateProblemStatus mockUpdateProblemStatus
		want                    *models.ProblemCreationResponse
		wantErr                 error
	}{
		{
			name: "Success to accept problems",
			args: args{
				context.TODO(),
				id,
			},
			mockUpdateProblemStatus: mockUpdateProblemStatus{
				err: nil,
			},
			want: &models.ProblemCreationResponse{
				Status:  "Success",
				Message: "Problem Updated Succesfully",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			problemRepoMock := new(mocks.ProblemRepository)
			svc := problem.NewService(sqlxDB)
			svc.InjectRepository(problemRepoMock)
			problemRepoMock.On("UpdateProblemStatus", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockUpdateProblemStatus.err)

			got, err := svc.RejectProblem(tt.args.ctx, tt.args.problemId)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}

func TestProblemService_GetProblemDetail(t *testing.T) {
	type mockGetCandidateById struct {
		res *db_models.ProblemCandidate
		err error
	}

	type args struct {
		ctx       context.Context
		problemId string
	}

	tests := []struct {
		name                 string
		args                 args
		mockGetCandidateById mockGetCandidateById
		want                 *models.ProblemDetail
		wantErr              error
	}{
		{
			name: "Success to get problem detail by id",
			args: args{
				context.TODO(),
				id,
			},
			mockGetCandidateById: mockGetCandidateById{
				res: &db_models.ProblemCandidate{
					ID:         id,
					Title:      title,
					Topic:      topic,
					Type:       "pilgan",
					Difficulty: difficulty,
					Detail:     "{}",
					Status:     status,
					Creator:    creator,
				},
				err: nil,
			},
			want: &models.ProblemDetail{
				ID:     id,
				Detail: "{}",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			problemRepoMock := new(mocks.ProblemRepository)
			svc := problem.NewService(sqlxDB)
			svc.InjectRepository(problemRepoMock)
			problemRepoMock.On("GetCandidateById", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetCandidateById.res, tt.mockGetCandidateById.err)

			got, err := svc.GetProblemDetail(tt.args.ctx, tt.args.problemId)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestProblemService_GetProblemList(t *testing.T) {
	type args struct {
		ctx    context.Context
		filter *models.ProblemFilter
	}

	type mockGetProblemList struct {
		res []*db_models.ProblemCandidate
		err error
	}

	tests := []struct {
		name               string
		args               args
		mockGetProblemList mockGetProblemList
		want               *models.ProblemCandidateList
		wantErr            error
	}{
		{
			name: "Success to get problem list",
			args: args{
				context.TODO(),
				&models.ProblemFilter{},
			},
			mockGetProblemList: mockGetProblemList{
				res: []*db_models.ProblemCandidate{
					{
						ID:         problem1,
						Creator:    creator,
						Title:      title,
						Topic:      topic,
						Difficulty: difficulty,
						Type:       "pilgan",
						Status:     status,
						Detail:     "{}",
					},
				},
				err: nil,
			},
			want: &models.ProblemCandidateList{
				Problems: []*models.ProblemCandidateTable{
					{
						ID:         problem1,
						Title:      title,
						Topic:      topic,
						Difficulty: difficulty,
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			problemRepoMock := new(mocks.ProblemRepository)
			svc := problem.NewService(sqlxDB)
			svc.InjectRepository(problemRepoMock)
			problemRepoMock.On("GetProblemList", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetProblemList.res, tt.mockGetProblemList.err)

			got, err := svc.GetProblemList(tt.args.ctx, *tt.args.filter)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}

func TestProblemService_GetProblemCandidateList(t *testing.T) {
	type args struct {
		ctx    context.Context
		filter *models.ProblemFilter
	}

	type mockGetCandidateProblemList struct {
		res []*db_models.ProblemCandidate
		err error
	}

	tests := []struct {
		name                        string
		args                        args
		mockGetCandidateProblemList mockGetCandidateProblemList
		want                        *models.ProblemCandidateList
		wantErr                     error
	}{
		{
			name: "Success to get candidate problem list",
			args: args{
				context.TODO(),
				&models.ProblemFilter{},
			},
			mockGetCandidateProblemList: mockGetCandidateProblemList{
				res: []*db_models.ProblemCandidate{
					{
						ID:         problem1,
						Creator:    creator,
						Title:      title,
						Topic:      topic,
						Difficulty: difficulty,
						Type:       "pilgan",
						Status:     status,
						Detail:     "{}",
					},
				},
				err: nil,
			},
			want: &models.ProblemCandidateList{
				Problems: []*models.ProblemCandidateTable{
					{
						ID:         problem1,
						Title:      title,
						Topic:      topic,
						Difficulty: difficulty,
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			problemRepoMock := new(mocks.ProblemRepository)
			svc := problem.NewService(sqlxDB)
			svc.InjectRepository(problemRepoMock)
			problemRepoMock.On("GetCandidateProblemList", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetCandidateProblemList.res, tt.mockGetCandidateProblemList.err)

			got, err := svc.GetProblemCandidateList(tt.args.ctx, *tt.args.filter)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}
