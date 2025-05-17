package authctx

import (
	"context"
	"errors"
	"log"

	"fiber-be-template/database"
	"fiber-be-template/ent"
	entuser "fiber-be-template/ent/user"
	"fiber-be-template/models"
)

type ctxKey string

const (
	userIDKey ctxKey = "user_id"
	emailKey  ctxKey = "email"
)

// Inject user ID into context
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// Retrieve user ID from context
func UserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}

// Inject email into context
func WithUserEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, emailKey, email)
}

// Retrieve email from context
func UserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(emailKey).(string)
	return email, ok
}

func GetAuthenticatedUserWithRole(ctx context.Context) (*models.User, error) {
	email, ok := UserEmailFromContext(ctx)
	if !ok {
		log.Println("[AUTH] No email in context")
		return nil, errors.New("unauthorized")
	}

	log.Println("[AUTH] Looking up user by email:", email)

	entUser, err := database.EntClient.User.
		Query().
		Where(entuser.EmailEQ(email)).
		WithUserRoles(func(q *ent.UserRoleQuery) {
			q.WithRole()
		}).
		Only(ctx)

	if err != nil {
		log.Println("[AUTH] Failed to fetch user:", err)
		return nil, err
	}

	userModel := &models.User{
		ID:    entUser.ID,
		Name:  entUser.Username,
		Email: entUser.Email,
	}

	if len(entUser.Edges.UserRoles) > 0 && entUser.Edges.UserRoles[0].Edges.Role != nil {
		entRole := entUser.Edges.UserRoles[0].Edges.Role
		userModel.Role = &models.Role{
			ID:   entRole.ID,
			Name: entRole.Name,
		}
		log.Println("[AUTH] Role:", entRole.Name)
	} else {
		log.Println("[AUTH] No role attached to user")
	}

	return userModel, nil
}