package models

import (
	"github.com/labstack/echo/v4"
)

type ProblemCandidate struct {
	ID         string      `json:"id"`
	Creator    string      `json:"creator"`
	Title      string      `json:"title"`
	Type       string      `json:"type"`
	Topic      string      `json:"topic"`
	Difficulty string      `json:"difficulty"`
	Status     string      `json:"status"`
	Detail     interface{} `json:"content"`
}

type ProblemStatus struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Topic      string `json:"topic"`
	Difficulty string `json:"difficulty"`
	Status     string `json:"status"`
}

type ProblemDetail struct {
	ID 		string `json:"id"`
	Detail 	interface{} `json:"content"`
}

type ProblemCreationInput struct {
	Creator    string `json:"creator" validate:"required" label:"creator"`
	Title      string `json:"title" validate:"required" label:"title"`
	Type       string `json:"type" validate:"required" label:"type"`
	Topic      string `json:"topic" validate:"required" label:"topic"`
	Difficulty string `json:"difficulty" validate:"required" label:"difficulty"`
	Detail     string `json:"detail" validate:"required" label:"content"`
}

type ProblemCreationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ProblemCandidateList struct {
	Problems []*ProblemCandidateTable
}

type ProblemCandidateTable struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Topic      string `json:"topic"`
	Difficulty string `json:"difficulty"`
}

type ProblemFilter struct {
	Difficulty string `json:"difficulty"`
	Category   string `json:"category"`
}

func (f *ProblemFilter) FromContext(c echo.Context) *ProblemFilter {
	f.Difficulty = c.QueryParam("difficulty")
	f.Category = c.QueryParam("category")
	return f
}

type ProblemStatusList struct {
	Problems []*ProblemStatus
}

type ProblemStatusUpdate struct {
	Id     string
	Status string
}