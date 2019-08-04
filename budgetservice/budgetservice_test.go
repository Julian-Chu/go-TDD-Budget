package budgetservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type BudgetService struct {
}

func (s BudgetService) Query(date time.Time, time time.Time) float64 {
	return 0
}

func Test_NoBudget(t *testing.T) {
	s := &BudgetService{}
	actual := s.Query(
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(0)
	assert.Equal(t, expected, actual, "")

}
