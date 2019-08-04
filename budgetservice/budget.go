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
	date, err := time.Parse("200601", b.YearMonth)

	if err != nil {
		fmt.Print(err.Error())
	}
	return float64(date.AddDate(0, 1, 0).Sub(date).Hours() / 24)
}

func (b Budget) FirstDay() time.Time {
	date, err := time.Parse("200601", b.YearMonth)
	if err != nil {
		fmt.Print(err.Error())
	}
	return date
}
