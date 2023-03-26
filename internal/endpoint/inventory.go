package endpoint

import (
	"log"

	staffDomain "example.com/go-inventory-grpc/internal/domain/staff"
	staffModel "example.com/go-inventory-grpc/internal/model"
	repo "example.com/go-inventory-grpc/internal/repository"

	"github.com/pkg/errors"

	"golang.org/x/net/context"
)

type server struct {
	staffD staffDomain.StaffDomain
	db     repo.DB
	UnimplementedInventoryServiceServer
}

type Config struct {
	StaffD staffDomain.StaffDomain
	DB     repo.DB
}

func (s *server) Register(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{
		Body: "hello from the server!",
	}, nil
}

func New(cfg Config) InventoryServiceServer {
	s := &server{
		staffD: cfg.StaffD,
		db:     cfg.DB,
	}
	return s
}

// func (s *Server) CreateStaff(ctx context.Context, message *CreateStaffRequest) (*CreateStaffResponse, error) {
// 	log.Printf("Received message body from client: %s", message)
// 	staffrepo := staffRepo.New()
// 	ss := staffDomain.New(staffrepo)
// 	staffInfo, err := ss.CreateStaff(ctx, staffModel.Staff{
// 		Name:  message.Name,
// 		EMAIL: message.Email,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &CreateStaffResponse{
// 		Id:    int32(staffInfo.ID),
// 		Name:  staffInfo.Name,
// 		Email: staffInfo.EMAIL,
// 	}, nil
// }

func (s *server) CreateStaff(ctx context.Context, message *CreateStaffRequest) (*CreateStaffResponse, error) {
	log.Printf("Received message body from client: %s", message)
	staffInfo, err := s.staffD.StaffCre(ctx, staffModel.Staff{
		Name:  message.Name,
		Email: message.Email,
	})
	if err != nil {
		return nil, err
	}

	return &CreateStaffResponse{
		Id:    int32(staffInfo.Id),
		Name:  staffInfo.Name,
		Email: staffInfo.Email,
	}, nil
}

func (s *server) GetStaffById(ctx context.Context, message *GetStaffByIdRequest) (*GetStaffByIdResponse, error) {
	log.Printf("Get staff request: %s", message)

	getStaff, err := s.staffD.GetStaffById(ctx, int(message.Id))
	if err != nil {
		return &GetStaffByIdResponse{}, err
	}

	return &GetStaffByIdResponse{
		Id:    int32(getStaff.Id),
		Name:  getStaff.Name,
		Email: getStaff.Email,
	}, nil
}

func (s *server) DeleteStaffById(ctx context.Context, message *DeleteStaffByIdRequest) (*DeleteStaffByIdResponse, error) {

	err := s.staffD.DeleteStaffById(ctx, int(message.Id))
	if err != nil {
		return &DeleteStaffByIdResponse{}, err
	}

	return &DeleteStaffByIdResponse{}, nil
}

func (s *server) UpdateStaffById(ctx context.Context, message *UpdateStaffByIdRequest) (*UpdateStaffByIdResponse, error) {

	updatedStaff, err := s.staffD.UpdateStaffById(ctx, int(message.Id), staffModel.Staff{
		Name:  message.Name,
		Email: message.Email,
	})
	if err != nil {
		return nil, errors.Wrap(err, "endpoint failed to update staff by id")
	}

	return &UpdateStaffByIdResponse{
		Id:    int32(updatedStaff.ID),
		Name:  updatedStaff.Name,
		Email: updatedStaff.Email,
	}, nil

}

func (s *server) GetAllStaff(ctx context.Context, message *GetAllStaffRequest) (*GetAllStaffResponse, error) {

	getAllStaff, err := s.staffD.GetAllStaff(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get list of staffs")
	}
	list := []*Staff{}
	for _, l := range getAllStaff {
		staff := Staff{
			Id:    int32(l.ID),
			Name:  l.Name,
			Email: l.Email,
		}
		list = append(list, &staff)

	}

	return &GetAllStaffResponse{
		StaffList: list,
	}, nil
}

// func (s *server) CreateStaff(ctx context.Context, message *CreateStaffRequest) (*CreateStaffResponse, error) {
// 	log.Printf("Received message body from client: %s", message)
// 	staffrepo := staffRepo.New()
// 	staffDomain := staffDomain.New(staffrepo)
// 	staffInfo, err := staffDomain.CreateStaff1(ctx, staffModel.Staff{
// 		Name:  message.Name,
// 		Email: message.Email,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &CreateStaffResponse{
// 		Id:    int32(staffInfo.Id),
// 		Name:  staffInfo.Name,
// 		Email: staffInfo.Email,
// 	}, nil
// }
