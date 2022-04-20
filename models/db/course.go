package db

type Course struct {
	ID          string `db:"id"`
	CourseName  string `db:"course_name"`
	Description string `db:"description"`
	Thumbnail   string `db:"thumbnail"`
	Creator     string `db:"creator"`
	// TODO: Topic
}

type UserProgress struct {
	UserID     string `db:"user_id"`
	CourseID   string `db:"course_id"`
	MaterialID string `db:"material_id"`
	Score      int    `db:"score"`
}
