package users

import (
    "database/sql"
    "fiber-be-template/database"
    "fiber-be-template/models"
    "fiber-be-template/dtos/users/requests"
    "fiber-be-template/dtos/users/responses"
    "fiber-be-template/mappers/users"
)

func GetAllUsers() ([]responses.UserResponseDto, error) {
    rows, err := database.DB.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var result []responses.UserResponseDto
    for rows.Next() {
        var u models.User
        if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
            return nil, err
        }
        result = append(result, users.ToUserResponseDto(u))
    }

    return result, nil
}

func GetUserByID(id int) (*responses.UserResponseDto, error) {
    var u models.User
    err := database.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).
        Scan(&u.ID, &u.Name, &u.Email)

    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }

    dto := users.ToUserResponseDto(u)
    return &dto, nil
}

func CreateUser(req requests.CreateUserRequestDto) (*responses.UserResponseDto, error) {
    var id int
    err := database.DB.QueryRow(
        "INSERT INTO users(name, email) VALUES($1, $2) RETURNING id",
        req.Name, req.Email,
    ).Scan(&id)

    if err != nil {
        return nil, err
    }

    dto := responses.UserResponseDto{
        ID:    id,
        Name:  req.Name,
        Email: req.Email,
    }
    return &dto, nil
}
