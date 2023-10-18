package usecase

import (
	"context"
	"time"

	"github.com/dilyara4949/clean-project/domain"
	"github.com/dilyara4949/clean-project/internal/tokenutil"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}



	func (lu *loginUsecase)GetUserByEmail(c context.Context, email string) (domain.User, error) {
		ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
		defer cancel()

		return lu.userRepository.GetByEmail(ctx, email)
	}

	func (lu *loginUsecase)CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
		return tokenutil.CreateAccessToken(user, secret, expiry)
	}

	func (lu *loginUsecase)CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
		return tokenutil.CreateRefreshToken(user, secret, expiry)
	}