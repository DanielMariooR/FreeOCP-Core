package course_repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	er "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/error"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	db_models "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/pagination"
)

type courseRepository struct{}

func NewRepository() CourseRepository {
	return &courseRepository{}
}

func (repo *courseRepository) GetTableName() string {
	return "Course"
}

func (repo *courseRepository) querySelectCourse() sq.SelectBuilder {
	builder := sq.Select(
		"id",
		"course_name",
		"description",
		"thumbnail",
		"creator",
	).From(repo.GetTableName())

	return builder
}

func (repo *courseRepository) querySelectCourseSyllabus() sq.SelectBuilder {
	builder := sq.Select(
		"id",
		"course_id",
		"name",
		"type",
		"section_id",
	).From("course_material")

	return builder
}

func (repo *courseRepository) querySelectCourseMaterial() sq.SelectBuilder {
	builder := sq.Select(
		"id",
		"course_id",
		"name",
		"type",
		"section_id",
		"content",
		"content_text",
	).From("course_material")

	return builder
}

func (repo *courseRepository) querySelectUserProgress() sq.SelectBuilder {
	builder := sq.Select(
		"user_id",
		"course_id",
		"material_id",
		"score",
	).From("user_progress")

	return builder
}

func (repo *courseRepository) queryCountCourse() sq.SelectBuilder {
	builder := sq.Select(
		"count(id)",
	).From(repo.GetTableName())

	return builder
}

func (repo *courseRepository) queryInsertCourseData() sq.InsertBuilder {
	builder := sq.Insert(repo.GetTableName()).Columns(
		"id",
		"course_name",
		"description",
		"thumbnail",
		"creator",
	)

	return builder
}

func (repo *courseRepository) queryInsertCourseMaterial() sq.InsertBuilder {
	builder := sq.Insert("course_material").Columns(
		"id",
		"course_id",
		"name",
		"type",
		"section_id",
		"content",
		"content_text",
	)

	return builder
}

func (repo *courseRepository) GetCourseById(ctx context.Context, db *sqlx.DB, id string) (*db_models.Course, error) {
	out := new(db_models.Course)
	query, args, err := repo.querySelectCourse().Where(sq.Eq{"id": id}).ToSql()

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

func (repo *courseRepository) GetCompletedCourseByUserID(ctx context.Context, db *sqlx.DB, meta *pagination.Meta, userid string) ([]*db_models.Course, uint64, error) {
	var courses []*db_models.Course
	var count uint64

	selectQuery, args, err := repo.querySelectCourse().
		Join("solved_course on solved_course.course_id = Course.id").
		Where(sq.Eq{"solved_course.user_id": userid}).ToSql()
	if err != nil {
		return courses, count, err
	}

	query := fmt.Sprintf("%s limit %d,%d", selectQuery, (meta.Page-1)*meta.Limit, meta.Limit)
	err = db.SelectContext(ctx, &courses, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return courses, count, nil
		}

		return courses, count, err
	}

	countQuery, args, err := repo.queryCountCourse().
		Join("solved_course on solved_course.course_id = Course.id").
		Where(sq.Eq{"solved_course.user_id": userid}).ToSql()
	if err != nil {
		return courses, count, err
	}

	err = db.GetContext(ctx, &count, countQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return courses, count, nil
		}

		return courses, count, err
	}

	return courses, count, nil
}

func (repo *courseRepository) GetOnProgressCourseByUserID(ctx context.Context, db *sqlx.DB, meta *pagination.Meta, userid string) ([]*db_models.Course, uint64, error) {
	var courses []*db_models.Course
	var count uint64

	selectQuery, args, err := repo.querySelectCourse().
		Join("on_progress_course on on_progress_course.course_id = Course.id").
		Where(sq.Eq{"on_progress_course.user_id": userid}).ToSql()
	if err != nil {
		return courses, count, err
	}

	query := fmt.Sprintf("%s limit %d,%d", selectQuery, (meta.Page-1)*meta.Limit, meta.Limit)
	err = db.SelectContext(ctx, &courses, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return courses, count, nil
		}

		return courses, count, err
	}

	countQuery, args, err := repo.queryCountCourse().
		Join("on_progress_course on on_progress_course.course_id = Course.id").
		Where(sq.Eq{"on_progress_course.user_id": userid}).ToSql()
	if err != nil {
		return courses, count, err
	}

	err = db.GetContext(ctx, &count, countQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return courses, count, nil
		}

		return courses, count, err
	}

	return courses, count, nil
}

func (repo *courseRepository) GetCourseList(ctx context.Context, db *sqlx.DB, meta *pagination.Meta) ([]*db_models.Course, uint64, error) {
	var courses []*db_models.Course
	var count uint64

	selectQuery, args, err := repo.querySelectCourse().ToSql()
	if err != nil {
		return courses, count, err
	}

	query := fmt.Sprintf("%s limit %d,%d", selectQuery, (meta.Page-1)*meta.Limit, meta.Limit)
	err = db.SelectContext(ctx, &courses, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return courses, count, nil
		}

		return courses, count, err
	}

	countQuery, args, err := repo.queryCountCourse().ToSql()
	if err != nil {
		return courses, count, err
	}

	err = db.GetContext(ctx, &count, countQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return courses, count, nil
		}

		return courses, count, err
	}

	return courses, count, nil
}

func (repo *courseRepository) GetCourseSyllabusByCourseID(ctx context.Context, db *sqlx.DB, courseId string) ([]*db_models.Syllabus, error) {
	var syllabus []*db_models.Syllabus

	query, args, err := repo.querySelectCourseSyllabus().
		Where(sq.Eq{"course_id": courseId}).ToSql()
	if err != nil {
		return syllabus, err
	}

	err = db.SelectContext(ctx, &syllabus, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return syllabus, nil
		}
		return syllabus, err
	}

	return syllabus, nil
}

func (repo *courseRepository) GetCourseMaterialByCourseIDAndSectionID(ctx context.Context, db *sqlx.DB, courseId string, sectionId string) ([]*db_models.Material, error) {
	var material []*db_models.Material

	query, args, err := repo.querySelectCourseMaterial().
		Where(sq.Eq{"course_id": courseId}).
		Where(sq.Eq{"section_id": sectionId}).
		ToSql()
	if err != nil {
		return material, err
	}

	err = db.SelectContext(ctx, &material, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return material, nil
		}
		return material, err
	}

	return material, nil
}

func (repo *courseRepository) IsUserEnrolledToCourse(ctx context.Context, db *sqlx.DB, userid string, courseid string) (bool, error) {
	var count uint64

	countQuery, args, err := sq.Select("count(*)").
		From("on_progress_course").
		Where(sq.Eq{"user_id": userid}).
		Where(sq.Eq{"course_id": courseid}).ToSql()

	if err != nil {
		return false, err
	}

	err = db.GetContext(ctx, &count, countQuery, args...)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (repo *courseRepository) InsertEnrollment(ctx context.Context, db *sqlx.DB, values *models.EnrollInput) error {
	query, args, err := sq.Insert("on_progress_course").
		Columns("user_id", "course_id").
		Values(values.UserID, values.CourseID).ToSql()

	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *courseRepository) GetCourseIDByMaterialID(ctx context.Context, db *sqlx.DB, materialID string) (string, error) {
	out := new(db_models.Syllabus)
	query, args, err := repo.querySelectCourseSyllabus().Where(sq.Eq{"id": materialID}).ToSql()
	if err != nil {
		return "", err
	}

	err = db.GetContext(ctx, out, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", er.NewError(fmt.Errorf("%s", "Material is not in course syllabus"), http.StatusBadRequest, nil)
		}
		return "", err
	}

	return out.CourseID, nil
}

func (repo *courseRepository) StoreUserProgress(ctx context.Context, db *sqlx.DB, materialID, courseID, userID string, score int) error {
	query, args, err := sq.Insert("user_progress").
		Columns("user_id", "course_id", "material_id", "score").
		Values(userID, courseID, materialID, score).ToSql()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *courseRepository) GetUserProgress(ctx context.Context, db *sqlx.DB, userID, courseID string) ([]*db_models.UserProgress, error) {
	var userProgress []*db_models.UserProgress

	query, args, err := repo.querySelectUserProgress().Where(sq.Eq{"user_id": userID, "course_id": courseID}).ToSql()
	if err != nil {
		return userProgress, err
	}

	err = db.SelectContext(ctx, &userProgress, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return userProgress, nil
		}
		return userProgress, err
	}

	return userProgress, nil
}

func (repo *courseRepository) CheckIsProgressLogged(ctx context.Context, db *sqlx.DB, userID, materialID string) (bool, error) {
	var userProgress []*db_models.UserProgress

	query, args, err := repo.querySelectUserProgress().Where(sq.Eq{"user_id": userID, "material_id": materialID}).ToSql()
	if err != nil {
		return false, err
	}

	err = db.SelectContext(ctx, &userProgress, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if len(userProgress) == 0 {
		return false, nil
	}

	return true, nil
}

func (repo *courseRepository) InsertCourseData(ctx context.Context, db *sqlx.DB, values *db_models.Course) error {
	query, args, err := repo.queryInsertCourseData().Values(
		values.ID,
		values.CourseName,
		values.Description,
		values.Thumbnail,
		values.Creator,
	).ToSql()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (repo *courseRepository) InsertCourseMaterialMany(ctx context.Context, db *sqlx.DB, values *[]db_models.Material) error {
	queryBuilder := repo.queryInsertCourseMaterial()

	for _, material := range *values {
		queryBuilder = queryBuilder.Values(
			material.ID,
			material.CourseID,
			material.Name,
			material.Type,
			material.SectionID,
			material.Content,
			material.ContentText,
		)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (repo *courseRepository) InsertCourse(ctx context.Context, db *sqlx.DB, input *models.CourseCreation) error {
	courseData := &db_models.Course{
		ID:          input.Course.ID,
		CourseName:  input.Course.CourseName,
		Description: input.Course.Description,
		Thumbnail:   input.Course.Thumbnail,
		Creator:     input.Course.Creator,
	}
	err := repo.InsertCourseData(ctx, db, courseData)
	if err != nil {
		return err
	}

	materialData := []db_models.Material{}

	for _, sect := range input.Sections {
		materialData = append(materialData, db_models.Material{
			ID:          sect.ID,
			CourseID:    courseData.ID,
			Name:        sect.Name,
			Type:        "section",
			SectionID:   "",
			Content:     "",
			ContentText: "",
		})
		for _, mat := range sect.Subsections {
			materialData = append(materialData, db_models.Material{
				ID:          mat.ID,
				CourseID:    courseData.ID,
				Name:        mat.Name,
				Type:        mat.Type,
				SectionID:   sect.ID,
				Content:     mat.Content,
				ContentText: mat.ContentText,
			})
		}
	}
	err = repo.InsertCourseMaterialMany(ctx, db, &materialData)
	if err != nil {
		return err
	}

	return nil
}

func (repo *courseRepository) GetMaterialByID(ctx context.Context, db *sqlx.DB, id string) (*db_models.Material, error) {
	out := new(db_models.Material)
	query, args, err := repo.querySelectCourseMaterial().Where(sq.Eq{"id": id}).ToSql()

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

func (repo *courseRepository) InsertCourseMaterial(ctx context.Context, db *sqlx.DB, material *db_models.Material) error {
	query, args, err := repo.queryInsertCourseMaterial().Values(
		material.ID,
		material.CourseID,
		material.Name,
		material.Type,
		material.SectionID,
		material.Content,
		material.ContentText,
	).ToSql()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (repo *courseRepository) GetCourseByCreatorID(ctx context.Context, db *sqlx.DB, creatorId string) ([]*db_models.Course, error) {
	var courses []*db_models.Course

	query, args, err := repo.querySelectCourse().Where(sq.Eq{"creator": creatorId}).ToSql()
	if err != nil {
		return nil, err
	}

	err = db.SelectContext(ctx, &courses, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return courses, nil
		}
		return courses, err
	}

	return courses, nil
}
