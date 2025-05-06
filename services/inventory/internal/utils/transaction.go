package utils

import (
	"context"

	"gorm.io/gorm"
)

// TransactionFunc represents a function that will be executed within a transaction
type TransactionFunc func(tx *gorm.DB) error

// WithTransaction executes the given function within a database transaction
// It will commit the transaction if the function returns nil, and will rollback if it returns an error
func WithTransaction(db *gorm.DB, fn TransactionFunc) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after rollback
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// TransactionalContext is a context key for storing transaction information
type TransactionalContext struct{}

// GetTxFromContext retrieves transaction from context if it exists
func GetTxFromContext(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(TransactionalContext{}).(*gorm.DB)
	return tx, ok
}

// WithTransactionContext wraps a function with transaction handling using context
func WithTransactionContext(ctx context.Context, db *gorm.DB, fn func(ctx context.Context) error) error {
	// Check if we're already in a transaction
	if _, ok := GetTxFromContext(ctx); ok {
		// Already in transaction, just execute the function
		return fn(ctx)
	}

	// Start a new transaction
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create new context with transaction
	txCtx := context.WithValue(ctx, TransactionalContext{}, tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after rollback
		}
	}()

	if err := fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Transactional is a decorator for repository methods that need transaction support
func Transactional(db *gorm.DB) func(fn func(ctx context.Context) error) func(ctx context.Context) error {
	return func(fn func(ctx context.Context) error) func(ctx context.Context) error {
		return func(ctx context.Context) error {
			return WithTransactionContext(ctx, db, fn)
		}
	}
}
