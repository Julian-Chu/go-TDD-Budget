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
	if s.isNoBudget(budgets) {
		return 0
	}

	sum := float64(0)
	for _, budget := range budgets {
		if s.isNoOverlap(queryEnd, queryStart, budget) {
			continue
		}
		end := queryEnd
		if budget.LastDay().Before(queryEnd) {
			end = budget.LastDay()
		}

		start := queryStart
		if budget.FirstDay().After(queryStart) {
			start = budget.FirstDay()
		}
		days := float64(end.Sub(start).Hours()/24 + 1)
		getDays := budget.GetDays()
		sum += (budget.Amount / getDays) * days
	}
	return sum
}

func (s BudgetService) isNoBudget(budgets []Budget) bool {
	return len(budgets) == 0
}

func (s BudgetService) isNoOverlap(end time.Time, start time.Time, budget Budget) bool {
	return end.Before(budget.FirstDay()) || start.After(budget.LastDay())
}
