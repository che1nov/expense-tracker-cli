package main

import (
	"expense-tracker-cli/intrernal/cli"
	"expense-tracker-cli/intrernal/repository"
	"expense-tracker-cli/intrernal/usecase"
	"os"
)

func main() {
	// Initialize repository, use case, and CLI
	repo := repository.NewExpenseRepository("expenses.json")
	usecase := usecase.NewExpenseUsecase(*repo)
	cli := cli.NewCLI(usecase)

	// Run the CLI with command-line arguments
	cli.Run(os.Args)
}
