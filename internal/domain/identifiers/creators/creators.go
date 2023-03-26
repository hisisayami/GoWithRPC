package creators

import (
	"context"

	staffExternal "example.com/go-inventory-grpc/internal/domain/external/staff"
	"example.com/go-inventory-grpc/internal/domain/identifiers/creators/email"
	"example.com/go-inventory-grpc/internal/domain/identifiers/creators/phone"
)

type Creator interface {
	CreateStaffAccount(ctx context.Context) error
}

type Config struct {
	StaffExternal staffExternal.StaffApi
}

// added into single accessibale strucutre for phone and Email Verifier
type Creators struct {
	Phone Creator
	Email Creator
}

func New(cfg Config) Creators {
	return Creators{
		Phone: phone.New(phone.Config{
			StaffExternal: cfg.StaffExternal,
		}),
		Email: email.New(email.Config{
			StaffExternal: cfg.StaffExternal,
		}),
	}
}
