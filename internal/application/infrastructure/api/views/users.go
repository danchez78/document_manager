package views

type RegisterUserResponse struct {
	Login string `json:"login"`
}

func NewRegisterUserResponse(login string) RegisterUserResponse {
	return RegisterUserResponse{Login: login}
}

type AuthUserResponse struct {
	Token string `json:"token"`
}

func NewAuthUserResponse(token string) AuthUserResponse {
	return AuthUserResponse{Token: token}
}

type DeauthUserResponse map[string]bool

func NewDeauthUserResponse(token string) DeauthUserResponse {
	return DeauthUserResponse{
		token: true,
	}
}
