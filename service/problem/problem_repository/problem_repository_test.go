package problem_repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/problem/problem_repository"
)

func TestProblemRepository_GetTableName(t *testing.T) {
	r := problem_repository.NewRepository()

	expectedTableName := "Candidate_Problem"
	assert.Equal(t, expectedTableName, r.GetTableName())
}

func TestProblemRepository_GetDetailTableName(t *testing.T) {
	r := problem_repository.NewRepository()

	expectedTableName := "Detail_Problem"
	assert.Equal(t, expectedTableName, r.GetDetailTableName())
}
