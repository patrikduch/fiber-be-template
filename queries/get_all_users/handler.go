package get_all_users

import (
    "context"
    "fmt"

    "fiber-be-template/database"
    "fiber-be-template/dtos/users/responses"
    "fiber-be-template/mappers/users"
    "fiber-be-template/models"
)

// Handler processes the GetAllUsers query.
type Handler struct{}

// NewHandler returns a new Handler instance.
func NewHandler() *Handler {
    return &Handler{}
}

// Handle executes the query and returns all users.
func (h *Handler) Handle(ctx context.Context, _ Query) ([]responses.UserResponseDto, error) {
    entUsers, err := database.EntClient.User.
        Query().
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
        result[i] = users.ToUserResponseDto(u)
    }

    return result, nil
}
