package staff

import (
	"context"

	"github.com/pkg/errors"

	staffClient "example.com/go-inventory-grpc/internal/service/staff"
)

// TODO: add CreateStaff func in interface and implement the CreateStaff func below like we did for Validate
// we would be calling this func in domain -> staffDomain since all business logic is added in staff domain
type StaffApi interface {
	Validate(ctx context.Context, fname string, lname string, sessionId string) (*staffClient.StaffInfoResponse, error)
}

type staffApi struct {
	staffClient staffClient.API
}

func New(staffClient staffClient.API) StaffApi {
	return &staffApi{
		staffClient: staffClient,
	}
}

func (s *staffApi) Validate(ctx context.Context, fname string, lname string, sessionId string) (*staffClient.StaffInfoResponse, error) {
	resp, err := s.staffClient.Validate(ctx, staffClient.StaffInfoInput{
		FirstName: fname,
		LastName:  lname,
		SessionID: sessionId,
	})
	if err != nil {
		return nil, errors.Wrap(err, "faile to validate staff")
	}

	return resp, nil
}
