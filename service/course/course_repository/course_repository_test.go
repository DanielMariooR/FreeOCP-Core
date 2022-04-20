package course_repository_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/pagination"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/course/course_repository"
)

var (
	courseId = uuid.New().String()
	
	creatorId = uuid.New().String()

	userId = uuid.New().String()
	userId2 = uuid.New().String()

	materialId = uuid.New().String()
	materialId2 = uuid.New().String()
	materialId3 = uuid.New().String()

	sectionId = uuid.New().String()
)

func TestCourseRepository_GetCourseById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}

	type mockQuery struct {
		res *db_models.Course
		err  error
	}

	tests := []struct {
		name     string
		args     args
		mock     mockQuery
		want    *db_models.Course
		wantErr  error
	}{
		{
			name: "[GetCourseById] Success to return course by ID",
			args: args {
				context.TODO(),
				courseId,
			},
			mock: mockQuery{
				res: &db_models.Course{
					ID:          courseId,
					CourseName:  "",
					Description: "",
					Thumbnail:   "",
					Creator:     creatorId,
				},
				err: nil,
			},
			want: &db_models.Course{
				ID:          courseId,
				CourseName:  "",
				Description: "",
				Thumbnail:   "",
				Creator:     creatorId,
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

			selectQuery := regexp.QuoteMeta("SELECT id, course_name, description, thumbnail, creator FROM Course WHERE id = ?")

			if tt.mock.res != nil {
				data := tt.mock.res
				rows := sqlmock.NewRows([]string{
					"id",
					"course_name",
					"description",
					"thumbnail",
					"creator",
				})

				rows.AddRow(data.ID, data.CourseName, data.Description, data.Thumbnail, data.Creator)

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, err := r.GetCourseById(tt.args.ctx, sqlxDB, tt.args.id)

			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_GetCompletedCourseById(t *testing.T) {
	type args struct {
		ctx     context.Context
		meta   *pagination.Meta
		userid  string
	}

	type mockQuery struct {
		res   []*db_models.Course
		count uint64
		err   error
	}

	tests := []struct {
		name      string
		args      args
		mock 	    mockQuery
		wantRes   []*db_models.Course
		wantCount	uint64
		wantErr   error
	}{
		{
			name: "[GetCompletedCourseById] Success to get completed courses by user ID",
			args: args{
				context.TODO(),
				&pagination.Meta{
					Limit: 3,
					Page:  3,
				},
				userId,
			},
			mock: mockQuery{
				res: []*db_models.Course {
					{
						ID:          courseId,
						CourseName:  "",
						Description: "",
						Thumbnail:   "",
						Creator:     creatorId,
					},
				},
				count: 1,
				err: nil,
			},
			wantRes: []*db_models.Course {
				{
					ID:          courseId,
					CourseName:  "",
					Description: "",
					Thumbnail:   "",
					Creator:     creatorId,
				},
			},
			wantCount: 1,
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

			selectQuery := regexp.QuoteMeta(`SELECT id, course_name, description, thumbnail, creator FROM Course`)
			
			countQuery := regexp.QuoteMeta(`SELECT count(id) FROM Course`)

			if tt.mock.res != nil {
				rows := sqlmock.NewRows([]string{
					"id",
					"course_name",
					"description",
					"thumbnail",
					"creator",
				})

				for _, row := range tt.mock.res {
					rows.AddRow(row.ID, row.CourseName, row.Description, row.Thumbnail, row.Creator)
				}

				countRow := sqlmock.NewRows([]string {
					"count(id)",
				})

				countRow.AddRow(tt.mock.count)

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectQuery(countQuery).WillReturnRows(countRow)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, cou, err := r.GetCompletedCourseByUserID(tt.args.ctx, sqlxDB, tt.args.meta, tt.args.userid)
			assert.Equal(t, tt.wantRes, got, tt.name)
			assert.Equal(t, tt.wantCount, cou, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_GetOnProgressCourseByUserID(t *testing.T) {
	type args struct {
		ctx     context.Context
		meta   *pagination.Meta
		userid  string
	}

	type mockQuery struct {
		res   []*db_models.Course
		count uint64
		err   error
	}

	tests := []struct {
		name      string
		args      args
		mock 	    mockQuery
		wantRes   []*db_models.Course
		wantCount	uint64
		wantErr   error
	}{
		{
			name: "[GetOnProgressCourseById] Success to get on progress courses by user ID",
			args: args{
				context.TODO(),
				&pagination.Meta{
					Limit: 3,
					Page:  3,
				},
				userId,
			},
			mock: mockQuery{
				res: []*db_models.Course {
					{
						ID:          courseId,
						CourseName:  "",
						Description: "",
						Thumbnail:   "",
						Creator:     creatorId,
					},
				},
				count: 1,
				err: nil,
			},
			wantRes: []*db_models.Course {
				{
					ID:          courseId,
					CourseName:  "",
					Description: "",
					Thumbnail:   "",
					Creator:     creatorId,
				},
			},
			wantCount: 1,
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

			selectQuery := regexp.QuoteMeta(`SELECT id, course_name, description, thumbnail, creator FROM Course`)
			
			countQuery := regexp.QuoteMeta(`SELECT count(id) FROM Course`)

			if tt.mock.res != nil {
				rows := sqlmock.NewRows([]string{
					"id",
					"course_name",
					"description",
					"thumbnail",
					"creator",
				})

				for _, row := range tt.mock.res {
					rows.AddRow(row.ID, row.CourseName, row.Description, row.Thumbnail, row.Creator)
				}

				countRow := sqlmock.NewRows([]string {
					"count(id)",
				})

				countRow.AddRow(tt.mock.count)

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectQuery(countQuery).WillReturnRows(countRow)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, cou, err := r.GetOnProgressCourseByUserID(tt.args.ctx, sqlxDB, tt.args.meta, tt.args.userid)
			assert.Equal(t, tt.wantRes, got, tt.name)
			assert.Equal(t, tt.wantCount, cou, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_GetCourseList(t *testing.T) {
	type args struct {
		ctx     context.Context
		meta   *pagination.Meta
	}

	type mockQuery struct {
		res   []*db_models.Course
		count uint64
		err   error
	}

	tests := []struct {
		name      string
		args      args
		mock 	    mockQuery
		wantRes   []*db_models.Course
		wantCount	uint64
		wantErr   error
	}{
		{
			name: "[GetCourseList] Success to get on progress courses by user ID",
			args: args{
				context.TODO(),
				&pagination.Meta{
					Limit: 3,
					Page:  3,
				},
			},
			mock: mockQuery{
				res: []*db_models.Course {
					{
						ID:          courseId,
						CourseName:  "",
						Description: "",
						Thumbnail:   "",
						Creator:     creatorId,
					},
				},
				count: 1,
				err: nil,
			},
			wantRes: []*db_models.Course {
				{
					ID:          courseId,
					CourseName:  "",
					Description: "",
					Thumbnail:   "",
					Creator:     creatorId,
				},
			},
			wantCount: 1,
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

			selectQuery := regexp.QuoteMeta(`SELECT id, course_name, description, thumbnail, creator FROM Course`)
			
			countQuery := regexp.QuoteMeta(`SELECT count(id) FROM Course`)

			if tt.mock.res != nil {
				rows := sqlmock.NewRows([]string{
					"id",
					"course_name",
					"description",
					"thumbnail",
					"creator",
				})

				for _, row := range tt.mock.res {
					rows.AddRow(row.ID, row.CourseName, row.Description, row.Thumbnail, row.Creator)
				}

				countRow := sqlmock.NewRows([]string {
					"count(id)",
				})

				countRow.AddRow(tt.mock.count)

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectQuery(countQuery).WillReturnRows(countRow)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, cou, err := r.GetCourseList(tt.args.ctx, sqlxDB, tt.args.meta)
			assert.Equal(t, tt.wantRes, got, tt.name)
			assert.Equal(t, tt.wantCount, cou, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_GetCourseSyllabusByCourseID(t *testing.T) {

	type args struct {
		ctx      context.Context
		courseId string
	}

	type mockQuery struct {
		res []*db_models.Syllabus
		err error
	}

	tests := []struct {
		name    string
		args    args
		mock 	  mockQuery
		want    []*db_models.Syllabus
		wantErr error
	}{
		{
			name: "[GetCourseSyllabusByCourseID] Success to get course syllabus by course ID",
			args: args{
				context.TODO(),
				courseId,
			},
			mock: mockQuery{
				res: []*db_models.Syllabus{
					{
						ID:        materialId,
						CourseID:  courseId,
						Name:      "",
						Type:      "section",
						SectionID: nil,
					},
					{
						ID:        materialId2,
						CourseID:  courseId,
						Name:      "",
						Type:      "video",
						SectionID: &materialId,
					},
					{
						ID:        materialId3,
						CourseID:  courseId,
						Name:      "",
						Type:      "assignment",
						SectionID: &materialId,
					},
				},
				err: nil,
			},
			want: []*db_models.Syllabus{
				{
					ID:        materialId,
					CourseID:  courseId,
					Name:      "",
					Type:      "section",
					SectionID: nil,
				},
				{
					ID:        materialId2,
					CourseID:  courseId,
					Name:      "",
					Type:      "video",
					SectionID: &materialId,
				},
				{
					ID:        materialId3,
					CourseID:  courseId,
					Name:      "",
					Type:      "assignment",
					SectionID: &materialId,
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

			selectQuery := regexp.QuoteMeta(`SELECT id, course_id, name, type, section_id FROM course_material`)

			if tt.mock.res != nil {
				rows := sqlmock.NewRows([]string{
					"id",
					"course_id",
					"name",
					"type",
					"section_id",
				})

				for _, row := range tt.mock.res {
					rows.AddRow(row.ID, row.CourseID, row.Name, row.Type, row.SectionID)
				}

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, err := r.GetCourseSyllabusByCourseID(tt.args.ctx, sqlxDB, tt.args.courseId)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_GetCourseMaterialByCourseIDAndSectionID(t *testing.T) {

	type args struct {
		ctx       context.Context
		courseId  string
		sectionId string
	}

	type mockQuery struct {
		res []*db_models.Material
		err	error
	}

	tests := []struct {
		name    string
		args    args
		mock    mockQuery
		want    []*db_models.Material
		wantErr error
	}{
		{
			name: "[GetCourseMaterialByCourseIDAndSectionID] Success to get course material by course ID and section ID",
			args: args{
				context.TODO(),
				courseId,
				sectionId,
			},
			mock: mockQuery{
				res: []*db_models.Material{
					{
						ID:          materialId,
						CourseID:    courseId,
						Name:        "",
						Type:        "",
						SectionID:   sectionId,
						Content:     "",
						ContentText: "",
					},
				},
				err: nil,
			},
			want: []*db_models.Material{
				{
					ID:          materialId,
					CourseID:    courseId,
					Name:        "",
					Type:        "",
					SectionID:   sectionId,
					Content:     "",
					ContentText: "",
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

			selectQuery := regexp.QuoteMeta(`SELECT id, course_id, name, type, section_id, content, content_text FROM course_material`)

			if tt.mock.res != nil {
				rows := sqlmock.NewRows([]string{
					"id",
					"course_id",
					"name",
					"type",
					"section_id",
					"content",
					"content_text",
				})

				for _, row := range tt.mock.res {
					rows.AddRow(row.ID, row.CourseID, row.Name, row.Type, row.SectionID, row.Content, row.ContentText)
				}

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, err := r.GetCourseMaterialByCourseIDAndSectionID(tt.args.ctx, sqlxDB, tt.args.courseId, tt.args.sectionId)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_IsUserEnrolledToCourse(t *testing.T) {

	type args struct {
		ctx      context.Context
		userId   string
		courseId string
	}

	type mockQuery struct {
		is  bool
		err error
	}

	tests := []struct {
		name    string
		args    args
		mock 	  mockQuery
		want    bool
		wantErr error
	}{
		{
			name: "[IsUserEnrolledToCourse] User 1 is enrolled to Course 1.",
			args: args{
				context.TODO(),
				userId,
				courseId,
			},
			mock: mockQuery{
				is: true,
				err: nil,
			},
			want: true,
			wantErr: nil,
		},
		{
			name: "[IsUserEnrolledToCourse] User 2 is not enrolled to Course 1.",
			args: args{
				context.TODO(),
				userId2,
				courseId,
			},
			mock: mockQuery{
				is: false,
				err: nil,
			},
			want: false,
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

			selectQuery := regexp.QuoteMeta(`SELECT count(*) FROM on_progress_course WHERE user_id = ? AND course_id = ?`)

			if tt.mock.is == true || tt.mock.is == false{
				rows := sqlmock.NewRows([]string{
					"count(*)",
				})

				if tt.mock.is == true {
					rows.AddRow(1);
				} else {
					rows.AddRow(0)
				}
			
				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, err := r.IsUserEnrolledToCourse(tt.args.ctx, sqlxDB, tt.args.userId, tt.args.courseId)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_InsertEnrollment(t *testing.T) {

	type args struct {
		ctx    context.Context
		input *models.EnrollInput
	}

	type mockQuery struct {
		data *models.EnrollInput
		err   error
	}

	tests := []struct {
		name    string
		args    args
		mock    mockQuery
		wantErr error
	}{
		{
			name: "[InsertEnrollment] Success to insert Enrollment",
			args: args{
				context.TODO(),
				&models.EnrollInput{
					UserID: userId,
					CourseID: courseId,
				},
			},
			mock: mockQuery{
				data: &models.EnrollInput{
					UserID: userId,
					CourseID: courseId,
				},
				err: nil,
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

			execQuery := regexp.QuoteMeta(`INSERT INTO on_progress_course (user_id,course_id) VALUES (?,?)`)

			mock.ExpectExec(execQuery).WillReturnResult(sqlmock.NewResult(1, 1))

			r := course_repository.NewRepository()
			err = r.InsertEnrollment(tt.args.ctx, sqlxDB, tt.args.input)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_GetCourseIDByMaterialID(t *testing.T) {

	type args struct {
		ctx        context.Context
		materialId string
	}

	type mockQuery struct {
		data []*db_models.Syllabus
		err  error
	}

	tests := []struct {
		name    string
		args    args
		mock 	  mockQuery
		want    string
		wantErr error
	}{
		{
			name: "[GetCourseIDByMaterialID] Success to get course ID by material ID",
			args: args {
				context.TODO(),
				materialId,
			},
			mock: mockQuery{
				data: []*db_models.Syllabus{
					{
						ID:        materialId,
						CourseID:  courseId,
						Name:      "",
						Type:      "section",
						SectionID: nil,
					},
				},
				err: nil,
			},
			want: courseId,
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

			selectQuery := regexp.QuoteMeta(`SELECT id, course_id, name, type, section_id FROM course_material`)

			if tt.mock.data != nil{
				rows := sqlmock.NewRows([]string{
					"id",
					"course_id",
					"name",
					"type",
					"section_id",
				})

				for _, row := range tt.mock.data {
					rows.AddRow(row.ID, row.CourseID, row.Name, row.Type, row.SectionID)
				}

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, err := r.GetCourseIDByMaterialID(tt.args.ctx, sqlxDB, tt.args.materialId)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_StoreUserProgress(t *testing.T) {

	type args struct {
		ctx        context.Context
		materialId string
		courseId   string
		userId     string
		score      int
	}

	type mockQuery struct {
		data *db_models.UserProgress
		err   error
	}

	tests := []struct {
		name    string
		args    args
		mock    mockQuery
		wantErr error
	}{
		{
			name: "[StoreUserProgress] Success to insert user progress.",
			args: args{
				context.TODO(),
				materialId,
				courseId,
				userId,
				100,
			},
			mock: mockQuery{
				data: &db_models.UserProgress{
					UserID:     userId,
					CourseID:   courseId,
					MaterialID: materialId,
					Score:      100,
				},
				err: nil,
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

			execQuery := regexp.QuoteMeta(`INSERT INTO user_progress (user_id,course_id,material_id,score) VALUES (?,?,?,?)`)

			mock.ExpectExec(execQuery).WillReturnResult(sqlmock.NewResult(1, 1))

			r := course_repository.NewRepository()
			err = r.StoreUserProgress(tt.args.ctx, sqlxDB,tt.args.materialId, tt.args.courseId, tt.args.userId, tt.args.score)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_GetUserProgress(t *testing.T) {

	type args struct {
		ctx      context.Context
		userId   string
		courseId string
	}

	type mockQuery struct {
		res []*db_models.UserProgress
		err error
	}

	tests := []struct {
		name    string
		args    args
		mock    mockQuery
		want    []*db_models.UserProgress
		wantErr error
	}{
		{
			name: "[GetUserProgress] Success to get user progress",
			args: args{
				context.TODO(),
				userId,
				courseId,
			},
			mock: mockQuery{
				res: []*db_models.UserProgress{
					{
						UserID:     userId,
						CourseID:   courseId,
						MaterialID: materialId,
						Score:      100,
					},
				},
				err: nil,
			},
			want: []*db_models.UserProgress{
				{
					UserID:     userId,
					CourseID:   courseId,
					MaterialID: materialId,
					Score:      100,
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

			selectQuery := regexp.QuoteMeta(`SELECT user_id, course_id, material_id, score FROM user_progress`)

			if tt.mock.res != nil{
				rows := sqlmock.NewRows([]string{
					"user_id",
					"course_id",
					"material_id",
					"score",
				})

				for _, row := range tt.mock.res {
					rows.AddRow(row.UserID, row.CourseID, row.MaterialID, row.Score)
				}

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, err := r.GetUserProgress(tt.args.ctx, sqlxDB, tt.args.userId, tt.args.courseId)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_CheckIsProgressLogged(t *testing.T) {

	type args struct {
		ctx        context.Context
		userId     string
		materialId string
	}

	type mockQuery struct {
		data []*db_models.UserProgress
		err  error
	}

	tests := []struct {
		name    string
		args    args
		mock 	  mockQuery
		want    bool
		wantErr error
	}{
		{
			name: "[CheckIsProgressLogged] Success to check that progress is logged",
			args: args {
				context.TODO(),
				userId,
				materialId,
			},
			mock: mockQuery {
				data: []*db_models.UserProgress {
					{
						UserID:     userId,
						CourseID:   courseId,
						MaterialID: materialId,
						Score:      100,
					},
				},
				err: nil,
			},
			want: true,
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

			selectQuery := regexp.QuoteMeta(`SELECT user_id, course_id, material_id, score FROM user_progress`)

			if tt.mock.data != nil{
				rows := sqlmock.NewRows([]string{
					"user_id",
					"course_id",
					"material_id",
					"score",
				})

				for _, row := range tt.mock.data {
					rows.AddRow(row.UserID, row.CourseID, row.MaterialID, row.Score)
				}

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, err := r.CheckIsProgressLogged(tt.args.ctx, sqlxDB, tt.args.userId, tt.args.materialId)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_InsertCourseData(t *testing.T) {

	type args struct {
		ctx    context.Context
		input *db_models.Course
	}

	type mockQuery struct {
		data *db_models.Course
		err   error
	}

	tests := []struct {
		name      string
		args      args
		mockQuery mockQuery
		wantErr   error
	}{
		{
			name: "[InsertCourseData] Success to insert course data",
			args: args{
				context.TODO(),
				&db_models.Course{
					ID:          courseId,
					CourseName:  "",
					Description: "",
					Thumbnail:   "",
					Creator:     creatorId,
				},
			},
			mockQuery: mockQuery{
				data: &db_models.Course{
					ID:          courseId,
					CourseName:  "",
					Description: "",
					Thumbnail:   "",
					Creator:     creatorId,
				},
				err: nil,
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

			execQuery := regexp.QuoteMeta(`INSERT INTO Course (id,course_name,description,thumbnail,creator) VALUES (?,?,?,?,?)`)

			mock.ExpectExec(execQuery).WillReturnResult(sqlmock.NewResult(1, 1))

			r := course_repository.NewRepository()
			err = r.InsertCourseData(tt.args.ctx, sqlxDB,tt.args.input)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_InsertCourse(t *testing.T) {
	type args struct {
		ctx    context.Context
		input *models.CourseCreation
	}

	type mockQuery struct {
		err   error
	}

	tests := []struct {
		name    string
		args    args
		mock    mockQuery
		wantErr error
	}{
		{
			name: "[InsertCourse] Success to insert course",
			args: args{
				context.TODO(),
				&models.CourseCreation{
					Course: models.Course{
						ID:          courseId,
						CourseName:  "",
						Description: "",
						Thumbnail:   "",
						Creator:     creatorId,
					},
					Sections: []*models.SectionCreation{
						{
							ID: sectionId,
							Name: "",
							Subsections: []*models.MaterialCreation{
								{
									ID:          materialId,
									Name:        "",
									Type:        "",
									Content:     "",
									ContentText: "",
								},
							},
						},
					},
				},
			},
			mock: mockQuery{
				err: nil,
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

			execCourseQuery := regexp.QuoteMeta(`INSERT INTO Course (id,course_name,description,thumbnail,creator) VALUES (?,?,?,?,?)`)

			execMaterialQuery := regexp.QuoteMeta(`INSERT INTO course_material (id,course_id,name,type,section_id,content,content_text) VALUES (?,?,?,?,?,?,?)`)

			mock.ExpectExec(execCourseQuery).WillReturnResult(sqlmock.NewResult(1, 1))
			
			mock.ExpectExec(execMaterialQuery).WillReturnResult(sqlmock.NewResult(1, 1))

			r := course_repository.NewRepository()
			err = r.InsertCourse(tt.args.ctx, sqlxDB,tt.args.input)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_GetMaterialByID(t *testing.T) {

	type args struct {
		ctx context.Context
		id  string
	}

	type mockQuery struct {
		res *db_models.Material
		err  error
	}

	tests := []struct {
		name     string
		args     args
		mock     mockQuery
		want 	  *db_models.Material
		wantErr  error
	}{
		{
			name: "[GetMaterialByID] Success to get material by ID.",
			args: args{
				context.TODO(),
				materialId,
			},
			mock: mockQuery{
				res: &db_models.Material{
					ID:          materialId,
					CourseID:    courseId,
					Name:        "",
					Type:        "",
					SectionID:   sectionId,
					Content:     "",
					ContentText: "",
				},
				err: nil,
			},
			want: &db_models.Material{
				ID:          materialId,
				CourseID:    courseId,
				Name:        "",
				Type:        "",
				SectionID:   sectionId,
				Content:     "",
				ContentText: "",
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

			selectQuery := regexp.QuoteMeta(`SELECT id, course_id, name, type, section_id, content, content_text FROM course_material`)

			if tt.mock.res != nil{
				rows := sqlmock.NewRows([]string{
					"id",
					"course_id",
					"name",
					"type",
					"section_id",
					"content",
					"content_text",
				})

				rows.AddRow(
					tt.mock.res.ID,
					tt.mock.res.CourseID,
					tt.mock.res.Name,
					tt.mock.res.Type,
					tt.mock.res.SectionID,
					tt.mock.res.Content,
					tt.mock.res.ContentText,
				)

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, err := r.GetMaterialByID(tt.args.ctx, sqlxDB, tt.args.id)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_InsertCourseMaterial(t *testing.T) {

	type args struct {
		ctx    context.Context
		input *db_models.Material
	}

	type mockQuery struct {
		data *db_models.Material
		err   error
	}

	tests := []struct {
		name      string
		args      args
		mockQuery mockQuery
		wantErr   error
	}{
		{
			name: "[InsertCourseMaterial] Success to insert course material.",
			args: args {
				context.TODO(),
				&db_models.Material{
					ID:          materialId,
					CourseID:    courseId,
					Name:        "",
					Type:        "",
					SectionID:   sectionId,
					Content:     "",
					ContentText: "",
				},
			},
			mockQuery: mockQuery{
				data: &db_models.Material{
					ID:          materialId,
					CourseID:    courseId,
					Name:        "",
					Type:        "",
					SectionID:   sectionId,
					Content:     "",
					ContentText: "",
				},
				err: nil,
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

			execQuery := regexp.QuoteMeta(`INSERT INTO course_material (id,course_id,name,type,section_id,content,content_text) VALUES (?,?,?,?,?,?,?)`)

			mock.ExpectExec(execQuery).WillReturnResult(sqlmock.NewResult(1, 1))

			r := course_repository.NewRepository()
			err = r.InsertCourseMaterial(tt.args.ctx, sqlxDB,tt.args.input)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseRepository_GetCourseByCreatorID(t *testing.T) {

	type args struct {
		ctx       context.Context
		creatorId string
	}

	type mockQuery struct {
		res []*db_models.Course
		err error
	}

	tests := []struct {
		name    string
		args    args
		mock    mockQuery
		want    []*db_models.Course
		wantErr error
	}{
		{
			name: "[GetCourseByCreatorID] Success to get course by creator ID.",
			args: args{
				context.TODO(),
				creatorId,
			},
			mock: mockQuery{
				res: []*db_models.Course {
					{
						ID:          courseId,
						CourseName:  "",
						Description: "",
						Thumbnail:   "",
						Creator:     creatorId,
					},
				},
				err: nil,
			},
			want: []*db_models.Course {
				{
					ID:          courseId,
					CourseName:  "",
					Description: "",
					Thumbnail:   "",
					Creator:     creatorId,
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

			selectQuery := regexp.QuoteMeta(`SELECT id, course_name, description, thumbnail, creator FROM Course WHERE creator = ?`)

			if tt.mock.res != nil{
				rows := sqlmock.NewRows([]string{
					"id",
					"course_name",
					"description",
					"thumbnail",
					"creator",
				})

				for _, row := range tt.mock.res {
					rows.AddRow(row.ID, row.CourseName, row.Description, row.Thumbnail, row.Creator)
				}

				mock.ExpectQuery(selectQuery).WillReturnRows(rows)
				mock.ExpectCommit()
			}

			r := course_repository.NewRepository()
			got, err := r.GetCourseByCreatorID(tt.args.ctx, sqlxDB, tt.args.creatorId)
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}
