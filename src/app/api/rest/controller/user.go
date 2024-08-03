package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/normatov07/mini-tweet/app"
	"github.com/normatov07/mini-tweet/common/response"
	"github.com/normatov07/mini-tweet/core/action"
	"github.com/normatov07/mini-tweet/core/service"
	"github.com/normatov07/mini-tweet/db/postgres"
)

type UserController struct {
	app *app.ApplicationContext
}

func GetUserController(app *app.ApplicationContext) UserController {
	return UserController{
		app: app,
	}
}

func (c *UserController) UserCreate(ctx *gin.Context) {
	var acn action.UserStore

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	token, err := service.GetUserService(new(postgres.UserRepo)).CreateUser(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(gin.H{"token": token}))
}

func (c *UserController) UserCreateFollow(ctx *gin.Context) {
	var acn action.StoreUserFollower

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	acn.FollowerID = c.app.User.ID

	err := service.GetUserService(new(postgres.UserRepo)).CreateUserFollower(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("OK"))
}

func (c *UserController) UserDeleteFollow(ctx *gin.Context) {
	var acn action.StoreUserFollower

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	acn.FollowerID = c.app.User.ID

	err := service.GetUserService(new(postgres.UserRepo)).DeleteUserFollower(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("OK"))
}
