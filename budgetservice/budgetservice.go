package budgetservice

import "time"

type BudgetService struct {
	repo IRepo
}

func NewBudgetService(repo IRepo) *BudgetService {
	return &BudgetService{repo: repo}
}

type Period struct {
	Start time.Time
	End   time.Time
}

func (s BudgetService) Query(queryStart time.Time, queryEnd time.Time) float64 {
	budgets := s.repo.GetAll()
	if s.isNoBudget(budgets) {
		return 0
	}
	period := Period{queryStart, queryEnd}
	sum := float64(0)
	for _, budget := range budgets {
		if period.isNoOverlap(budget) {
			continue
		}
		end := period.End
		if budget.LastDay().Before(period.End) {
			end = budget.LastDay()
		}

		start := period.Start
		if budget.FirstDay().After(period.Start) {
			start = budget.FirstDay()
		}
		days := float64(end.Sub(start).Hours()/24 + 1)
		sum += budget.DailyAmount() * days
	}
	return sum
}

func (s BudgetService) isNoBudget(budgets []Budget) bool {
	return len(budgets) == 0
}

func (p Period) isNoOverlap(budget Budget) bool {
	return p.End.Before(budget.FirstDay()) || p.Start.After(budget.LastDay())
}
