package service

import (
	"ashishi-banking/domain"
	"ashishi-banking/dto"
	"ashishi-banking/errs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_customerService_GetAllCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		status string
	}

	//arrange
	tests := []struct {
		name    string
		repo    *domain.MockCustomerRepository
		args    args
		want    []dto.CustomerResponse
		wantErr *errs.AppError
	}{
		{
			name: "should return slice customers and nil error",
			repo: func(mock *domain.MockCustomerRepository) *domain.MockCustomerRepository {
				mock.EXPECT().FindAll("1").Return([]domain.Customer{
					{
						Id:          "1",
						Name:        "name",
						City:        "city",
						Zipcode:     "zipcode",
						DateofBirth: "2001-01-01 11:04:05",
						Status:      "1",
					},
					{
						Id:          "3",
						Name:        "name",
						City:        "city",
						Zipcode:     "zipcode",
						DateofBirth: "2003-01-01 11:04:05",
						Status:      "1",
					},
					{Id: "4",
						Name:        "name4",
						City:        "city4",
						Zipcode:     "zipcode4",
						DateofBirth: "2004-01-01 11:04:05",
						Status:      "1"},
				}, nil)
				return mock
			}(domain.NewMockCustomerRepository(ctrl)),
			args: args{status: "active"},
			want: []dto.CustomerResponse{
				{
					Id:          "1",
					Name:        "name",
					City:        "city",
					Zipcode:     "zipcode",
					DateofBirth: "2001-01-01 11:04:05",
					Status:      "active",
				},
				{
					Id:          "3",
					Name:        "name",
					City:        "city",
					Zipcode:     "zipcode",
					DateofBirth: "2003-01-01 11:04:05",
					Status:      "active",
				},
				{Id: "4",
					Name:        "name4",
					City:        "city4",
					Zipcode:     "zipcode4",
					DateofBirth: "2004-01-01 11:04:05",
					Status:      "active"},
			},
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {
			t := &DefaultCustomerService{repo: test.repo}

			customers, err := t.GetAllCustomer(test.args.status)
			if err != nil {
				assert.Equal(t1, err, test.wantErr)
				return
			} else {
				assert.Equal(t1, customers, test.want)
			}
		})
	}
}
