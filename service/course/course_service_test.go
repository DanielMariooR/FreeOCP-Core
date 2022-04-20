package course_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	er "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/error"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/mocks"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/pagination"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/course"
)


var (
	courseId = uuid.New().String()
	courseId2 = uuid.New().String()
	courseId3 = uuid.New().String()
	
	syllabusId = uuid.New().String()
	syllabusId2 = uuid.New().String()
	syllabusId3 = uuid.New().String()

	materialId = uuid.New().String()

	userId = uuid.New().String()

	creatorId = uuid.New().String()
)

func TestCourseService_GetCourseDetail(t *testing.T) {
	type mockRepo struct {
		res *db_models.Course
		err error
	}

	type args struct {
		ctx context.Context
		id  string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.Course
		wantErr  error
	}{
		{
			name: "[GetCourseDetail] Success to get course detail",
			args: args{
				context.TODO(),
				courseId,
			},
			mock: mockRepo{
				res: &db_models.Course{
					ID: courseId,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: creatorId,
				},
				err: nil,
			},
			want: &models.Course{
				ID: courseId,
				CourseName: "",
				Description: "",
				Thumbnail: "",
				Creator: creatorId,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("GetCourseById",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.res, tt.mock.err)

			got, err := svc.GetCourseDetail(tt.args.ctx, tt.args.id)
			
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}
}

func TestCourseService_GetCompletedCourse(t *testing.T) {
	type mockRepo struct {
		res   []*db_models.Course
		count uint64
		err   error
	}

	type args struct {
		ctx     context.Context
		meta   *pagination.Meta
		userId  string
	}

	tests := []struct {
		name        string
		args        args
		mock        mockRepo
		wantCourses []*models.Course
		wantCount   uint64
		wantErr     error
	}{
		{
			name: "[GetCompletedCourse] Success to get list of completed courses",
			args: args{
				context.TODO(),
				&pagination.Meta{
					Limit: 3,
					Page: 3,
				},
				courseId,
			},
			mock: mockRepo{
				res: []*db_models.Course{
					{
						ID: courseId,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: creatorId,
					},
					{
						ID: courseId2,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: creatorId,
					},
					{
						ID: courseId3,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: creatorId,
					},
				},
				count: 3,
				err: nil,
			},
			wantCourses: []*models.Course{
				{
					ID: courseId,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: creatorId,
				},
				{
					ID: courseId2,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: creatorId,
				},
				{
					ID: courseId3,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: creatorId,
				},
			},
			wantCount: 3,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("GetCompletedCourseByUserID",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.res,tt.mock.count, tt.mock.err)

			got, cou, err := svc.GetCompeletedCourse(tt.args.ctx, tt.args.meta, tt.args.userId)
			
			assert.Equal(t, tt.wantCourses, got, tt.name)
			assert.Equal(t, tt.wantCount, cou, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseService_GetOnProgressCourse(t *testing.T) {
	type mockRepo struct {
		res   []*db_models.Course
		count uint64
		err   error
	}

	type args struct {
		ctx     context.Context
		meta   *pagination.Meta
		userId  string
	}

	tests := []struct {
		name        string
		args        args
		mock        mockRepo
		wantCourses []*models.Course
		wantCount   uint64
		wantErr     error
	}{
		{
			name: "[GetOnProgressCourse] Success to get list of on progress courses",
			args: args{
				context.TODO(),
				&pagination.Meta{
					Limit: 3,
					Page: 3,
				},
				courseId,
			},
			mock: mockRepo{
				res: []*db_models.Course{
					{
						ID: courseId,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: creatorId,
					},
					{
						ID: courseId2,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: creatorId,
					},
					{
						ID: courseId3,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: creatorId,
					},
				},
				count: 3,
				err: nil,
			},
			wantCourses: []*models.Course{
				{
					ID: courseId,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: creatorId,
				},
				{
					ID: courseId2,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: creatorId,
				},
				{
					ID: courseId3,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: creatorId,
				},
			},
			wantCount: 3,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("GetOnProgressCourseByUserID",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.res,tt.mock.count, tt.mock.err)

			got, cou, err := svc.GetOnProgressCourse(tt.args.ctx, tt.args.meta, tt.args.userId)
			
			assert.Equal(t, tt.wantCourses, got, tt.name)
			assert.Equal(t, tt.wantCount, cou, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseService_GetCoursePagination(t *testing.T) {
	type mockRepo struct {
		res   []*db_models.Course
		count uint64
		err   error
	}

	type args struct {
		ctx     context.Context
		meta   *pagination.Meta
	}

	tests := []struct {
		name        string
		args        args
		mock        mockRepo
		wantCourses []*models.Course
		wantCount   uint64
		wantErr     error
	}{
		{
			name: "[GetCoursePagination] Success to get list of courses",
			args: args{
				context.TODO(),
				&pagination.Meta{
					Limit: 3,
					Page: 3,
				},
			},
			mock: mockRepo{
				res: []*db_models.Course{
					{
						ID: courseId,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: "",
					},
					{
						ID: courseId2,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: "",
					},
					{
						ID: courseId3,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: "",
					},
				},
				count: 3,
				err: nil,
			},
			wantCourses: []*models.Course{
				{
					ID: courseId,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: "",
				},
				{
					ID: courseId2,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: "",
				},
				{
					ID: courseId3,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: "",
				},
			},
			wantCount: 3,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("GetCourseList",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.res,tt.mock.count, tt.mock.err)

			got, cou, err := svc.GetCoursePagination(tt.args.ctx, tt.args.meta)
			
			assert.Equal(t, tt.wantCourses, got, tt.name)
			assert.Equal(t, tt.wantCount, cou, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseService_GetCourseSyllabus(t *testing.T) {
	type mockRepo struct {
		res []*db_models.Syllabus
		err error
	}

	type args struct {
		ctx context.Context
		courseId  string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.SyllabusResponse
		wantErr  error
	}{
		{
			name: "[GetCourseSyllabus] Success to get course syllabus",
			args: args{
				context.TODO(),
				courseId,
			},
			mock: mockRepo{
				res: []*db_models.Syllabus{
					{
						ID: syllabusId,
						CourseID: courseId,
						Name: "",
						Type: "section",
						SectionID: nil,
					},
					{
						ID: syllabusId2,
						CourseID: courseId,
						Name: "",
						Type: "video",
						SectionID: &syllabusId,
					},
					{
						ID: syllabusId3,
						CourseID: courseId,
						Name: "",
						Type: "assignment",
						SectionID: &syllabusId,
					},
				},
				err: nil,
			},
			want: &models.SyllabusResponse{
				Syllabus: []*models.Section{
					{
						ID: syllabusId,
						Name: "",
						Subsections: []*models.Material{
							{
								ID: syllabusId2,
								Name: "",
								Type: "video",
							},
							{
								ID: syllabusId3,
								Name: "",
								Type: "assignment",
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "[GetCourseSyllabus] Success to get course syllabus with no section",
			args: args{
				context.TODO(),
				courseId,
			},
			mock: mockRepo{
				res: []*db_models.Syllabus{
					{
						ID: syllabusId2,
						CourseID: courseId,
						Name: "",
						Type: "video",
						SectionID: nil,
					},
					{
						ID: syllabusId3,
						CourseID: courseId,
						Name: "",
						Type: "assignment",
						SectionID: nil,
					},
				},
				err: nil,
			},
			want: &models.SyllabusResponse{
				Syllabus: []*models.Section{
					{
						ID: "0000",
						Name: "Syllabus",
						Subsections: []*models.Material{
							{
								ID: syllabusId2,
								Name: "",
								Type: "video",
							},
							{
								ID: syllabusId3,
								Name: "",
								Type: "assignment",
							},
						},
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("GetCourseSyllabusByCourseID",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.res, tt.mock.err)

			got, err := svc.GetCourseSyllabus(tt.args.ctx, tt.args.courseId)
			
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}
}

func TestCourseService_GetCourseMaterial(t *testing.T) {
	type mockRepo struct {
		res  []*db_models.Material
		err  error
	}

	type args struct {
		ctx       context.Context
		courseId  string
		sectionId string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.SectionContentResponse
		wantErr  error
	}{
		{
			name: "[GetCourseMaterial] Success to get course material",
			args: args{
				context.TODO(),
				courseId,
				syllabusId,
			},
			mock: mockRepo{
				res: []*db_models.Material{
					{	
						ID: syllabusId2,
						CourseID: courseId,
						Name: "",
						Type: "video",
						SectionID: syllabusId,
						Content: "",
						ContentText: "",
					},
					{	
						ID: syllabusId3,
						CourseID: courseId,
						Name: "",
						Type: "assignment",
						SectionID: syllabusId,
						Content: "",
						ContentText: "",
					},
				},
				err: nil,
			},
			want: &models.SectionContentResponse{
				ID: syllabusId,
				Subsections: []*models.MaterialContent{
					{
						ID: syllabusId2,
						Name: "",
						Type: "video",
						Content: "",
						ContentText: "",
					},
					{
						ID: syllabusId3,
						Name: "",
						Type: "assignment",
						Content: "",
						ContentText: "",
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("GetCourseMaterialByCourseIDAndSectionID",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.res, tt.mock.err)

			got, err := svc.GetCourseMaterial(tt.args.ctx, tt.args.courseId, tt.args.sectionId)
			
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}
}

func TestCourseService_Enroll(t *testing.T) {
	type mockRepo struct {
		check struct {
			is  bool
			err error
		}
		insert struct {
			err error
		}
	}

	type args struct {
		ctx      context.Context
		userId   string
		courseId string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.EnrollResponse
		wantErr  error
	}{
		{
			name: "[Enroll] Success to enroll to a course",
			args: args{
				context.TODO(),
				userId,
				courseId,
			},
			mock: mockRepo{
				check: struct{is bool; err error}{
					false,
					nil,
				},
				insert: struct{err error}{
					nil,
				},
			},
			want: &models.EnrollResponse{
				Status: "Success",
				Message: "You have successfully enrolled to the course.",
			},
			wantErr: nil,
		},
		{
			name: "[Enroll] Already enrolled to a course",
			args: args{
				context.TODO(),
				userId,
				courseId,
			},
			mock: mockRepo{
				check: struct{is bool; err error}{
					true,
					nil,
				},
				insert: struct{err error}{
					nil,
				},
			},
			want: nil,
			wantErr: er.NewError(fmt.Errorf("%s", "You are already enrolled to the course"), http.StatusBadRequest, nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("IsUserEnrolledToCourse",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.check.is, tt.mock.check.err)

			repoMock.
				On("InsertEnrollment", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.insert.err)

			got, err := svc.Enroll(tt.args.ctx, tt.args.userId, tt.args.courseId)
			
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}
}

func TestCourseService_StoreUserProgress(t *testing.T) {
	type mockRepo struct {
		get struct {
			res string
			err error
		}
		checkEnrolled struct {
			is  bool
			err error
		}
		checkLogged struct {
			is  bool
			err error
		}
		insert struct {
			err error
		}
	}

	type args struct {
		ctx context.Context
		userId     string
		materialId string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.StoreProgressResponse
		wantErr  error
	}{
		{
			name: "[StoreUserProgress] Success to store user progress",
			args: args{
				context.TODO(),
				userId,
				materialId,
			},
			mock: mockRepo{
				get: struct{res string; err error}{
					courseId,
					nil,
				},
				checkEnrolled: struct{is bool; err error}{
					true,
					nil,
				},
				checkLogged: struct{is bool; err error}{
					false,
					nil,
				},
				insert: struct{err error}{
					nil,
				},
			},
			want: &models.StoreProgressResponse{
				Status:  "Success",
				Message: "User progress updated successfully",
			},
			wantErr: nil,
		},
		{
			name: "[StoreUserProgress] User is not enrolled to the course",
			args: args{
				context.TODO(),
				userId,
				materialId,
			},
			mock: mockRepo{
				get: struct{res string; err error}{
					courseId,
					nil,
				},
				checkEnrolled: struct{is bool; err error}{
					false,
					nil,
				},
				checkLogged: struct{is bool; err error}{
					false,
					nil,
				},
				insert: struct{err error}{
					nil,
				},
			},
			want: nil,
			wantErr: er.NewError(fmt.Errorf("%s", "You are not enrolled to the course"), http.StatusBadRequest, nil),
		},
		{
			name: "[StoreUserProgress] User progress is already logged",
			args: args{
				context.TODO(),
				userId,
				materialId,
			},
			mock: mockRepo{
				get: struct{res string; err error}{
					courseId,
					nil,
				},
				checkEnrolled: struct{is bool; err error}{
					true,
					nil,
				},
				checkLogged: struct{is bool; err error}{
					true,
					nil,
				},
				insert: struct{err error}{
					nil,
				},
			},
			want: nil,
			wantErr: er.NewError(fmt.Errorf("%s", "Your progress has already been recorded"), http.StatusBadRequest, nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("GetCourseIDByMaterialID",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.get.res, tt.mock.get.err)

			repoMock.
				On("IsUserEnrolledToCourse",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.checkEnrolled.is, tt.mock.checkEnrolled.err)
			
			repoMock.
				On("CheckIsProgressLogged",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.checkLogged.is, tt.mock.checkLogged.err)

			repoMock.
				On("StoreUserProgress",mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.insert.err)
			
			got, err := svc.StoreUserProgress(tt.args.ctx, tt.args.userId, tt.args.materialId)
			
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseService_ComputeUserProgress(t *testing.T) {
	type mockRepo struct {
		check struct {
			is  bool
			err error
		}
		getUser struct {
			res []*db_models.UserProgress
			err error
		}
		getSyllabus struct {
			res []*db_models.Syllabus
			err error
		}
	}

	type args struct {
		ctx      context.Context
		userId   string
		courseId string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.GetProgressPercentageResponse
		wantErr  error
	}{
		{
			name: "[ComputeUserProgress] Success to compute user progress",
			args: args{
				context.TODO(),
				userId,
				courseId,
			},
			mock: mockRepo{
				check: struct{is bool; err error}{
					true,
					nil,
				},
				getUser: struct{res []*db_models.UserProgress; err error}{
					[]*db_models.UserProgress{
						{
							UserID: userId,
							CourseID: courseId,
							MaterialID: syllabusId2,
							Score: 100,
						},
						{
							UserID: userId,
							CourseID: courseId,
							MaterialID: syllabusId3,
							Score: 100,
						},
					},
					nil,
				},
				getSyllabus: struct{res []*db_models.Syllabus; err error}{
					[]*db_models.Syllabus{
						{
							ID: syllabusId,
							CourseID: courseId,
							Name: "",
							Type: "section",
							SectionID: nil,
						},
						{
							ID: syllabusId2,
							CourseID: courseId,
							Name: "",
							Type: "video",
							SectionID: &syllabusId,
						},
						{
							ID: syllabusId3,
							CourseID: courseId,
							Name: "",
							Type: "assignment",
							SectionID: &syllabusId,
						},
					},
					nil,
				},
			},
			want: &models.GetProgressPercentageResponse{
				Percentage: 100,
			},
			wantErr: nil,
		},
		{
			name: "[ComputeUserProgress] User is not enrolled to the course",
			args: args{
				context.TODO(),
				userId,
				courseId,
			},
			mock: mockRepo{
				check: struct{is bool; err error}{
					false,
					nil,
				},
				getUser: struct{res []*db_models.UserProgress; err error}{
					nil,
					nil,
				},
				getSyllabus: struct{res []*db_models.Syllabus; err error}{
					nil,
					nil,
				},
			},
			want: nil,
			wantErr: er.NewError(fmt.Errorf("%s", "You are not enrolled to the course"), http.StatusBadRequest, nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("IsUserEnrolledToCourse",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.check.is, tt.mock.check.err)

			if tt.mock.getUser.res != nil {
				repoMock.
					On("GetUserProgress",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(tt.mock.getUser.res, tt.mock.getUser.err)
			}

			if tt.mock.getSyllabus.res != nil {
				repoMock.
					On("GetCourseSyllabusByCourseID",mock.Anything, mock.Anything, mock.Anything).
					Return(tt.mock.getSyllabus.res, tt.mock.getSyllabus.err)
			}
			
			got, err := svc.ComputeUserProgress(tt.args.ctx, tt.args.userId, tt.args.courseId)
			
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseService_GetUserProgress(t *testing.T) {
	type mockRepo struct {
		check struct {
			is  bool
			err error
		}
		get struct {
			res []*db_models.UserProgress
			err error
		}
	}

	type args struct {
		ctx context.Context
		userId  string
		courseId string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.GetProgressResponse
		wantErr  error
	}{
		{
			name: "[GetUserProgress] Success to get user progress",
			args: args{
				context.TODO(),
				userId,
				courseId,
			},
			mock: mockRepo{
				check: struct{is bool; err error}{
					true,
					nil,
				},
				get: struct{res []*db_models.UserProgress; err error}{
					[]*db_models.UserProgress{
						{
							UserID: userId,
							CourseID: courseId,
							MaterialID: syllabusId2,
							Score: 100,
						},
						{
							UserID: userId,
							CourseID: courseId,
							MaterialID: syllabusId3,
							Score: 100,
						},
					},
					nil,
				},
			},
			want: &models.GetProgressResponse{
				Progress: []*models.UserProgress{
					{
						UserID: userId,
						CourseID: courseId,
						MaterialID: syllabusId2,
						Score: 100,
					},
					{
						UserID: userId,
						CourseID: courseId,
						MaterialID: syllabusId3,
						Score: 100,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "[GetUserProgress] User is not enrolled to the course",
			args: args{
				context.TODO(),
				userId,
				courseId,
			},
			mock: mockRepo{
				check: struct{is bool; err error}{
					false,
					nil,
				},
				get: struct{res []*db_models.UserProgress; err error}{
					nil,
					nil,
				},
			},
			want: nil,
			wantErr: er.NewError(fmt.Errorf("%s", "You are not enrolled to the course"), http.StatusBadRequest, nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("IsUserEnrolledToCourse",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.check.is, tt.mock.check.err)
			
			if tt.mock.get.res != nil {
				repoMock.
					On("GetUserProgress",mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(tt.mock.get.res, tt.mock.get.err)
			}

			got, err := svc.GetUserProgress(tt.args.ctx, tt.args.userId, tt.args.courseId)
			
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}
}

func TestCourseService_CreateNewCourse(t *testing.T) {
	type mockRepo struct {
		err error
	}

	type args struct {
		ctx context.Context
		course *models.CourseCreation
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.CourseCreationResponse
		wantErr  error
	}{
		{
			name: "[CreateNewCourse] Success to create a new course",
			args: args{
				context.TODO(),
				&models.CourseCreation{
					Course: models.Course{},
					Sections: []*models.SectionCreation{
						{
							Subsections: []*models.MaterialCreation{
								{},
							},
						},
					},
				},
			},
			mock: mockRepo{
				err: nil,
			},
			want: &models.CourseCreationResponse{
				Status:  "Success",
				Message: "Course Created Succesfully",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("InsertCourse",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.err)

			got, err := svc.CreateNewCourse(tt.args.ctx, tt.args.course)
			
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)

		})
	}
}

// TODO: TestCourseService_UploadImage

// TODO: TestCourseService_CreateCourseDesc
func TestCourseService_CreateCourseDesc(t *testing.T) {
	type mockRepo struct {
		err error
	}

	type args struct {
		ctx        context.Context
		course    *models.CourseDescriptionInput
		creatorId  string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.CourseCreationResponse
		wantErr  error
	}{
		{
			name: "[CreateCourseDesc] Success to create course with description",
			args: args{
				context.TODO(),
				&models.CourseDescriptionInput{
					CourseName: "",
					Description: "",
					Thumbnail: "",
				},
				creatorId,
			},
			mock: mockRepo{
				err: nil,
			},
			want: &models.CourseCreationResponse{
				Status:  "Success",
				Message: "Course Description Created Succesfully",
				Id: "",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("InsertCourseData",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.err)

			got, err := svc.CreateCourseDesc(tt.args.ctx, tt.args.course, tt.args.creatorId)
			
			assert.Equal(t, tt.want.Status, got.Status, tt.name)
			assert.Equal(t, tt.want.Message, got.Message, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseService_CreateCourseSection(t *testing.T) {
	type mockRepo struct {
		err error
	}

	type args struct {
		ctx        context.Context
		input     *models.CourseSectionInput
		creatorId  string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.CourseCreationResponse
		wantErr  error
	}{
		{
			name: "[CreateCourseSection] Success to create course with section",
			args: args{
				context.TODO(),
				&models.CourseSectionInput{
					Name: "",
					CourseID: "",
				},
				creatorId,
			},
			mock: mockRepo{
				err: nil,
			},
			want: &models.CourseCreationResponse{
				Status:  "Success",
				Message: "Section Created Succesfully",
				Id: "",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("InsertCourseMaterial",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.err)

			got, err := svc.CreateCourseSection(tt.args.ctx, tt.args.input, tt.args.creatorId)
			
			assert.Equal(t, tt.want.Status, got.Status, tt.name)
			assert.Equal(t, tt.want.Message, got.Message, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseService_CreateCourseMaterial(t *testing.T) {
	type mockRepo struct {
		getId struct {
			id  string
			err error
		}
		getMat struct {
			res *db_models.Material
			err  error
		}
		insert struct {
			err error
		}
	}

	type args struct {
		ctx        context.Context
		input     *models.CourseMaterialInput
		creatorId  string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want    *models.CourseCreationResponse
		wantErr  error
	}{
		{
			name: "[CreateCourseMaterial] Success to create course with material",
			args: args{
				context.TODO(),
				&models.CourseMaterialInput{
					Name: "",
					Type: "",
					Content: "",
					ContentText: "",
					SectionID: "",
				},
				creatorId,
			},
			mock: mockRepo{
				getId: struct{id string; err error}{
					syllabusId,
					nil,
				},
				getMat: struct{res *db_models.Material; err error}{
					nil,
					nil,
				},
				insert: struct{err error}{
					nil,
				},
			},
			want: &models.CourseCreationResponse{
				Status:  "Success",
				Message: "Section Created Succesfully",
				Id: "",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("GetCourseIDByMaterialID",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.getId.id, tt.mock.getId.err)

			repoMock.
				On("GetMaterialByID",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.getMat.res, tt.mock.getMat.err)

			repoMock.
				On("InsertCourseMaterial",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.insert.err)

			got, err := svc.CreateCourseMaterial(tt.args.ctx, tt.args.input, tt.args.creatorId)
			
			assert.Equal(t, tt.want.Status, got.Status, tt.name)
			assert.Equal(t, tt.want.Message, got.Message, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestCourseService_GetCourseByCreatorID(t *testing.T) {
	type mockRepo struct {
		res []*db_models.Course
		err error
	}

	type args struct {
		ctx context.Context
		creatorId  string
	}

	tests := []struct {
		name     string
		args     args
		mock     mockRepo
		want     []*models.Course
		wantErr  error
	}{
		{
			name: "[GetCourseByCreatorID] Success to get course by creator ID",
			args: args{
				context.TODO(),
				creatorId,
			},
			mock: mockRepo{
				res: []*db_models.Course{
					{
						ID: courseId,
						CourseName: "",
						Description: "",
						Thumbnail: "",
						Creator: creatorId,
					},
				},
				err: nil,
			},
			want: []*models.Course{
				{
					ID: courseId,
					CourseName: "",
					Description: "",
					Thumbnail: "",
					Creator: creatorId,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")

			repoMock := new(mocks.CourseRepository)
			svc := course.NewService(sqlxDB)
			svc.InjectCourseRepository(repoMock)

			repoMock.
				On("GetCourseByCreatorID",mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mock.res, tt.mock.err)

			got, err := svc.GetCourseByCreatorID(tt.args.ctx, tt.args.creatorId)
			
			assert.Equal(t, tt.want, got, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}
