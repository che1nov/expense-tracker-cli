package repository

import (
	"encoding/json"
	"expense-tracker-cli/intrernal/entity"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// ExpenseRepository manages the storage of expenses
type ExpenseRepository struct {
	filePath string // Path to the JSON file
	mu       sync.Mutex
}

// NewExpenseRepository creates a new instance of ExpenseRepository
func NewExpenseRepository(filePath string) *ExpenseRepository {
	return &ExpenseRepository{filePath: filePath}
}

// Load loads expenses from the file
func (r *ExpenseRepository) Load() ([]entity.Expense, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var expenses []entity.Expense

	// Check if the file exists
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		return expenses, nil
	}

	// Read the file
	data, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal JSON data
	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return expenses, nil
}

// Save saves expenses to the file
func (r *ExpenseRepository) Save(expenses []entity.Expense) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Marshal data into JSON format
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write data to the file
	err = ioutil.WriteFile(r.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
