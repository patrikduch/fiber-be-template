package get_all_users

import (
	"context"
	"fmt"

	"fiber-be-template/database"
	"fiber-be-template/dtos/users/responses"
	"fiber-be-template/mappers/users"
	"fiber-be-template/models"
	"fiber-be-template/ent"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, _ Query) ([]responses.UserResponseDto, error) {
	entUsers, err := database.EntClient.User.
		Query().
		WithUserRoles(func(q *ent.UserRoleQuery) {
			q.WithRole()
		}).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying users: %w", err)
	}

	result := make([]responses.UserResponseDto, len(entUsers))
	for i, entUser := range entUsers {
		u := models.User{
			ID:    entUser.ID,
			Name:  entUser.Username,
			Email: entUser.Email,
		}

		// Check for UserRoles edge
		if len(entUser.Edges.UserRoles) > 0 && entUser.Edges.UserRoles[0].Edges.Role != nil {
			role := entUser.Edges.UserRoles[0].Edges.Role
			u.Role = &models.Role{
				ID:   role.ID,
				Name: role.Name,
			}
		}

		result[i] = users.ToUserResponseDto(u)
	}

	return result, nil
}