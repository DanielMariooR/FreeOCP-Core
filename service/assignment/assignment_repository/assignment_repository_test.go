package assignment_repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	er "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/error"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/assignment/assignment_repository"
)

func TestAssignmentRepository_GetAssignmentById(t *testing.T) {
	var (
		assignmentId = uuid.New().String()
		creator      = "DanielMr"
		title        = "intro to python"
		duration     = 2000
		topic        = "programming"
		difficulty   = "mudah"
	)

	type args struct {
		ctx context.Context
		id  string
	}

	type mockSelect struct {
		assignmentById *db_models.Assignment
		err            error
	}

	tests := []struct {
		name       string
		args       args
		mockSelect mockSelect
		want       *db_models.Assignment
		wantErr    error
	}{
		{
			name: "Success to return assignment by ID",
			args: args{
				context.TODO(),
				assignmentId,
			},
			mockSelect: mockSelect{
				assignmentById: &db_models.Assignment{
					ID:         assignmentId,
					Creator:    creator,
					Title:      title,
					Duration:   duration,
					Topic:      topic,
					Difficulty: difficulty,
				},
			},
			want: &db_models.Assignment{
				ID:         assignmentId,
				Creator:    creator,
				Title:      title,
				Duration:   duration,
				Topic:      topic,
				Difficulty: difficulty,
			},
			wantErr: nil,
		},
		{
			name: "Fail to return assignment by ID",
			args: args{
				context.TODO(),
				assignmentId,
			},
			mockSelect: mockSelect{
				err: sql.ErrNoRows,
			},
			want:    nil,
			wantErr: er.NewError(fmt.Errorf("%s", "Assignment Not Found!"), http.StatusBadRequest, nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "sqlmock")

			if tt.mockSelect.assignmentById != nil {
				data := tt.mockSelect.assignmentById
				rows := sqlmock.NewRows([]string{
					"id",
					"creator",
					"title",
					"duration",
					"topic",
					"difficulty",
				})
				rows.AddRow(data.ID, data.Creator, data.Title, data.Duration, data.Topic, data.Difficulty)
				mock.ExpectQuery(`SELECT id, creator, title, duration, topic, difficulty FROM assignment`).WillReturnRows(rows)
				mock.ExpectCommit()

			}

			if tt.mockSelect.err != nil {
				mock.ExpectQuery(`SELECT id, creator, title, duration, topic, difficulty FROM assignment`).WillReturnError(tt.mockSelect.err)
			}

			r := assignment_repository.NewRepository()
			got, err := r.GetAssignmentById(tt.args.ctx, sqlxDB, tt.args.id)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}
}

func TestAssignmentRepository_GetAssignmentProblemsById(t *testing.T) {
	var (
		assignmentId = uuid.New().String()
		problem1     = uuid.New().String()
		problem2     = uuid.New().String()
		problem3     = uuid.New().String()
		problems     []*db_models.ProblemTypeDetail
	)

	type args struct {
		ctx          context.Context
		assignmentId string
	}

	type mockSelect struct {
		assignmentProblemsById []*db_models.ProblemTypeDetail
		err                    error
	}

	tests := []struct {
		name       string
		args       args
		mockSelect mockSelect
		want       []*db_models.ProblemTypeDetail
		wantErr    error
	}{
		{
			name: "Success to get assignment problems by ID",
			args: args{
				context.TODO(),
				assignmentId,
			},
			mockSelect: mockSelect{
				assignmentProblemsById: []*db_models.ProblemTypeDetail{
					{
						ID:     problem1,
						Detail: "",
						Type:   "checkbox",
					},
					{
						ID:     problem2,
						Detail: "",
						Type:   "pilgan",
					},
					{
						ID:     problem3,
						Detail: "",
						Type:   "isian",
					},
				},
			},
			want: []*db_models.ProblemTypeDetail{
				{
					ID:     problem1,
					Detail: "",
					Type:   "checkbox",
				},
				{
					ID:     problem2,
					Detail: "",
					Type:   "pilgan",
				},
				{
					ID:     problem3,
					Detail: "",
					Type:   "isian",
				},
			},
			wantErr: nil,
		},
		{
			name: "Success to get empty assignment",
			args: args{
				context.TODO(),
				assignmentId,
			},
			mockSelect: mockSelect{
				err: sql.ErrNoRows,
			},
			want:    problems,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "sqlmock")

			queryString := "SELECT p.id as id, p.detail as detail, cp.type as type FROM Detail_Problem p INNER JOIN assignment_problems ap on p.id = ap.problem_id INNER JOIN Candidate_Problem cp on cp.id = p.id WHERE ap.assignment_id = ?"

			if tt.mockSelect.assignmentProblemsById != nil {
				rows := sqlmock.NewRows([]string{"id", "detail", "type"})
				for _, row := range *&tt.mockSelect.assignmentProblemsById {
					rows.AddRow(row.ID, row.Detail, row.Type)
				}

				mock.ExpectQuery(queryString).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			if tt.mockSelect.err != nil {
				mock.ExpectQuery(queryString).WillReturnError(tt.mockSelect.err)
			}

			r := assignment_repository.NewRepository()
			got, err := r.GetAssignmentProblemsById(tt.args.ctx, sqlxDB, tt.args.assignmentId)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}

func TestAssignmentRepository_InsertAssignmentDesc(t *testing.T) {
	var (
		assignmentId = uuid.New().String()
		creator      = "DanielMr"
		title        = "intro to python"
		duration     = 2000
		topic        = "programming"
		difficulty   = "mudah"
	)

	type args struct {
		ctx   context.Context
		input *db_models.Assignment
	}

	type mockExec struct {
		assignmentByID *db_models.Assignment
		err            error
	}

	tests := []struct {
		name     string
		args     args
		mockExec mockExec
		want     *db_models.Assignment
		wantErr  error
	}{
		{
			name: "Insert success",
			args: args{
				context.TODO(),
				&db_models.Assignment{
					ID:         assignmentId,
					Creator:    creator,
					Title:      title,
					Duration:   duration,
					Topic:      topic,
					Difficulty: difficulty,
				},
			},
			mockExec: mockExec{
				assignmentByID: &db_models.Assignment{
					ID:         assignmentId,
					Creator:    creator,
					Title:      title,
					Duration:   duration,
					Topic:      topic,
					Difficulty: difficulty,
				},
				err: nil,
			},
			want: &db_models.Assignment{
				ID:         assignmentId,
				Creator:    creator,
				Title:      title,
				Duration:   duration,
				Topic:      topic,
				Difficulty: difficulty,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "sqlmock")

			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO assignment (id,creator,title,duration,topic,difficulty) VALUES (?,?,?,?,?,?)`)).WillReturnResult(sqlmock.NewResult(1, 1))

			r := assignment_repository.NewRepository()
			err = r.InsertAssignmentDesc(tt.args.ctx, sqlxDB, tt.args.input)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}

func TestAssignmentRepository_InsertAssignmentProblem(t *testing.T) {
	var (
		assignmentId = uuid.New().String()
		problem1     = uuid.New().String()
		problem2     = uuid.New().String()
		problem3     = uuid.New().String()
	)

	input := []db_models.AssignmentProblem{
		{
			AssignmentID: assignmentId,
			ProblemID:    problem1,
		},
		{
			AssignmentID: assignmentId,
			ProblemID:    problem2,
		},
		{
			AssignmentID: assignmentId,
			ProblemID:    problem3,
		},
	}

	type args struct {
		ctx   context.Context
		input *[]db_models.AssignmentProblem
	}

	type mockExec struct {
		assignmentProblemByID []*db_models.AssignmentProblem
		err                   error
	}

	tests := []struct {
		name     string
		args     args
		mockExec mockExec
		wantErr  error
	}{
		{
			name: "Success to insert Assignment Problems",
			args: args{
				context.TODO(),
				&input,
			},
			mockExec: mockExec{
				assignmentProblemByID: []*db_models.AssignmentProblem{
					{
						AssignmentID: assignmentId,
						ProblemID:    problem1,
					},
					{
						AssignmentID: assignmentId,
						ProblemID:    problem2,
					},
					{
						AssignmentID: assignmentId,
						ProblemID:    problem3,
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "sqlmock")

			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO assignment_problem (assignment_id,problem_id) VALUES (?,?),(?,?),(?,?)`)).WillReturnResult(sqlmock.NewResult(1, 1))

			r := assignment_repository.NewRepository()
			err = r.InsertAssignmentProblem(tt.args.ctx, sqlxDB, tt.args.input)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}

func TestAssignmentRepository_GetTableName(t *testing.T) {
	r := assignment_repository.NewRepository()

	expectedTableName := "assignment"
	assert.Equal(t, expectedTableName, r.GetTableName())
}

func TestAssignmentRepository_GetProblemTableName(t *testing.T) {
	r := assignment_repository.NewRepository()

	expectedTableName := "assignment_problem"
	assert.Equal(t, expectedTableName, r.GetProblemTableName())
}
