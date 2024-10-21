package repository

import (
	"fmt"
	. "server/model"
)

type ExpenseRepository struct {
	dbAdapter *DbAdapter[Expense]
}

func CreateExpenseRepository(dbAdapter *DbAdapter[Expense]) *ExpenseRepository {
	return &ExpenseRepository{
		dbAdapter: dbAdapter,
	}
}

func (r *ExpenseRepository) Insert(expense Expense) (Expense, error) {
	query := "INSERT INTO public.expense (amount) VALUES ($1) RETURNING id"
	err := r.dbAdapter.ExecuteUpdatingField(query, []interface{}{expense.Amount}, &expense.ID)
	if err != nil {
		return Expense{}, fmt.Errorf("could not insert expense: %v", err)
	}
	return expense, nil
}

func (r *ExpenseRepository) GetAll() ([]Expense, error) {
	query := "SELECT id, amount, timestamp FROM public.expense"
	return r.dbAdapter.ExecuteForAll(query)
}
