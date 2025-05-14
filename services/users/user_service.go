package users

import (
    "context"
    "fmt"
    
    "github.com/google/uuid"
    
    "fiber-be-template/database"
    "fiber-be-template/ent"
    "fiber-be-template/ent/user"
    "fiber-be-template/dtos/users/responses"
    "fiber-be-template/mappers/users"
    "fiber-be-template/models"
)

// GetAllUsers retrieves all users from the DB
func GetAllUsers() ([]responses.UserResponseDto, error) {
    // Get the context
    ctx := context.Background()
    
    // Query all users from the database using the EntClient from database package
    entUsers, err := database.EntClient.User.
        Query().
        All(ctx)
    
    if err != nil {
        return nil, fmt.Errorf("failed querying users: %w", err)
    }
    
    // Convert to response DTOs
    result := make([]responses.UserResponseDto, len(entUsers))
    for i, entUser := range entUsers {
        // Create a models.User to match your original structure
        var u models.User
        u.ID = entUser.ID
        u.Name = entUser.Username
        u.Email = entUser.Email
        
        // Use the existing mapper function
        result[i] = users.ToUserResponseDto(u)
    }
    
    return result, nil
}

// GetUserByID retrieves one user by ID
func GetUserByID(id string) (*responses.UserResponseDto, error) {
    // Parse the string ID to UUID
    uid, err := uuid.Parse(id)
    if err != nil {
        return nil, err
    }
    
    // Create a models.User to match your original function
    var u models.User
    
    // Query using Ent with the EntClient from database package
    entUser, err := database.EntClient.User.
        Query().
        // Use the ID method (the exact method depends on your Ent schema)
        Where(user.ID(uid)).
        Only(context.Background())
    
    // Handle case where no user is found
    if ent.IsNotFound(err) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    
    // Map Ent user to your models.User
    u.ID = entUser.ID
    u.Name = entUser.Username
    u.Email = entUser.Email
    
    // Use the existing mapper function
    dto := users.ToUserResponseDto(u)
    return &dto, nil
}
