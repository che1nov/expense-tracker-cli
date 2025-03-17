package cli

import (
	"expense-tracker-cli/intrernal/entity"
	"expense-tracker-cli/intrernal/usecase"
	"fmt"
	"strconv"
	"strings"
)

// CLI manages the command-line interface
type CLI struct {
	usecase *usecase.ExpenseUsecase
}

// NewCLI creates a new instance of CLI
func NewCLI(usecase *usecase.ExpenseUsecase) *CLI {
	return &CLI{usecase: usecase}
}

// Run starts the CLI
func (c *CLI) Run(args []string) {
	if len(args) < 2 {
		printUsage()
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 5 {
			fmt.Println("Error: Missing arguments. Usage: add --description <desc> --amount <amount>")
			return
		}
		desc, amount, category := parseAddArgs(args[2:])
		id, err := c.usecase.AddExpense(desc, amount, category)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", id)

	case "delete":
		if len(args) < 3 {
			fmt.Println("Error: Missing ID. Usage: delete --id <id>")
			return
		}
		id, err := strconv.Atoi(strings.TrimPrefix(args[2], "--id="))
		if err != nil {
			fmt.Println("Error: Invalid ID.")
			return
		}
		err = c.usecase.DeleteExpense(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Expense deleted successfully (ID: %d)\n", id)

	case "list":
		expenses, err := c.usecase.ListExpenses()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printExpenses(expenses)

	case "summary":
		if len(args) > 2 && strings.HasPrefix(args[2], "--month=") {
			month, err := strconv.Atoi(strings.TrimPrefix(args[2], "--month="))
			if err != nil || month < 1 || month > 12 {
				fmt.Println("Error: Invalid month.")
				return
			}
			total, err := c.usecase.GetMonthlySummary(month)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("Total expenses for month %d: $%.2f\n", month, total)
		} else {
			total, err := c.usecase.GetSummary()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("Total expenses: $%.2f\n", total)
		}

	default:
		fmt.Printf("Error: Unknown command '%s'.\n", command)
		printUsage()
	}
}

// printUsage prints usage instructions
func printUsage() {
	fmt.Println("Usage: expense-tracker <command> [args]")
	fmt.Println("Commands:")
	fmt.Println("  add --description <desc> --amount <amount> [--category <category>]  Add a new expense")
	fmt.Println("  delete --id <id>                                                   Delete an expense")
	fmt.Println("  list                                                              List all expenses")
	fmt.Println("  summary [--month=<month>]                                         Show total expenses or for a specific month")
}

// parseAddArgs parses arguments for the 'add' command
func parseAddArgs(args []string) (string, float64, string) {
	var desc, category string
	var amount float64

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--description":
			desc = args[i+1]
			i++
		case "--amount":
			amount, _ = strconv.ParseFloat(args[i+1], 64)
			i++
		case "--category":
			category = args[i+1]
			i++
		}
	}

	return desc, amount, category
}

// printExpenses displays a list of expenses
func printExpenses(expenses []entity.Expense) {
	fmt.Println("ID\tDate\t\tDescription\tAmount\tCategory")
	for _, expense := range expenses {
		fmt.Printf("%d\t%s\t%s\t$%.2f\t%s\n",
			expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount, expense.Category)
	}
}
