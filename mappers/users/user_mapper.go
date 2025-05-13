package users

import (
    "fiber-be-template/models"
    "fiber-be-template/dtos/users/requests"
    "fiber-be-template/dtos/users/responses"
    "github.com/google/uuid"
)

// Convert User model to response DTO
func ToUserResponseDto(user models.User) responses.UserResponseDto {
    return responses.UserResponseDto{
        ID:    user.ID.String(), // Convert UUID to string
        Name:  user.Name,
        Email: user.Email,
    }
}

// Convert CreateUserRequestDto to User model with generated UUID
func ToUserModel(dto requests.CreateUserRequestDto) models.User {
    return models.User{
        ID:    uuid.New(), // Generate new UUID
        Name:  dto.Name,
        Email: dto.Email,
    }
}
