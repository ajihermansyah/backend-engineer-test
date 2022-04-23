package user

import (
	"backend-engineer-test/app-auth/model"
	"backend-engineer-test/app-auth/repository"
	"encoding/json"

	respModel "backend-engineer-test/app-auth/model/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepo repository.UserRepositoryInterface
}

func (h *UserController) PingController(ctx *gin.Context) {
	ctx.JSON(200, "Server is working fine...")
}

func (h *UserController) RegisterUserController(ctx *gin.Context) {
	var user model.User

	body, err := ctx.GetRawData()
	if err != nil {
		fmt.Println("Error read request body :", err)
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: "Error reading request body",
		})

		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("Error unmarshalling json body :", err)
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: "Error unmarshaling JSON body",
		})

		return
	}

	psw, err := h.UserRepo.CreateUser(user)
	if err != nil {
		fmt.Println("Error creating user :", err)
		ctx.JSON(500, respModel.FailedResponse{
			Code:    500,
			Status:  "failed",
			Message: err.Error(),
		})

		return
	}

	user.Password = psw

	ctx.JSON(201, respModel.SuccessResponse{
		Code:    201,
		Status:  "success",
		Message: "Register user successfuly",
		Data:    user.Password,
	})

}
