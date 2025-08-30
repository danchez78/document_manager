package controllers

type RegisterUserRequest struct {
	RegisterUserController
}

type RegisterUserController struct {
	Token    string `json:"token"`
	Login    string `json:"login"`
	Password string `json:"pswd"`
}

type AuthUserRequest struct {
	AuthUserController
}

type AuthUserController struct {
	Login    string `json:"login"`
	Password string `json:"pswd"`
}

type DeauthUserRequest struct {
	Token string `param:"token"`
}
