package course

import (
	"context"
	"net/http"

	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/pagination"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/course/course_repository"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user/user_repository"
)

type CourseService interface {
	InjectCourseRepository(course_repository.CourseRepository) error
	InjectUserRepository(repo user_repository.UserRepository) error
	GetCourseDetail(ctx context.Context, id string) (*models.Course, error)
	GetCompeletedCourse(ctx context.Context, meta *pagination.Meta, userId string) ([]*models.Course, uint64, error)
	GetOnProgressCourse(ctx context.Context, meta *pagination.Meta, userId string) ([]*models.Course, uint64, error)
	GetCoursePagination(ctx context.Context, meta *pagination.Meta) ([]*models.Course, uint64, error)
	GetCourseSyllabus(ctx context.Context, courseId string) (*models.SyllabusResponse, error)
	GetCourseMaterial(ctx context.Context, courseId string, sectionId string) (*models.SectionContentResponse, error)
	Enroll(ctx context.Context, userId string, courseId string) (*models.EnrollResponse, error)
	StoreUserProgress(ctx context.Context, userId, materialId string) (*models.StoreProgressResponse, error)
	ComputeUserProgress(ctx context.Context, userId, courseId string) (*models.GetProgressPercentageResponse, error)
	GetUserProgress(ctx context.Context, userId, courseId string) (*models.GetProgressResponse, error)
	CreateNewCourse(ctx context.Context, course *models.CourseCreation) (*models.CourseCreationResponse, error)
	CreateCourseDesc(ctx context.Context, course *models.CourseDescriptionInput, creatorId string) (*models.CourseCreationResponse, error)
	CreateCourseSection(ctx context.Context, course *models.CourseSectionInput, creatorId string) (*models.CourseCreationResponse, error)
	CreateCourseMaterial(ctx context.Context, course *models.CourseMaterialInput, creatorId string) (*models.CourseCreationResponse, error)
	UploadImage(ctx context.Context, request *http.Request, baseURL string) (*models.UploadImageResponse, error)
	GetCourseByCreatorID(ctx context.Context, creatorId string) ([]*models.Course, error)
}
