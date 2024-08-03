package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/normatov07/mini-tweet/app"
	"github.com/normatov07/mini-tweet/common/response"
	"github.com/normatov07/mini-tweet/core/action"
	"github.com/normatov07/mini-tweet/core/service"
	"github.com/normatov07/mini-tweet/db/postgres"
)

type PostController struct {
	app *app.ApplicationContext
}

func GetPostController(app *app.ApplicationContext) PostController {
	return PostController{
		app: app,
	}
}

func (c *PostController) StorePost(ctx *gin.Context) {
	var acn action.PostStore

	if err := ctx.Bind(&acn); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	acn.File, _ = ctx.FormFile("file")
	acn.UserID = c.app.User.ID

	err := service.GetPostService(new(postgres.PostRepo), new(postgres.ResourceRepo), new(postgres.UserRepo)).StorePost(acn)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("OK"))
}

func (c *PostController) DeletePost(ctx *gin.Context) {
	var acn action.RepostDelete

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	acn.UserID = c.app.User.ID

	err := service.GetPostService(new(postgres.PostRepo), new(postgres.ResourceRepo), new(postgres.UserRepo)).DeletePost(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("OK"))
}

func (c *PostController) AddLikePost(ctx *gin.Context) {
	var acn action.PostUser

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	acn.UserID = c.app.User.ID

	err := service.GetPostService(new(postgres.PostRepo), new(postgres.ResourceRepo), new(postgres.UserRepo)).AddLikePost(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("OK"))
}

func (c *PostController) DelPostLike(ctx *gin.Context) {
	var acn action.PostUser

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	acn.UserID = c.app.User.ID

	err := service.GetPostService(new(postgres.PostRepo), new(postgres.ResourceRepo), new(postgres.UserRepo)).DelPostLike(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("OK"))
}

func (c *PostController) GetPosts(ctx *gin.Context) {
	var acn action.PostPagination

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	resp, err := service.GetPostService(new(postgres.PostRepo), new(postgres.ResourceRepo), new(postgres.UserRepo)).GetPosts(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(resp))
}

func (c *PostController) GetFollowerPosts(ctx *gin.Context) {
	var acn action.PostPagination

	if err := ctx.Bind(&acn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	log.Println(c.app.User)
	acn.UserID = c.app.User.ID

	resp, err := service.GetPostService(new(postgres.PostRepo), new(postgres.ResourceRepo), new(postgres.UserRepo)).GetFollowerPosts(acn)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse(resp))
}
