package formatter

import (
	"time"

	"kotak-amal/entity"
)

type UserFormatter struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation"`
	Email      string    `json:"email"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatUser(u *entity.User) *UserFormatter {
	user := &UserFormatter{
		ID:         u.ID,
		Name:       u.Name,
		Occupation: u.Occupation,
		Email:      u.Email,
		Token:      u.Token,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}

	return user
}
