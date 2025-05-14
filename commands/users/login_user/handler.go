package login_user

import (
    "context"
    "errors"
    "os"
    "strings"
    "time"

    "fiber-be-template/database"
    "fiber-be-template/dtos/users/responses"
    "fiber-be-template/ent"
    "fiber-be-template/ent/user"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

type Handler struct{}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) (*responses.LoginUserResponseDto, error) {
    normalizedEmail := strings.ToLower(strings.TrimSpace(cmd.Payload.Email))

    u, err := database.EntClient.User.
        Query().
        Where(user.NormalizedEmailEQ(normalizedEmail)).
        Only(ctx)

    if ent.IsNotFound(err) || err != nil {
        return nil, errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(cmd.Payload.Password)); err != nil {
        return nil, errors.New("invalid credentials")
    }

    token, err := generateJWT(u.ID.String(), u.Email)
    if err != nil {
        return nil, err
    }

    return &responses.LoginUserResponseDto{AccessToken: token}, nil
}

func generateJWT(userID string, email string) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return "", errors.New("JWT_SECRET not set")
    }

    claims := jwt.MapClaims{
        "sub":   userID,
        "email": email,
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
        "iat":   time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}
