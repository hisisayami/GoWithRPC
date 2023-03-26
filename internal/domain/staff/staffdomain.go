package staff

import (
	"context"

	"example.com/go-inventory-grpc/ent"
	staffModel "example.com/go-inventory-grpc/internal/model"
	staffRepository "example.com/go-inventory-grpc/internal/repository/staff"
	"github.com/pkg/errors"
)

type StaffDomain interface {
	GetAllStaff(ctx context.Context) ([]*ent.Staff, error)
	GetStaffById(ctx context.Context, id int) (staffModel.Staff, error)
	DeleteStaffById(ctx context.Context, id int) error
	UpdateStaffById(ctx context.Context, staffId int, user staffModel.Staff) (*ent.Staff, error)
	StaffCre(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error)
}

type staffDomain struct {
	staffRepo staffRepository.Repository
}

func New(staffRepo staffRepository.Repository) StaffDomain {
	s := &staffDomain{
		staffRepo: staffRepo,
	}

	return s
}

func (s *staffDomain) StaffCre(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error) {

	staff, err := s.staffRepo.StaffCre(ctx, staffDetails)
	if err != nil {
		return staffModel.Staff{}, err
	}

	var staffDBRes = staffModel.Staff{
		Id:    staff.ID,
		Name:  staff.Name,
		Email: staff.Email,
	}

	return staffDBRes, nil

}

func (s *staffDomain) GetAllStaff(ctx context.Context) ([]*ent.Staff, error) {

	staffList, err := s.staffRepo.GetAllStaff(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get list of staffs")
	}

	return staffList, nil
}

func (s *staffDomain) GetStaffById(ctx context.Context, id int) (staffModel.Staff, error) {

	getStaff, err := s.staffRepo.GetStaffByID(ctx, id)
	if err != nil {
		return staffModel.Staff{}, errors.Wrap(err, "failed to get staff by id")
	}

	return staffModel.Staff{
		Name:  getStaff.Name,
		Email: getStaff.Email,
	}, nil
}

func (s *staffDomain) DeleteStaffById(ctx context.Context, id int) error {
	err := s.staffRepo.DeleteStaffById(ctx, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete staff by id")
	}

	return nil
}

func (s *staffDomain) UpdateStaffById(ctx context.Context, staffId int, user staffModel.Staff) (*ent.Staff, error) {
	updateStaffById, err := s.staffRepo.UpdateStaffById(ctx, staffId, user)
	if err != nil {
		return nil, errors.Wrap(err, "domain failed to update staff by id")
	}

	return updateStaffById, nil
}
