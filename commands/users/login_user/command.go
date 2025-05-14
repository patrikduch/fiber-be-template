package login_user

import "fiber-be-template/dtos/users/requests"

type Command struct {
    Payload requests.LoginUserRequestDto
}
