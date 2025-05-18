package vo

type UserRegistratorRequest struct {
	Email    string `json:"email"`
	Purpose  string `json:"purpose"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
