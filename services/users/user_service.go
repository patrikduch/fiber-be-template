package users

import (
    "context"
    "fmt"
    
    "fiber-be-template/database"
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