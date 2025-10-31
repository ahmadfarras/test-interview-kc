package response

import (
	"reflect"
	"test-interview-kc/internal/domain"
	"testing"
	"time"
)

func TestToWalletAccountDetailResponse(t *testing.T) {
	timeNow := time.Now()
	tests := []struct {
		name  string
		input domain.WalletAccount
		want  WalletAccountDetailResponse
	}{
		{
			name: "all fields set",
			input: domain.WalletAccount{
				ID:        "acc1",
				Name:      "Test Account",
				Balance:   100.5,
				CreatedAt: &timeNow,
				CreatedBy: "admin",
				UpdatedAt: &timeNow,
				UpdatedBy: "admin",
			},
			want: WalletAccountDetailResponse{
				ID:        "acc1",
				Name:      "Test Account",
				Balance:   100.5,
				CreatedAt: &timeNow,
				CreatedBy: "admin",
				UpdatedAt: &timeNow,
				UpdatedBy: "admin",
			},
		},
		{
			name:  "zero values",
			input: domain.WalletAccount{},
			want:  WalletAccountDetailResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToWalletAccountDetailResponse(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToWalletAccountDetailResponse() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
