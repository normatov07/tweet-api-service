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

type RepostController struct {
	app *app.ApplicationContext
}

func GetRepostController(app *app.ApplicationContext) RepostController {
	return RepostController{
		app: app,
	}
}

func (c *RepostController) StoreRepost(ctx *gin.Context) {
	var acn action.RepostCreate

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	acn.UserID = c.app.User.ID

	err := service.GetRepostService(new(postgres.PostRepo)).StoreRepost(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("OK"))
}

func (c *RepostController) DeleteRepost(ctx *gin.Context) {
	var acn action.RepostDelete

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	acn.UserID = c.app.User.ID

	err := service.GetRepostService(new(postgres.PostRepo)).DeleteRepost(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("OK"))
}
