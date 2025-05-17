package register_user

import (
	"context"
	"fmt"
	"log"
	"strings"

	"fiber-be-template/database"
	"fiber-be-template/dtos/users/responses"
	"fiber-be-template/ent/role"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) (*responses.RegisterUserResponseDto, error) {
	// Start a transaction with Ent
	tx, err := database.EntClient.Tx(ctx)
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return nil, fmt.Errorf("starting transaction: %w", err)
	}

	// Rollback in case of error
	rollback := true
	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	id := uuid.New()
	log.Println("Starting user registration with ID:", id.String())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, err
	}

	normalizedEmail := strings.ToLower(strings.TrimSpace(cmd.Payload.Email))
	normalizedUsername := strings.ToLower(strings.TrimSpace(cmd.Payload.Name))

	log.Println("Creating user...")
	// Create the user within the transaction
	user, err := tx.User.
		Create().
		SetID(id).
		SetUsername(cmd.Payload.Name).
		SetNormalizedUsername(normalizedUsername).
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
		log.Printf("Error creating user: %v", err)
		return nil, err
	}
	log.Printf("User created with ID: %s", user.ID.String())

	log.Println("Finding Customer role...")
	// Get the "Customer" role
	customerRole, err := tx.Role.
		Query().
		Where(role.NormalizedNameEQ("CUSTOMER")).
		Only(ctx)

	if err != nil {
		log.Printf("Error finding Customer role: %v", err)
		return nil, err
	}
	log.Printf("Found Customer role with ID: %s", customerRole.ID.String())

	// Try setting up the UserRole relationship within the transaction
	log.Println("Creating UserRole relationship...")
	
	// Try using the direct UserRole creation on the transaction
	_, err = tx.UserRole.
		Create().
		SetUserID(user.ID).
		SetRoleID(customerRole.ID).
		Save(ctx)
	
	if err != nil {
		log.Printf("Error creating UserRole relationship: %v", err)
		return nil, err
	}
	
	// If we reached here, we can commit the transaction
	rollback = false
	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		return nil, fmt.Errorf("committing transaction: %w", err)
	}
	
	log.Println("User registration completed successfully")
	return &responses.RegisterUserResponseDto{ID: id.String()}, nil
}