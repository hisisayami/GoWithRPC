package staff

import (
	"context"
	"fmt"

	"example.com/go-inventory-grpc/ent"
	staffExternal "example.com/go-inventory-grpc/internal/domain/external/staff"
	staffManager "example.com/go-inventory-grpc/internal/domain/identifiers"
	staffModel "example.com/go-inventory-grpc/internal/model"
	staffRepository "example.com/go-inventory-grpc/internal/repository/staff"

	"github.com/pkg/errors"
)

const (
	REJECTED string = "REJECTED"
	APPROVED string = "APPROVED"
)

type StaffDomain interface {
	GetAllStaff(ctx context.Context) ([]*ent.Staff, error)
	GetStaffById(ctx context.Context, id int) (staffModel.Staff, error)
	DeleteStaffById(ctx context.Context, id int) error
	UpdateStaffById(ctx context.Context, staffId int, user staffModel.Staff) (*ent.Staff, error)
	StaffCre(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error)
}

type staffDomain struct {
	staffRepo     staffRepository.Repository
	staffExternal staffExternal.StaffApi
	staffManager  staffManager.Manager
}

func New(staffRepo staffRepository.Repository, staffExternal staffExternal.StaffApi, staffManager staffManager.Manager) StaffDomain {
	s := &staffDomain{
		staffRepo:     staffRepo,
		staffExternal: staffExternal,
		staffManager:  staffManager,
	}

	return s
}

// TODO: in rejected  case we need to call CreateStaff from domain -> external staff
func (s *staffDomain) StaffCre(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error) {

	res, err := s.staffExternal.Validate(ctx, staffDetails.Name, staffDetails.Name, "12313212")
	if err != nil {
		return staffModel.Staff{}, err
	}

	fmt.Println("res", res)

	switch res.Status {
	case REJECTED:
		fmt.Println("need to created staff by phone number")
		// TODO instead of returning staff object we need to call external staff to create the staff and phone number field is mandatory
		// call CreateStaffByPhone from staff manager
		var staffDBRes1 = staffModel.Staff{
			Id:    1,
			Name:  REJECTED,
			Email: REJECTED,
		}
		return staffDBRes1, nil
	case APPROVED:
		fmt.Println("staff can be crated by normal flow")
	default:
		// TODO we need to call external staff to create the staff and email field is mandatory
		// call CreateStaffByEmail from staff manager
	}

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
