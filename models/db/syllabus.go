package db

type Syllabus struct {
	ID        string  `db:"id"`
	CourseID  string  `db:"course_id"`
	Name      string  `db:"name"`
	Type      string  `db:"type"`
	SectionID *string `db:"section_id"`
}

type Material struct {
	ID 					string `db:"id"`
	CourseID		string `db:"course_id"`
	Name				string `db:"name"`
	Type				string `db:"type"`
	SectionID		string `db:"section_id"`
	Content			string `db:"content"`
	ContentText	string `db:"content_text"`
}