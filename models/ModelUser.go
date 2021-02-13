package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Role	int `json:"role`
}

func (b *User) TableName() string {
	return "user"
}
