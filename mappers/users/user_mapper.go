package users

import (
    "fiber-be-template/models"
    "fiber-be-template/dtos/users/requests"
    "fiber-be-template/dtos/users/responses"
    "github.com/google/uuid"
)

func ToUserResponseDto(user models.User) responses.UserResponseDto {
    return responses.UserResponseDto{
        ID:    user.ID.String(),
        Name:  user.Name,
        Email: user.Email,
    }
}

func ToUserModel(dto requests.CreateUserRequestDto) models.User {
    return models.User{
        ID:    uuid.New(), // Generate new UUID
        Name:  dto.Name,
        Email: dto.Email,
    }
}

func ToAuthMeResponseDto(user models.User) responses.AuthMeResponseDto {
	dto := responses.AuthMeResponseDto{
		ID:    user.ID.String(), // FIXED
		Email: user.Email,
		Name:  user.Name,
	}
	if user.Role != nil {
		dto.Role = user.Role.Name
	}
	return dto
}