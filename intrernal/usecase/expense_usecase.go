package usecase

import (
	"errors"
	"expense-tracker-cli/intrernal/entity"
	"expense-tracker-cli/intrernal/repository"
	"time"
)

// ExpenseUsecase contains business logic for managing expenses
type ExpenseUsecase struct {
	repo repository.ExpenseRepository
}

// NewExpenseUsecase creates a new instance of ExpenseUsecase
func NewExpenseUsecase(repo repository.ExpenseRepository) *ExpenseUsecase {
	return &ExpenseUsecase{repo: repo}
}

// AddExpense adds a new expense
func (uc *ExpenseUsecase) AddExpense(description string, amount float64, category string) (int, error) {
	if amount <= 0 {
		return 0, errors.New("amount must be positive")
	}

	expenses, err := uc.repo.Load()
	if err != nil {
		return 0, err
	}

	newExpense := entity.Expense{
		ID:          generateExpenseID(expenses),
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
		Category:    category,
	}

	expenses = append(expenses, newExpense)
	err = uc.repo.Save(expenses)
	if err != nil {
		return 0, err
	}

	return newExpense.ID, nil
}

// DeleteExpense deletes an expense by ID
func (uc *ExpenseUsecase) DeleteExpense(id int) error {
	expenses, err := uc.repo.Load()
	if err != nil {
		return err
	}

	newExpenses := []entity.Expense{}
	for _, expense := range expenses {
		if expense.ID != id {
			newExpenses = append(newExpenses, expense)
		}
	}

	if len(newExpenses) == len(expenses) {
		return errors.New("expense not found")
	}

	return uc.repo.Save(newExpenses)
}

// ListExpenses returns all expenses
func (uc *ExpenseUsecase) ListExpenses() ([]entity.Expense, error) {
	return uc.repo.Load()
}

// GetSummary calculates the total amount of all expenses
func (uc *ExpenseUsecase) GetSummary() (float64, error) {
	expenses, err := uc.repo.Load()
	if err != nil {
		return 0, err
	}

	total := 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}

	return total, nil
}

// GetMonthlySummary calculates the total amount of expenses for a specific month
func (uc *ExpenseUsecase) GetMonthlySummary(month int) (float64, error) {
	expenses, err := uc.repo.Load()
	if err != nil {
		return 0, err
	}

	total := 0.0
	now := time.Now()
	for _, expense := range expenses {
		if expense.Date.Month() == time.Month(month) && expense.Date.Year() == now.Year() {
			total += expense.Amount
		}
	}

	return total, nil
}

// Generate unique ID for expenses
func generateExpenseID(expenses []entity.Expense) int {
	if len(expenses) == 0 {
		return 1
	}

	maxID := 0
	for _, expense := range expenses {
		if expense.ID > maxID {
			maxID = expense.ID
		}
	}
	return maxID + 1
}
