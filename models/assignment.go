package models

type AssignmentResponse struct {
	ID         string               `json:"id"`
	Creator    string               `json:"creator"`
	Title      string               `json:"title"`
	Duration   int                  `json:"duration"`
	Topic      string               `json:"topic"`
	Difficulty string               `json:"difficulty"`
	Problems   []*ProblemTypeDetail `json:"problems"`
}

type ProblemTypeDetail struct {
	ID       string      `json:"id"`
	Type     string      `json:"type"`
	Choice   interface{} `json:"choice"`
	Question interface{} `json:"question"`
}

type ProblemTypeDetailAnswer struct {
	ID 			string 			`json:"id"`
	Type 		string			`json:"type"`
	Choice 	interface{}	`json:"choice"`
	Answer 	interface{}	`json:"answer"`
}

type Assignment struct {
	ID	        string `json:"id"`
	Creator	    string `json:"creator"`
	Title       string `json:"title" validate:"required" label:"title"`
	Duration    int    `json:"duration" validate:"required" label:"duration"`
	Topic	    string `json:"topic"`
	Difficulty	string `json:"difficulty"`
}

type AssignmentProblem struct {
	ProblemID    string `json:"problem_id" validate:"required" label:"problem_id"`
}

type AssignmentCreation struct {
	Desc     Assignment          `json:"desc" validate:"required" label:"desc"`
	Problems []AssignmentProblem `json:"list_problem_id" validate:"required" label:"lisT_problem_id"`
}

type AssignmentCreationResponse struct {
	Status  string `json:"status"` 
	Message string `json:"message"`
	ID      string `json:"assignmentID"`
}

type AssignmentSubmission struct {
	ID 			string 					`json:"id" validate:"required"`
	Answers []ProblemAnswer	`json:"answers" validate:"required"` 
}

type ProblemAnswer struct {
	ID 			string			`json:"id"`
	Type 		string			`json:"type"`
	Answer 	interface{}	`json:"answer"`
}

type AssignmentScore struct {
	Score int `json:"score"`
}