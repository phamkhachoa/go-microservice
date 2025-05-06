package dtos

type LoginInput struct {
	UserAccount  string `json:"user_account"`
	UserPassword string `json:"user_password"`
}

type LoginOutput struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
