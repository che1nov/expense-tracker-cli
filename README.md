# Expense Tracker

Expense Tracker is a simple command-line application for managing your personal finances. It allows users to add, update, delete, and view their expenses. The application also provides summaries of expenses for better financial management.

## Features

- Add expenses with a description and amount.
- Update existing expenses.
- Delete expenses.
- View all expenses.
- View a summary of all expenses.
- View a summary of expenses for a specific month (current year).

### Additional Features

- Add categories for expenses and filter expenses by category.
- Set a budget for each month and show a warning if the user exceeds the budget.
- Export expenses to a CSV file.

## Commands and Expected Output

### Add Expense

```sh
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)
```

### List Expenses

```sh
$ expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10
```

### View Summary

```sh
$ expense-tracker summary
# Total expenses: $30
```

### Delete Expense

```sh
$ expense-tracker delete --id 2
# Expense deleted successfully
```

### View Summary After Deletion

```sh
$ expense-tracker summary
# Total expenses: $20
```

### View Monthly Summary

```sh
$ expense-tracker summary --month 8
# Total expenses for August: $20
```

## Implementation

You can implement the application using any programming language of your choice. Here are some suggestions:

- Use a command-line argument parsing module (e.g., `argparse` for Python, `commander` for Node.js).
- Use a simple text file to store expense data. You can use JSON, CSV, or any other format for data storage.
- Add error handling to manage invalid inputs and edge cases (e.g., negative amounts, non-existent expense IDs).
- Use functions to modularize the code, making it easier to test and maintain.

## Example Usage

### Adding Expenses

```sh
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)
```

### Listing Expenses

```sh
$ expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10
```

### Viewing Summary

```sh
$ expense-tracker summary
# Total expenses: $30
```

### Deleting an Expense

```sh
$ expense-tracker delete --id 2
# Expense deleted successfully
```

### Viewing Updated Summary

```sh
$ expense-tracker summary
# Total expenses: $20
```

### Viewing Monthly Summary

```sh
$ expense-tracker summary --month 8
# Total expenses for August: $20
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

https://roadmap.sh/projects/expense-tracker