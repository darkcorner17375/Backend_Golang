package dto

//POST from /login url when client loging
type LoginDTO struct {
	Email    string `json:"email" from:"email" binding:"required,email"`
	Password string `json:"password" from:"password" binding:"required,min=6"`
}
