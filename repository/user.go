package repository

import (
	"context"
	"database/sql"

	"kotak-amal/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}

	res, err := tx.ExecContext(ctx, `INSERT INTO users (
		name,
		occupation,
		email,
		password_hash,
		avatar_file_name,
		role,
		token,
		created_at,
		updated_at) Values (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.Name,
		user.Occupation,
		user.Email,
		user.PasswordHash,
		user.AvatarFileName,
		user.Role,
		user.Token,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	user.ID = int(id)
	return user, nil
}
