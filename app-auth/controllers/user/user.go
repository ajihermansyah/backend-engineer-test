package user

import (
	"backend-engineer-test/app-auth/model"
	"backend-engineer-test/app-auth/repository"
	"encoding/json"

	httpHelper "backend-engineer-test/app-auth/helper/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Helper   httpHelper.HTTPHelper
	UserRepo repository.UserRepositoryInterface
}

func (uc *UserController) PingController(ctx *gin.Context) {
	ctx.JSON(200, "Server is working fine...")
}

func (uc *UserController) RegisterUserController(ctx *gin.Context) {
	var user model.User

	body, err := ctx.GetRawData()
	if err != nil {
		fmt.Println("Error read request body :", err)
		uc.Helper.SendError(ctx, 500, "Error reading request body", "error", uc.Helper.EmptyJsonMap())
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("Error unmarshalling json body :", err)
		uc.Helper.SendError(ctx, 500, "Error unmarshaling JSON body", "error", uc.Helper.EmptyJsonMap())
		return
	}

	psw, err := uc.UserRepo.CreateUser(user)
	if err != nil {
		fmt.Println("Error creating user :", err)
		uc.Helper.SendError(ctx, 500, err.Error(), "error", uc.Helper.EmptyJsonMap())
		return
	}

	user.Password = psw
	uc.Helper.SendSuccess(ctx, "Register user successfully", user)

}
