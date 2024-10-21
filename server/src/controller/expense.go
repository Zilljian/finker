package controllers

import (
	"encoding/json"
	"net/http"
	"server/controller/util"
	. "server/model"
	"server/repository"
)

type ExpenseController struct {
	repository *repository.ExpenseRepository
}

func CreateExpenseController(repository *repository.ExpenseRepository) *ExpenseController {
	return &ExpenseController{
		repository: repository,
	}
}

func (c *ExpenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var newExpense Expense
	if err := json.NewDecoder(r.Body).Decode(&newExpense); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	expense, e := c.repository.Insert(newExpense)
	if e != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}

func (c *ExpenseController) GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, e := c.repository.GetAll()
	if e != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	util.SetHeaders(w)
	json.NewEncoder(w).Encode(expenses)
}
