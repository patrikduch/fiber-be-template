package users

import (
    "fiber-be-template/models"
    "fiber-be-template/dtos/users/requests"
    "fiber-be-template/dtos/users/responses"
)

func ToUserResponseDto(user models.User) responses.UserResponseDto {
    return responses.UserResponseDto{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }
}

func ToUserModel(dto requests.CreateUserRequestDto, nextID int) models.User {
    return models.User{
        ID:    nextID,
        Name:  dto.Name,
        Email: dto.Email,
    }
}
