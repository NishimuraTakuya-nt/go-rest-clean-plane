package response

import (
	"time"

	"github.com/NishimuraTakuya-nt/go-rest-clean-plane/internal/core/domain/models"
)

// UserResponse は基本的なユーザー情報を表す構造体です
// @Description User account information
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Roles     []string  `json:"roles"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToUserResponse はドメインモデルからレスポンスモデルへの変換を行います
func ToUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Roles: user.Roles,
		Email: user.Email,
	}
}

// ListUserResponse は複数ユーザー情報を返すためのレスポンス構造体です
// @Description User account list information
type ListUserResponse struct {
	Users      []UserResponse `json:"users"`
	TotalCount int            `json:"total_count"`
	Offset     int            `json:"offset"`
	Limit      int            `json:"limit"`
}

// ToListUserResponse は複数のユーザーモデルを変換します
func ToListUserResponse(models []models.User, offset, limit int) *ListUserResponse {
	users := make([]UserResponse, len(models))
	for i, model := range models {
		users[i] = ToUserResponse(&model)
	}

	return &ListUserResponse{
		Users:      users,
		TotalCount: len(users),
		Offset:     offset,
		Limit:      limit,
	}
}
