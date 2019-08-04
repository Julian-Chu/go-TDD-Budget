package budgetservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var m = MockRepo{func() []Budget {
	return []Budget{}
}}
var service = NewBudgetService(m)

func Test_NoBudget(t *testing.T) {
	m.getAll = func() []Budget {
		return []Budget{}
	}
	actual := service.Query(
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(0)
	assert.Equal(t, expected, actual, "")
}

func Test_period_inside_budget_month(t *testing.T) {
	m.getAll = func() []Budget {
		return []Budget{
			{YearMonth: "201904", Amount: 30},
		}
	}
	actual := service.Query(
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(0)
	assert.Equal(t, expected, actual, "")
}

type MockRepo struct {
	getAll func() []Budget
}

func (r MockRepo) GetAll() []Budget {
	return r.getAll()
}
