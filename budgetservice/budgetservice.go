package budgetservice

import "time"

type BudgetService struct {
	repo IRepo
}

func NewBudgetService(repo IRepo) *BudgetService {
	return &BudgetService{repo: repo}
}

func (s BudgetService) Query(start time.Time, end time.Time) float64 {
	budgets := s.repo.GetAll()
	if s.isNoBudget(budgets) || s.isNoOverlap(end, budgets, start) {
		return 0
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
