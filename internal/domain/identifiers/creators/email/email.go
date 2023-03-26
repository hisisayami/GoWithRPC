package email

import (
	"context"

	staffExternal "example.com/go-inventory-grpc/internal/domain/external/staff"
)

type Config struct {
	StaffExternal staffExternal.StaffApi
}

type Creator struct {
	staffExternal staffExternal.StaffApi
}

func New(cfg Config) *Creator {
	return &Creator{
		staffExternal: cfg.StaffExternal,
	}
}

func (v *Creator) CreateStaffAccount(ctx context.Context) error {
	//TODO : call /create staff from staffExternal.StaffApi

	return nil
}
