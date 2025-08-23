// Package challenge7 contains the solution for Challenge 7: Bank Account with Error Handling.
package challenge7

import (
	"sync"
	// Add any other necessary imports
)

// BankAccount represents a bank account with balance management and minimum balance requirements.
type BankAccount struct {
	ID         string
	Owner      string
	Balance    float64
	MinBalance float64
	mu         sync.Mutex // For thread safety
}

// Constants for account operations
const (
	MaxTransactionAmount = 10000.0 // Example limit for deposits/withdrawals
)

// Custom error types

// AccountError is a general error type for bank account operations.
type AccountError struct {
	// Implement this error type
}

func (e *AccountError) Error() string {
	// Implement error message
	return "Account error"
}

// InsufficientFundsError occurs when a withdrawal or transfer would bring the balance below minimum.
type InsufficientFundsError struct {
	// Implement this error type
}

func (e *InsufficientFundsError) Error() string {
	// Implement error message
	return "Insuffcient funds error"
}

// NegativeAmountError occurs when an amount for deposit, withdrawal, or transfer is negative.
type NegativeAmountError struct {
	// Implement this error type
}

func (e *NegativeAmountError) Error() string {
	// Implement error message
	return "Negative amount error"
}

// ExceedsLimitError occurs when a deposit or withdrawal amount exceeds the defined limit.
type ExceedsLimitError struct {
	// Implement this error type
}

func (e *ExceedsLimitError) Error() string {
	// Implement error message
	return "Amount exceeds limit"
}

// NewBankAccount creates a new bank account with the given parameters.
// It returns an error if any of the parameters are invalid.
func NewBankAccount(id, owner string, initialBalance, minBalance float64) (*BankAccount, error) {
	// Implement account creation with validation
	if initialBalance < 0 || minBalance < 0 {
		return nil, &NegativeAmountError{}
	} else if initialBalance < minBalance {
		return nil, &InsufficientFundsError{}
	} else if id == "" || owner == "" {
		return nil, &AccountError{}
	}

	return &BankAccount{
		ID:         id,
		Owner:      owner,
		Balance:    initialBalance,
		MinBalance: minBalance,
	}, nil
}

// Deposit adds the specified amount to the account balance.
// It returns an error if the amount is invalid or exceeds the transaction limit.
func (a *BankAccount) Deposit(amount float64) error {
	// Implement deposit functionality with proper error handling
	if amount < 0 {
		return &NegativeAmountError{}
	} else if amount > MaxTransactionAmount {
		return &ExceedsLimitError{}
	}

	a.mu.Lock()
	a.Balance += amount
	a.mu.Unlock()

	return nil
}

// Withdraw removes the specified amount from the account balance.
// It returns an error if the amount is invalid, exceeds the transaction limit,
// or would bring the balance below the minimum required balance.
func (a *BankAccount) Withdraw(amount float64) error {
	// Implement withdrawal functionality with proper error handling
	if amount > MaxTransactionAmount {
		return &ExceedsLimitError{}
	} else if a.Balance-amount < a.MinBalance {
		return &InsufficientFundsError{}
	} else if amount < 0 {
		return &NegativeAmountError{}
	}

	a.mu.Lock()
	a.Balance -= amount
	a.mu.Unlock()

	return nil
}

// Transfer moves the specified amount from this account to the target account.
// It returns an error if the amount is invalid, exceeds the transaction limit,
// or would bring the balance below the minimum required balance.
func (a *BankAccount) Transfer(amount float64, target *BankAccount) error {
	// Implement transfer functionality with proper error handling
	if amount < 0 {
		return &NegativeAmountError{}
	} else if amount > MaxTransactionAmount {
		return &ExceedsLimitError{}
	} else if a.Balance-amount < a.MinBalance {
		return &InsufficientFundsError{}
	}

	a.mu.Lock()
	a.Balance -= amount
	a.mu.Unlock()

	target.mu.Lock()
	target.Balance += amount
	target.mu.Unlock()

	return nil
}

