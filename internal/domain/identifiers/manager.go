package identifiers

import (
	"context"

	"example.com/go-inventory-grpc/internal/domain/identifiers/creators"
)

type (
	validateNameFn func(ctx context.Context, name string) bool
)

// Interface to Create Staff
type Manager interface {
	CreateStaffByEmail(ctx context.Context) error
	CreateStaffByPhone(ctx context.Context) error
}

type manager struct {
	validateNameFn validateNameFn
	creators       creators.Creators
}

func New(config creators.Config) Manager {
	m := &manager{
		creators: creators.New(config),
	}
	m.validateNameFn = m.validateName
	return m
}

func (m *manager) CreateStaffByEmail(ctx context.Context) error {
	if m.validateNameFn(ctx, "name") {
		return nil
	}
	err := m.creators.Email.CreateStaffAccount(ctx)
	if err != nil {
		return nil
	}
	return nil
}

func (m *manager) CreateStaffByPhone(ctx context.Context) error {
	if m.validateNameFn(ctx, "name") {
		return nil
	}
	err := m.creators.Phone.CreateStaffAccount(ctx)
	if err != nil {
		return nil
	}
	return nil
}

func (m *manager) validateName(ctx context.Context, name string) bool {
	return true
}
