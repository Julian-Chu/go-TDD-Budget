package budgetservice

import "time"

type BudgetService struct {
	repo IRepo
}

func NewBudgetService(repo IRepo) *BudgetService {
	return &BudgetService{repo: repo}
}

func (s BudgetService) Query(queryStart time.Time, queryEnd time.Time) float64 {
	budgets := s.repo.GetAll()
	if s.isNoBudget(budgets) || s.isNoOverlap(queryEnd, budgets, queryStart) {
		return 0
	}
	end := queryEnd
	if budgets[0].LastDay().Before(queryEnd) {
		end = budgets[0].LastDay()
	}

	start := queryStart
	if budgets[0].FirstDay().After(queryStart) {
		start = budgets[0].FirstDay()
	}
	days := float64(end.Sub(start).Hours()/24 + 1)
	getDays := budgets[0].GetDays()
	return (budgets[0].Amount / getDays) * days
}

func (s BudgetService) isNoBudget(budgets []Budget) bool {
	return len(budgets) == 0
}

func (s BudgetService) isNoOverlap(end time.Time, budgets []Budget, start time.Time) bool {
	return end.Before(budgets[0].FirstDay()) || start.After(budgets[0].LastDay())
}
