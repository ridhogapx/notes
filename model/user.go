package model

type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string
	Password string
	Name     string
	Role     uint8
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
