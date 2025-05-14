package get_user_by_email

import (
    "context"
    "fiber-be-template/database"
    "fiber-be-template/dtos/users/responses"
    "fiber-be-template/mappers/users"
    "fiber-be-template/models"
    "fiber-be-template/ent"
    "fiber-be-template/ent/user"
)

type Handler struct{}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, q Query) (*responses.UserResponseDto, error) {
    entUser, err := database.EntClient.User.
        Query().
        Where(user.EmailEQ(q.Email)).
        Only(ctx)

    if ent.IsNotFound(err) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }

    u := models.User{
        ID:    entUser.ID,
        Name:  entUser.Username,
        Email: entUser.Email,
    }

    dto := users.ToUserResponseDto(u)
    return &dto, nil
}
