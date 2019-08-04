package budgetservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type BudgetService struct {
	repo IRepo
}

func NewBudgetService(repo IRepo) *BudgetService {
	return &BudgetService{repo: repo}
}

func (s BudgetService) Query(date time.Time, time time.Time) float64 {
	budgets := s.repo.GetAll()
	if len(budgets) == 0 {
		return 0
	}
	days := float64(time.Day() - date.Day() + 1)
	return budgets[0].Amount / days
}

func Test_NoBudget(t *testing.T) {
	m := MockRepo{}
	m.getAll = func() []Budget {
		return []Budget{}
	}
	s := NewBudgetService(m)
	actual := s.Query(
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(0)
	assert.Equal(t, expected, actual, "")
}

func Test_period_inside_budget_month(t *testing.T) {
	m := MockRepo{}
	m.getAll = func() []Budget {
		return []Budget{
			{YearMonth: "201904", Amount: 30},
		}
	}
	s := NewBudgetService(m)
	actual := s.Query(
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(0)
	assert.Equal(t, expected, actual, "")
}

type Budget struct {
	YearMonth string
	Amount    float64
}

type IRepo interface {
	GetAll() []Budget
}

type MockRepo struct {
	getAll func() []Budget
}

func (r MockRepo) GetAll() []Budget {
	return r.getAll()
}
