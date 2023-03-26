package staff

import (
	"context"
	"fmt"
	"testing"

	"example.com/go-inventory-grpc/ent"
	staffExternal "example.com/go-inventory-grpc/internal/domain/external/staff"
	terror "example.com/go-inventory-grpc/internal/error"
	staffModel "example.com/go-inventory-grpc/internal/model"
	staffRepository "example.com/go-inventory-grpc/internal/repository/staff"
	staffClient "example.com/go-inventory-grpc/internal/service/staff"
	"github.com/pkg/errors"
)

func Test_Approved(t *testing.T) {

	type input struct {
		ctx           context.Context
		input         staffModel.Staff
		staffRepo     staffRepository.Repository
		staffExternal staffExternal.StaffApi
	}

	type want struct {
		//status string
		err terror.Result
	}

	cases := []struct {
		name  string
		input input
		want  want
	}{
		{
			"handle approved case",
			input{
				ctx: context.Background(),
				input: staffModel.Staff{
					Id:    1,
					Name:  "Jhon Ena",
					Email: "eee@gmail.com",
				},
				staffRepo:     &mockStaffRepo{},
				staffExternal: &mockStaffExternal{},
			},
			want{
				err: terror.Result{
					WantErr: false,
				},
			},
		},
		{
			"handle error case",
			input{
				ctx: context.Background(),
				input: staffModel.Staff{
					Id:    1,
					Name:  "Jhon Ena",
					Email: "eee@gmail.com",
				},
				staffRepo: &mockStaffRepo{
					wantErr: true,
				},
				staffExternal: &mockStaffExternal{
					wantErr: true,
				},
			},
			want{
				err: terror.Result{
					WantErr:  true,
					Contains: "faile to validate staff",
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			staffDomain := New(tt.input.staffRepo, tt.input.staffExternal)

			res, err := staffDomain.StaffCre(tt.input.ctx, tt.input.input)

			fmt.Println("res", res)

			terror.AssertError(t, tt.want.err, err)

		})
	}

}

type mockStaffRepo struct {
	staffRepository.Repository

	wantErr bool
}

func (m *mockStaffRepo) StaffCre(ctx context.Context, newStaff staffModel.Staff) (*ent.Staff, error) {
	if m.wantErr {
		return nil, errors.New("failed to create")
	}
	ent := ent.Staff{
		ID:    1,
		Name:  "",
		Email: "",
	}

	return &ent, nil
}

type mockStaffExternal struct {
	staffExternal.StaffApi
	wantErr bool
}

func (mm *mockStaffExternal) Validate(ctx context.Context, fname string, lname string, sessionId string) (*staffClient.StaffInfoResponse, error) {
	if mm.wantErr {
		return nil, errors.New("faile to validate staff")
	}

	res := staffClient.StaffInfoResponse{
		Status:    "APPROVED",
		Comment:   "",
		SessionId: "122132",
	}

	return &res, nil
}
