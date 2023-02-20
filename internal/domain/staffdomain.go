package domain

import (
	"context"

	config "example.com/go-inventory-grpc/config"
	"example.com/go-inventory-grpc/ent"
	staffModel "example.com/go-inventory-grpc/internal/model"
	staffRepository "example.com/go-inventory-grpc/internal/repository/staff"
	"github.com/pkg/errors"
	//"google.golang.org/protobuf/internal/errors"
)

type StaffDomain interface {
	CreateStaff(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error)
	CreateStaff1(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error)
	GetAllStaff(ctx context.Context) ([]*ent.Staff, error)
	GetStaffById(ctx context.Context, id int) (staffModel.Staff, error)
	DeleteStaffById(ctx context.Context, id int) error
	UpdateStaffById(ctx context.Context, staffId int, user staffModel.Staff) (*ent.Staff, error)
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

// func StaffGetAll(ctx context.Context) ([]*ent.Staff, error) {
// 	staffs, err := staffRepository.New1(ctx).StaffGetAll()
// 	if err != nil {
// 		log.Printf("err : %s", err)
// 		return nil, err
// 	}
// 	return staffs, nil
// }

func (s *staffDomain) CreateStaff(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error) {

	staff, err := s.staffRepo.StaffCreate(ctx, staffDetails)
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

func (s *staffDomain) CreateStaff1(ctx context.Context, staffDetails staffModel.Staff) (staffModel.Staff, error) {
	entClient := config.GetClient()

	staff, err := staffRepository.New1(ctx, entClient).StaffCreate1(staffDetails)
	if err != nil {
		return staffModel.Staff{}, err
	}

	staffDBRes := staffModel.Staff{
		Id:    staff.ID,
		Name:  staff.Name,
		Email: staff.Email,
	}

	return staffDBRes, nil
}

// func (s *staffDomain) GetStaffById(ctx context.Context, id int) (staffModel.Staff, error) {

// 	getStaff, err := s.staffRepo.GetStaffByID(ctx, id)
// 	if err != nil {
// 		return staffModel.Staff{}, errors.Wrap(err, "failed to get staff by id")
// 	}

// 	return staffModel.Staff{
// 		Id:    getStaff.ID,
// 		Name:  getStaff.Name,
// 		Email: getStaff.Email,
// 	}, nil

// }

func (s *staffDomain) GetAllStaff(ctx context.Context) ([]*ent.Staff, error) {
	entClient := config.GetClient()
	staffRepo := staffRepository.New1(ctx, entClient)

	staffList, err := staffRepo.GetAllStaff()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get list of staffs")
	}

	return staffList, nil
}

func (s *staffDomain) GetStaffById(ctx context.Context, id int) (staffModel.Staff, error) {
	entClient := config.GetClient()

	getStaff, err := staffRepository.New1(ctx, entClient).GetStaffByID(ctx, id)
	if err != nil {
		return staffModel.Staff{}, errors.Wrap(err, "failed to get staff by id")
	}

	return staffModel.Staff{
		Name:  getStaff.Name,
		Email: getStaff.Email,
	}, nil
}

func (s *staffDomain) DeleteStaffById(ctx context.Context, id int) error {
	entClient := config.GetClient()

	err := staffRepository.New1(ctx, entClient).DeleteStaffById(ctx, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete staff by id")
	}

	return nil
}

func (s *staffDomain) UpdateStaffById(ctx context.Context, staffId int, user staffModel.Staff) (*ent.Staff, error) {
	entClient := config.GetClient()

	updateStaffById, err := staffRepository.New1(ctx, entClient).UpdateStaffById(ctx, staffId, user)
	if err != nil {
		return nil, errors.Wrap(err, "domain failed to update staff by id")
	}

	return updateStaffById, nil
}
