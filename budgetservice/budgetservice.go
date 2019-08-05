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
	if len(budgets) == 0 {
		return 0
	}
	if end.Before(budgets[0].FirstDay()) {
		return 0
	}
	if start.After(budgets[0].LastDay()) {
		return 0
	}
	days := float64(end.Sub(start).Hours()/24 + 1)
	getDays := budgets[0].GetDays()
	return (budgets[0].Amount / getDays) * days
}
