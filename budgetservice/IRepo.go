package budgetservice

type IRepo interface {
	GetAll() []Budget
}
