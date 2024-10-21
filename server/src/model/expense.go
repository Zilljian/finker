package model

import "time"

type Expense struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Amount    float64   `json:"amount"`
	Expenses  []string  `json:"Expenses"`
}

type ExpenseFactory struct{}

func CreateExpenseFactory() ExpenseFactory {
	return ExpenseFactory{}
}

func (ExpenseFactory) CreateTemplate() (*Expense, []interface{}) {
	var Expense Expense
	return &Expense, []interface{}{&Expense.ID, &Expense.Amount, &Expense.Timestamp}
}
