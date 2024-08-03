package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/normatov07/mini-tweet/app/api/rest/controller"
	"github.com/normatov07/mini-tweet/app/api/rest/middleware"
)

func (s *Server) Routes() http.Handler {
	r := gin.Default()

	r.Use(middleware.AuthMiddleware(s.app))
	r.MaxMultipartMemory = 8 << 20 // 8 MB LIMIT

	postCtr := controller.GetPostController(s.app)
	usrCtr := controller.GetUserController(s.app)
	repostCtr := controller.GetRepostController(s.app)

	post := r.Group("/api/post")
	{
		post.GET("/", postCtr.GetPosts)
		post.GET("/follower-posts", postCtr.GetFollowerPosts)
		post.POST("/create", postCtr.StorePost)
		post.DELETE("/delete", postCtr.DeletePost)
		post.POST("/like/add", postCtr.AddLikePost)
		post.DELETE("/like/delete", postCtr.DelPostLike)
	}

	repost := r.Group("/api/repost")
	{
		repost.POST("/create", repostCtr.StoreRepost)
		repost.DELETE("/delete", repostCtr.DeleteRepost)
	}

	user := r.Group("/api/user")
	{
		user.POST("/create", usrCtr.UserCreate)
		user.POST("/follower/create", usrCtr.UserCreateFollow)
		user.DELETE("/follower/delete", usrCtr.UserDeleteFollow)
	}

	return r
}
