package factory

import "github.com/joaoeduardodev/golang"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}