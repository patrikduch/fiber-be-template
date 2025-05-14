package register_user

import "fiber-be-template/dtos/users/requests"

type Command struct {
    Payload requests.RegisterUserRequestDto
}
