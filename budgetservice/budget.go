package budgetservice

import (
	"fmt"
	"time"
)

type Budget struct {
	YearMonth string
	Amount    float64
}

func (b Budget) GetDays() float64 {
	date := b.getDate()
	return float64(date.AddDate(0, 1, 0).Sub(date).Hours() / 24)
}

func (b Budget) getDate() time.Time {
	date, err := time.Parse("200601", b.YearMonth)
	if err != nil {
		fmt.Print(err.Error())
	}
	return date
}

func (b Budget) FirstDay() time.Time {
	return b.getDate()
}

func (b Budget) LastDay() time.Time {
	return b.FirstDay().AddDate(0, 1, -1)
}

func (b Budget) DailyAmount() float64 {
	return b.Amount / b.GetDays()
}
