package get_authenticated_user

import (
	"context"
	"errors"
	"fmt"
	
	"fiber-be-template/database"
	"fiber-be-template/dtos/users/responses"
	"fiber-be-template/mappers/users"
	"fiber-be-template/models"
	"fiber-be-template/utils/authctx"
	"fiber-be-template/ent"
	"fiber-be-template/ent/user"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, _ Query) (*responses.AuthMeResponseDto, error) {
	// ✅ Extract email instead of user ID
	email, ok := authctx.UserEmailFromContext(ctx)
	if !ok || email == "" {
		return nil, errors.New("unauthorized")
	}

	entUser, err := database.EntClient.User.
		Query().
		Where(user.EmailEQ(email)).
		WithUserRoles(func(q *ent.UserRoleQuery) {
			q.WithRole()
		}).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed retrieving user: %w", err)
	}

	userModel := models.User{
		ID:    entUser.ID,
		Name:  entUser.Username,
		Email: entUser.Email,
	}

	if len(entUser.Edges.UserRoles) > 0 && entUser.Edges.UserRoles[0].Edges.Role != nil {
		role := entUser.Edges.UserRoles[0].Edges.Role
		userModel.Role = &models.Role{
			ID:   role.ID,
			Name: role.Name,
		}
	}

	dto := users.ToAuthMeResponseDto(userModel)
	return &dto, nil
}