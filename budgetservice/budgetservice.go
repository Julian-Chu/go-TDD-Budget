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
	days := float64(end.Sub(start).Hours() / 24)
	return budgets[0].Amount / days
}
