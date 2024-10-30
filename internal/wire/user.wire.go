//go:build wireinject

package wire

import (
	"ecommerce/internal/controller"
	"ecommerce/internal/repo"
	"ecommerce/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
