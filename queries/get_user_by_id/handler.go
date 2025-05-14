package get_user_by_id

import (
    "context"
    "fmt"

    "github.com/google/uuid"
    "fiber-be-template/database"
    "fiber-be-template/dtos/users/responses"
    "fiber-be-template/ent"
    "fiber-be-template/ent/user"
    "fiber-be-template/mappers/users"
    "fiber-be-template/models"
)

type Handler struct{}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, q Query) (*responses.UserResponseDto, error) {
    uid, err := uuid.Parse(q.ID)
    if err != nil {
        return nil, err
    }

    entUser, err := database.EntClient.User.
        Query().
        Where(user.ID(uid)).
        Only(ctx)

    if ent.IsNotFound(err) {
        return nil, nil
    }
    if err != nil {
        return nil, fmt.Errorf("failed retrieving user by ID: %w", err)
    }

    u := models.User{
        ID:    entUser.ID,
        Name:  entUser.Username,
        Email: entUser.Email,
    }

    dto := users.ToUserResponseDto(u)
    return &dto, nil
}
