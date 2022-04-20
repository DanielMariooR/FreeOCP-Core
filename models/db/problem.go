package db

type ProblemCandidate struct {
	ID         string `db:"id"`
	Creator    string `db:"creator"`
	Title      string `db:"title"`
	Type       string `db:"type"`
	Topic      string `db:"topic"`
	Difficulty string `db:"difficulty"`
	Status     string `db:"status"`
	Detail     string `db:"detail"`
}

type ProblemDetail struct {
	ID     string `db:"id"`
	Detail string `db:"detail"`
}

type ProblemTypeDetail struct {
	ID     string `db:"id"`
	Detail string `db:"detail"`
	Type   string `db:"type"`
}
