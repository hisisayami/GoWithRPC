package staff

import (
	"context"

	"example.com/go-inventory-grpc/ent"
	staffModel "example.com/go-inventory-grpc/internal/model"

	repo "example.com/go-inventory-grpc/internal/repository"
	"github.com/pkg/errors"
)

type Repository interface {
	StaffCre(ctx context.Context, newStaff staffModel.Staff) (*ent.Staff, error)
	GetAllStaff(ctx context.Context) ([]*ent.Staff, error)
	GetStaffByID(ctx context.Context, id int) (*ent.Staff, error)
	UpdateStaff(ctx context.Context, user ent.Staff) (*ent.Staff, error)
	DeleteStaffById(ctx context.Context, id int) error
	UpdateStaffById(ctx context.Context, staffId int, user staffModel.Staff) (*ent.Staff, error)
}

type repository struct {
	db repo.DB
}

func New(db repo.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllStaff(ctx context.Context) ([]*ent.Staff, error) {
	entC := r.db.GetEntClient()

	staffs, err := entC.Staff.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return staffs, nil
}

func (r *repository) GetStaffByID(ctx context.Context, id int) (*ent.Staff, error) {
	entC := r.db.GetEntClient()

	user, err := entC.Staff.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) StaffCre(ctx context.Context, newStaff staffModel.Staff) (*ent.Staff, error) {
	entC := r.db.GetEntClient()
	// tx, err := r.db.NewTransaction(ctx)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to get transaction from context")
	// }

	staffEntity, err := entC.Staff.Create().
		SetEmail(newStaff.Email).
		SetName(newStaff.Name).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create staff")
	}

	return staffEntity, nil
}

func (r *repository) UpdateStaff(ctx context.Context, user ent.Staff) (*ent.Staff, error) {
	entC := r.db.GetEntClient()

	updatedUser, err := entC.Staff.UpdateOneID(user.ID).
		SetEmail(user.Email).
		SetName(user.Name).Save(ctx)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *repository) DeleteStaffById(ctx context.Context, id int) error {
	entC := r.db.GetEntClient()

	err := entC.Staff.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateStaffById(ctx context.Context, staffId int, user staffModel.Staff) (*ent.Staff, error) {
	entC := r.db.GetEntClient()

	updateStaffById, err := entC.Staff.UpdateOneID(staffId).
		SetEmail(user.Email).
		SetName(user.Name).Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repo failed to update staff entity")
	}

	return updateStaffById, nil
}
