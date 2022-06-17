package usecase

import (
	"context"
	"net/http"

	"kotak-amal/entity"
	"kotak-amal/helper"
	"kotak-amal/repository"
	"kotak-amal/transport/request"
	"kotak-amal/transport/response"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, req *request.RegisterUserRequest) (*entity.User, *response.GeneralResponse)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(
	userRepo repository.UserRepository,
) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, req *request.RegisterUserRequest) (*entity.User, *response.GeneralResponse) {
	now := helper.Now()

	/*
	 * Handle Passoword
	 */
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		errorResp := &response.GeneralResponse{
			Meta: response.Meta{
				Code:    http.StatusBadRequest,
				Status:  "failed",
				Message: err.Error(),
			},
		}
		return nil, errorResp
	}

	/*
	 * Mapping request to struct user
	 */
	userRequest := &entity.User{
		Name:           req.Name,
		Occupation:     req.Occupation,
		Email:          req.Email,
		PasswordHash:   string(passwordHash),
		AvatarFileName: "avatar.jpg",
		Role:           "user",
		Token:          "rahasia",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	/*
	 * Parse request to repository
	 */
	user, err := u.userRepo.Create(ctx, userRequest)
	if err != nil {
		errorResp := &response.GeneralResponse{
			Meta: response.Meta{
				Code:    http.StatusUnprocessableEntity,
				Status:  "failed",
				Message: "Un Processable Entity",
			},
			Data: nil,
		}
		return nil, errorResp
	}

	return user, nil
}
