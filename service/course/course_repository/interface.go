package course_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/pagination"
)

type CourseRepository interface {
	GetTableName() string
	GetCourseById(ctx context.Context, db *sqlx.DB, id string) (*db_models.Course, error)
	GetCourseList(ctx context.Context, db *sqlx.DB, meta *pagination.Meta) ([]*db_models.Course, uint64, error)
	GetCompletedCourseByUserID(ctx context.Context, db *sqlx.DB, meta *pagination.Meta, userid string) ([]*db_models.Course, uint64, error)
	GetOnProgressCourseByUserID(ctx context.Context, db *sqlx.DB, meta *pagination.Meta, userid string) ([]*db_models.Course, uint64, error)
	GetCourseSyllabusByCourseID(ctx context.Context, db *sqlx.DB, courseId string) ([]*db_models.Syllabus, error)
	IsUserEnrolledToCourse(ctx context.Context, db *sqlx.DB, userid string, courseid string) (bool, error)
	InsertEnrollment(ctx context.Context, db *sqlx.DB, values *models.EnrollInput) error
	GetCourseMaterialByCourseIDAndSectionID(ctx context.Context, db *sqlx.DB, courseId string, sectionId string) ([]*db_models.Material, error)
	GetCourseIDByMaterialID(ctx context.Context, db *sqlx.DB, materialID string) (string, error)
	StoreUserProgress(ctx context.Context, db *sqlx.DB, materialID, courseID, userID string, score int) error
	GetUserProgress(ctx context.Context, db *sqlx.DB, userID, courseID string) ([]*db_models.UserProgress, error)
	CheckIsProgressLogged(ctx context.Context, db *sqlx.DB, userID, materialID string) (bool, error)
	InsertCourse(ctx context.Context, db *sqlx.DB, course *models.CourseCreation) error
	GetMaterialByID(ctx context.Context, db *sqlx.DB, id string) (*db_models.Material, error)
	InsertCourseData(ctx context.Context, db *sqlx.DB, course *db_models.Course) error
	InsertCourseMaterial(ctx context.Context, db *sqlx.DB, course *db_models.Material) error
	GetCourseByCreatorID(ctx context.Context, db *sqlx.DB, creatorId string) ([]*db_models.Course, error)
}
