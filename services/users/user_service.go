package users

import (

	"github.com/google/uuid"


	"database/sql"
	"fiber-be-template/database"
	"fiber-be-template/dtos/users/requests"
	"fiber-be-template/dtos/users/responses"
	"fiber-be-template/mappers/users"
	"fiber-be-template/models"
		  
)

// GetAllUsers retrieves all users from the DB
func GetAllUsers() ([]responses.UserResponseDto, error) {
	rows, err := database.DB.Query(`SELECT "Id", "UserName", "Email" FROM "User"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []responses.UserResponseDto
	for rows.Next() {
		var u models.User
		var id string // UUID is returned as string
		if err := rows.Scan(&id, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		u.ID, err = uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		result = append(result, users.ToUserResponseDto(u))
	}

	return result, nil
}

// GetUserByID retrieves one user by ID
func GetUserByID(id string) (*responses.UserResponseDto, error) {
	var u models.User
	err := database.DB.QueryRow(`SELECT Id, UserName, Email FROM "User" WHERE Id = $1`, id).
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

// CreateUser inserts a user and returns the result
func CreateUser(req requests.CreateUserRequestDto) (*responses.UserResponseDto, error) {
	var id string
	err := database.DB.QueryRow(
		`INSERT INTO "User"(Id, UserName, Email) VALUES($1, $2, $3) RETURNING Id`,
		uuid.New(), req.Name, req.Email,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &responses.UserResponseDto{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}
