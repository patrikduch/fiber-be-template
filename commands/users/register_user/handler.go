package register_user

import (
	"context"
	"strings"

	"fiber-be-template/database"
	"fiber-be-template/dtos/users/responses"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) (*responses.RegisterUserResponseDto, error) {
	id := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	normalizedEmail := strings.ToLower(strings.TrimSpace(cmd.Payload.Email))

	_, err = database.EntClient.User.
		Create().
		SetID(id).
		SetUsername(cmd.Payload.Name).
		SetEmail(cmd.Payload.Email).
		SetNormalizedEmail(normalizedEmail).
		SetPasswordHash(string(hashedPassword)).
		SetEmailConfirmed(false).
		SetPhoneNumberConfirmed(false).
		SetTwoFactorEnabled(false).
		SetLockoutEnabled(false).
		SetAccessFailedCount(0).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return &responses.RegisterUserResponseDto{ID: id.String()}, nil
}
