package db

type Assignment struct {
	ID         string `db:"id"`
	Creator    string `db:"creator"`
	Title      string `db:"title"`
	Duration   int    `db:"duration"`
	Topic      string `db:"topic"`
	Difficulty string `db:"difficulty"`
}

type AssignmentProblem struct {
	AssignmentID string `db:"assignment_id"`
	ProblemID    string `db:"problem_id"`
}

type AssignmentCreation struct {
	Desc     Assignment
	Problems []AssignmentProblem
}
