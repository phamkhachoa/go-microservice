package request

type UserFilter struct {
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}
