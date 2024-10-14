package usecases

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/domain/models"
	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/mocks/mockpiyographql"
	"github.com/golang/mock/gomock"
)

func Test_lenID(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test lenID",
			args: args{
				ID: "123",
			},
			want: 3,
		},
		{
			name: "Test lenID 10",
			args: args{
				ID: "1234567890",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenID(tt.args.ID); got != tt.want {
				t.Errorf("lenID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockpiyographql.NewMockClient(ctrl)
	target := NewUserUseCase(mockClient)

	t.Run("get user", func(t *testing.T) {
		ID := "123"
		// モックの振る舞いを設定
		mockClient.EXPECT().GetUser(context.Background(), ID).
			Return(&models.User{ID: "123", Name: "Test User"}, nil)

		// テストケースを実行
		user, err := target.Get(context.Background(), ID)

		assert.NoError(t, err)
		assert.Equal(t, "123", user.ID)
	})

	t.Run("get user 2", func(t *testing.T) {
		ID := "aaa"
		// モックの振る舞いを設定
		mockClient.EXPECT().GetUser(context.Background(), ID).
			Return(&models.User{ID: "aaa", Name: "Test User"}, nil)

		// テストケースを実行
		user, err := target.Get(context.Background(), ID)

		assert.NoError(t, err)
		assert.Equal(t, "aaa", user.ID)
	})
}
