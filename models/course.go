package models

type Course struct {
	ID          string `json:"id"`
	CourseName  string `json:"course_name"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Creator     string `json:"creator"`
	// TODO: Topic
}

type EnrollInput struct {
	UserID   string `db:"user_id"`
	CourseID string `db:"course_id"`
}

type EnrollResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type StoreProgressResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetProgressPercentageResponse struct {
	Percentage int `json:"percentage"`
}

type UserProgress struct {
	UserID     string `json:"userId"`
	CourseID   string `json:"courseId"`
	MaterialID string `json:"materialId"`
	Score      int    `json:"score"`
}

type GetProgressResponse struct {
	Progress []*UserProgress `json:"progress"`
}

type CourseCreation struct {
	Course   Course             `json:"data" validate:"required" label:"data"`
	Sections []*SectionCreation `json:"sections" validate:"required" label:"sections"` 
}

type SectionCreation struct {
	ID          string              `json:"sectionID"`
	Name        string              `json:"sectionName" validate:"required" label:"sectionName"`
	Subsections []*MaterialCreation `json:"subSections" validate:"required" label:"subSections"`
}

type MaterialCreation struct {
	ID          string `json:"materialID"`
	Name        string `json:"materialName" validate:"required" label:"materialName"`
	Type        string `json:"materialType" validate:"required" label:"materialType"`
	Content     string `json:"materialContent" validate:"required"`
	ContentText string `json:"materialContentText" validate:"required"`
}

type CourseCreationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Id      string `json:"id"`
}

type UploadImageResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	URL     string `json:"imageURL"`
}

type CourseDescriptionInput struct {
	CourseName  string `json:"course_name" validate:"required" label:"course_name"`
	Description string `json:"description" validate:"required" label:"description"`
	Thumbnail   string `json:"thumbnail" validate:"required" label:"thumbnail"`
}

type CourseSectionInput struct {
	Name     string `json:"sectionName" validate:"required" label:"sectionName"`
	CourseID string `json:"courseID" validate:"required" label:"courseID"`
}

type CourseMaterialInput struct {
	Name   string `json:"materialName" validate:"required" label:"materialName"`
	Type        string `json:"materialType" validate:"required" label:"materialType"`
	Content     string `json:"materialContent"`
	ContentText string `json:"materialContentText"`
	SectionID string `json:"sectionID" validate:"required" label:"sectionID"`
}