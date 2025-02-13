package store

import "github.com/google/uuid"


type TransactionModelInterface interface {

}

type Transaction struct {
    ID uuid.UUID
    Description string
    ValueInCents int
}
