package dto

import (
	"strings"

	"github.com/ShaizanKhan/go-banking-lib/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("deposit atleast 5000.00")
	}

	if strings.ToLower(r.AccountType) != "savings" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("account type should be checking or savings")
	}

	return nil
}
