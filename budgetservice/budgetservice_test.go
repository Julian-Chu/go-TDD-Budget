package budgetservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var m = &MockRepo{func() []Budget {
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

	expected := float64(1)
	assert.Equal(t, expected, actual, "")
}

func Test_no_overlap_before_budget_firstday(t *testing.T) {
	m.getAll = func() []Budget {
		return []Budget{
			{YearMonth: "201904", Amount: 30},
		}
	}
	actual := service.Query(
		time.Date(2019, 03, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 03, 31, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(0)
	assert.Equal(t, expected, actual, "")
}

func Test_no_overlap_after_budget_lastday(t *testing.T) {
	m.getAll = func() []Budget {
		return []Budget{
			{YearMonth: "201904", Amount: 30},
		}
	}
	actual := service.Query(
		time.Date(2019, 05, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 05, 01, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(0)
	assert.Equal(t, expected, actual, "")
}

func Test_period_overlap_budget_lastday(t *testing.T) {
	m.getAll = func() []Budget {
		return []Budget{
			{YearMonth: "201904", Amount: 30},
		}
	}
	actual := service.Query(
		time.Date(2019, 04, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 05, 01, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(1)
	assert.Equal(t, expected, actual, "")
}

func Test_period_cross_budget_month(t *testing.T) {
	m.getAll = func() []Budget {
		return []Budget{
			{YearMonth: "201904", Amount: 30},
		}
	}
	actual := service.Query(
		time.Date(2019, 03, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 05, 01, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(30)
	assert.Equal(t, expected, actual, "")
}

func Test_invalid_period(t *testing.T) {
	m.getAll = func() []Budget {
		return []Budget{
			{YearMonth: "201904", Amount: 30},
		}
	}
	actual := service.Query(
		time.Date(2019, 04, 02, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(0)
	assert.Equal(t, expected, actual, "")
}

func Test_dailyamount_is_10(t *testing.T) {
	m.getAll = func() []Budget {
		return []Budget{
			{YearMonth: "201904", Amount: 300},
		}
	}
	actual := service.Query(
		time.Date(2019, 04, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 04, 02, 0, 0, 0, 0, time.UTC),
	)

	expected := float64(20)
	assert.Equal(t, expected, actual, "")
}

type MockRepo struct {
	getAll func() []Budget
}

func (r MockRepo) GetAll() []Budget {
	return r.getAll()
}
