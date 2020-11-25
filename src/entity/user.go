package entity

// User is user models property
type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}
