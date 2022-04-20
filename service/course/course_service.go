package course

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	er "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/error"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/db"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/pagination"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/course/course_repository"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user/user_repository"
)

type courseService struct {
	db               *sqlx.DB
	courseRepository course_repository.CourseRepository
	userRepository   user_repository.UserRepository
}

func NewService(db *sqlx.DB) CourseService {
	return &courseService{
		db: db,
	}
}

func (serv *courseService) GetCourseDetail(ctx context.Context, id string) (*models.Course, error) {
	course, err := serv.courseRepository.GetCourseById(ctx, serv.db, id)
	if err != nil {
		return nil, err
	}

	var username string
	user, err := serv.userRepository.GetUserById(ctx, serv.db, course.Creator)
	if err != nil {
		return nil, err
	}

	if user == nil {
		username = "anon"
	} else {
		username = user.Username
	}

	resp := models.Course{
		ID:          course.ID,
		CourseName:  course.CourseName,
		Description: course.Description,
		Thumbnail:   course.Thumbnail,
		Creator:     username,
	}

	return &resp, nil
}

func (serv *courseService) GetCompeletedCourse(ctx context.Context, meta *pagination.Meta, userId string) ([]*models.Course, uint64, error) {
	var (
		count   uint64
		courses []*models.Course
	)

	db_courses, count, err := serv.courseRepository.GetCompletedCourseByUserID(ctx, serv.db, meta, userId)
	if err != nil {
		return courses, count, err
	}

	for _, course := range db_courses {

		var username string
		user, err := serv.userRepository.GetUserById(ctx, serv.db, course.Creator)
		if err != nil {
			return courses, count, err
		}

		if user == nil {
			username = "anon"
		} else {
			username = user.Username
		}

		temp := models.Course{
			ID:          course.ID,
			CourseName:  course.CourseName,
			Description: course.Description,
			Thumbnail:   course.Thumbnail,
			Creator:     username,
		}

		courses = append(courses, &temp)
	}

	return courses, count, nil
}

func (serv *courseService) GetOnProgressCourse(ctx context.Context, meta *pagination.Meta, userId string) ([]*models.Course, uint64, error) {
	var (
		count   uint64
		courses []*models.Course
	)

	db_courses, count, err := serv.courseRepository.GetOnProgressCourseByUserID(ctx, serv.db, meta, userId)
	if err != nil {
		return courses, count, err
	}

	for _, course := range db_courses {

		var username string
		user, err := serv.userRepository.GetUserById(ctx, serv.db, course.Creator)
		if err != nil {
			return courses, count, err
		}

		if user == nil {
			username = "anon"
		} else {
			username = user.Username
		}

		temp := models.Course{
			ID:          course.ID,
			CourseName:  course.CourseName,
			Description: course.Description,
			Thumbnail:   course.Thumbnail,
			Creator:     username,
		}

		courses = append(courses, &temp)
	}

	return courses, count, nil
}

func (serv *courseService) GetCoursePagination(ctx context.Context, meta *pagination.Meta) ([]*models.Course, uint64, error) {
	var (
		count   uint64
		courses []*models.Course
	)

	db_courses, count, err := serv.courseRepository.GetCourseList(ctx, serv.db, meta)
	if err != nil {
		return courses, count, err
	}

	for _, course := range db_courses {

		var username string
		user, err := serv.userRepository.GetUserById(ctx, serv.db, course.Creator)
		if err != nil {
			return courses, count, err
		}

		if user == nil {
			username = "anon"
		} else {
			username = user.Username
		}

		temp := models.Course{
			ID:          course.ID,
			CourseName:  course.CourseName,
			Description: course.Description,
			Thumbnail:   course.Thumbnail,
			Creator:     username,
		}

		courses = append(courses, &temp)
	}

	return courses, count, nil
}

func (serv *courseService) GetCourseSyllabus(ctx context.Context, courseId string) (*models.SyllabusResponse, error) {
	db_syllabus, err := serv.courseRepository.GetCourseSyllabusByCourseID(ctx, serv.db, courseId)
	if err != nil {
		return nil, err
	}

	var sections []*models.Section

	sectionCount := 0
	for _, syllabus := range db_syllabus {
		if syllabus.Type == "section" {
			sectionCount += 1
			var subSections []*models.Material
			var sectionId = syllabus.ID

			for _, material := range db_syllabus {
				if material.SectionID != nil {
					if *material.SectionID == sectionId {
						temp := models.Material{
							ID:   material.ID,
							Name: material.Name,
							Type: material.Type,
						}

						subSections = append(subSections, &temp)
					}
				}
			}

			temp := models.Section{
				ID:          syllabus.ID,
				Name:        syllabus.Name,
				Subsections: subSections,
			}

			sections = append(sections, &temp)
		}
	}

	if sectionCount == 0 {
		var subSections []*models.Material
		for _, material := range db_syllabus {
			temp := models.Material{
				ID:   material.ID,
				Name: material.Name,
				Type: material.Type,
			}

			subSections = append(subSections, &temp)
		}

		section := models.Section{
			ID:          "0000",
			Name:        "Syllabus",
			Subsections: subSections,
		}

		sections = append(sections, &section)
	}

	data := models.SyllabusResponse{
		Syllabus: sections,
	}

	return &data, nil
}

func (serv *courseService) GetCourseMaterial(ctx context.Context, courseId string, sectionId string) (*models.SectionContentResponse, error) {
	db_material, err := serv.courseRepository.GetCourseMaterialByCourseIDAndSectionID(ctx, serv.db, courseId, sectionId)
	if err != nil {
		return nil, err
	}

	var materials []*models.MaterialContent

	for _, material := range db_material {
		temp := models.MaterialContent{
			ID:          material.ID,
			Name:        material.Name,
			Type:        material.Type,
			Content:     material.Content,
			ContentText: material.ContentText,
		}

		materials = append(materials, &temp)
	}

	data := models.SectionContentResponse{
		ID:          sectionId,
		Subsections: materials,
	}

	return &data, nil
}

func (serv *courseService) Enroll(ctx context.Context, userId string, courseId string) (*models.EnrollResponse, error) {
	check, err := serv.courseRepository.IsUserEnrolledToCourse(ctx, serv.db, userId, courseId)
	if err != nil {
		return nil, err
	}

	if check {
		return nil, er.NewError(fmt.Errorf("%s", "You are already enrolled to the course"), http.StatusBadRequest, nil)
	}

	values := &models.EnrollInput{
		UserID:   userId,
		CourseID: courseId,
	}

	err = serv.courseRepository.InsertEnrollment(ctx, serv.db, values)
	if err != nil {
		return nil, err
	}

	out := &models.EnrollResponse{
		Status:  "Success",
		Message: "You have successfully enrolled to the course.",
	}

	return out, nil
}

func (serv *courseService) StoreUserProgress(ctx context.Context, userId, materialId string) (*models.StoreProgressResponse, error) {
	courseId, err := serv.courseRepository.GetCourseIDByMaterialID(ctx, serv.db, materialId)
	if err != nil {
		return nil, err
	}

	check, err := serv.courseRepository.IsUserEnrolledToCourse(ctx, serv.db, userId, courseId)
	if err != nil {
		return nil, err
	}

	if !check {
		return nil, er.NewError(fmt.Errorf("%s", "You are not enrolled to the course"), http.StatusBadRequest, nil)
	}

	isLogged, err := serv.courseRepository.CheckIsProgressLogged(ctx, serv.db, userId, materialId)
	if err != nil {
		return nil, err
	}

	if isLogged {
		return nil, er.NewError(fmt.Errorf("%s", "Your progress has already been recorded"), http.StatusBadRequest, nil)
	}

	err = serv.courseRepository.StoreUserProgress(ctx, serv.db, materialId, courseId, userId, 100)
	if err != nil {
		return nil, err
	}

	out := &models.StoreProgressResponse{
		Status:  "Success",
		Message: "User progress updated successfully",
	}

	return out, nil
}

func (serv *courseService) ComputeUserProgress(ctx context.Context, userId, courseId string) (*models.GetProgressPercentageResponse, error) {
	check, err := serv.courseRepository.IsUserEnrolledToCourse(ctx, serv.db, userId, courseId)
	if err != nil {
		return nil, err
	}

	if !check {
		return nil, er.NewError(fmt.Errorf("%s", "You are not enrolled to the course"), http.StatusBadRequest, nil)
	}

	userProgress, err := serv.courseRepository.GetUserProgress(ctx, serv.db, userId, courseId)
	if err != nil {
		return nil, err
	}

	syllabus, err := serv.GetCourseSyllabus(ctx, courseId)
	if err != nil {
		return nil, err
	}

	materialDone := len(userProgress)
	courseLength := 0

	for _, section := range syllabus.Syllabus {
		for range section.Subsections {
			courseLength += 1
		}
	}

	intResult := 0
	if courseLength != 0 {
		result := float64(materialDone) / float64(courseLength) * 100
		intResult = int(math.Round(result))
	}

	out := &models.GetProgressPercentageResponse{
		Percentage: intResult,
	}

	return out, nil
}

func (serv *courseService) GetUserProgress(ctx context.Context, userId, courseId string) (*models.GetProgressResponse, error) {
	check, err := serv.courseRepository.IsUserEnrolledToCourse(ctx, serv.db, userId, courseId)
	if err != nil {
		return nil, err
	}

	if !check {
		return nil, er.NewError(fmt.Errorf("%s", "You are not enrolled to the course"), http.StatusBadRequest, nil)
	}

	userProgress, err := serv.courseRepository.GetUserProgress(ctx, serv.db, userId, courseId)
	if err != nil {
		return nil, err
	}

	var out []*models.UserProgress
	for _, progress := range userProgress {
		temp := models.UserProgress{
			UserID:     progress.UserID,
			CourseID:   progress.CourseID,
			MaterialID: progress.MaterialID,
			Score:      progress.Score,
		}

		out = append(out, &temp)
	}

	resp := &models.GetProgressResponse{
		Progress: out,
	}

	return resp, nil
}

func (serv *courseService) CreateNewCourse(ctx context.Context, course *models.CourseCreation) (*models.CourseCreationResponse, error) {
	newID := uuid.New().String()
	course.Course.ID = newID

	for _, sect := range course.Sections {
		newID = uuid.New().String()
		sect.ID = newID
		for _, mat := range sect.Subsections {
			newID = uuid.New().String()
			mat.ID = newID
		}
	}

	err := serv.courseRepository.InsertCourse(ctx, serv.db, course)
	if err != nil {
		return nil, err
	}

	resp := &models.CourseCreationResponse{
		Status:  "Success",
		Message: "Course Created Succesfully",
	}

	return resp, nil
}

func (serv *courseService) UploadImage(ctx context.Context, request *http.Request, baseURL string) (*models.UploadImageResponse, error) {

	imageId := uuid.New().String()
	url := imageId + ".png"

	request.ParseMultipartForm(10 << 20)

	file, _, err := request.FormFile("imageFile")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("static/image/", "course-*-"+url)
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// return that we have successfully uploaded our file!
	url = baseURL + tempFile.Name()
	resp := &models.UploadImageResponse{
		Status:  "Success",
		Message: "File Uploaded Succesfully",
		URL:     url,
	}

	return resp, nil
}

func (serv *courseService) CreateCourseDesc(ctx context.Context, course *models.CourseDescriptionInput, creatorId string) (*models.CourseCreationResponse, error) {
	id := uuid.New().String()

	input := db.Course{
		ID:          id,
		CourseName:  course.CourseName,
		Description: course.Description,
		Thumbnail:   course.Thumbnail,
		Creator:     creatorId,
	}

	err := serv.courseRepository.InsertCourseData(ctx, serv.db, &input)
	if err != nil {
		return nil, err
	}

	resp := &models.CourseCreationResponse{
		Status:  "Success",
		Message: "Course Description Created Succesfully",
		Id:      id,
	}

	return resp, nil
}

func (serv *courseService) CreateCourseSection(ctx context.Context, input *models.CourseSectionInput, creatorId string) (*models.CourseCreationResponse, error) {
	id := uuid.New().String()

	material := &db.Material{
		ID:          id,
		CourseID:    input.CourseID,
		Name:        input.Name,
		Type:        "section",
		SectionID:   "",
		Content:     "",
		ContentText: "",
	}

	err := serv.courseRepository.InsertCourseMaterial(ctx, serv.db, material)
	if err != nil {
		return nil, err
	}

	resp := &models.CourseCreationResponse{
		Status:  "Success",
		Message: "Section Created Succesfully",
		Id:      id,
	}

	return resp, nil
}

func (serv *courseService) CreateCourseMaterial(ctx context.Context, input *models.CourseMaterialInput, creatorId string) (*models.CourseCreationResponse, error) {
	id := uuid.New().String()
	courseId, err := serv.courseRepository.GetCourseIDByMaterialID(ctx, serv.db, input.SectionID)
	if err != nil {
		return nil, err
	}

	_, err = serv.courseRepository.GetMaterialByID(ctx, serv.db, input.SectionID)
	if err != nil {
		return nil, err
	}

	material := &db.Material{
		ID:          id,
		CourseID:    courseId,
		Name:        input.Name,
		Type:        input.Type,
		SectionID:   input.SectionID,
		Content:     input.Content,
		ContentText: input.ContentText,
	}

	err = serv.courseRepository.InsertCourseMaterial(ctx, serv.db, material)
	if err != nil {
		return nil, err
	}

	resp := &models.CourseCreationResponse{
		Status:  "Success",
		Message: "Section Created Succesfully",
		Id:      id,
	}

	return resp, nil
}

func (serv *courseService) GetCourseByCreatorID(ctx context.Context, creatorId string) ([]*models.Course, error) {
	db_courses, err := serv.courseRepository.GetCourseByCreatorID(ctx, serv.db, creatorId)
	if err != nil {
		return nil, err
	}

	var courses []*models.Course
	for _, course := range db_courses {
		courses = append(courses, &models.Course{
			ID:          course.ID,
			CourseName:  course.CourseName,
			Description: course.Description,
			Thumbnail:   course.Thumbnail,
			Creator:     course.Creator,
		})
	}

	return courses, nil
}
