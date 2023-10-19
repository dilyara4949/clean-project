package route

import (
	"time"

	"github.com/dilyara4949/clean-project/api/controller"
	"github.com/dilyara4949/clean-project/bootstrap"
	"github.com/dilyara4949/clean-project/domain"
	"github.com/dilyara4949/clean-project/mongo"
	"github.com/dilyara4949/clean-project/repository"
	"github.com/dilyara4949/clean-project/usecase"
	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}