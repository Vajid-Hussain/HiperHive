package router_server_svc

import (
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	handler_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/handler"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
)

func ServerRouter(engin *echo.Group, handler *handler_server_svc.ServerService, authMiddleWire *middlewire_auth_svc.Middlewire, soketioServer *socketio.Server) {
	engin.Use(authMiddleWire.UserAuthMiddlewire)
	{
		engin.POST("", handler.CreateServer)
		engin.GET("/:id", handler.GetServer)
		engin.POST("/join", handler.JoinToServer)
		engin.GET("/userserver", handler.GetUserServer)
		engin.GET("/message", handler.GetChannelMessage)
		engin.PATCH("/image", handler.UpdateServerPhoto)
		engin.PATCH("/description", handler.UpdateServerDescription)
		engin.GET("/members", handler.GetServerMembers)

		engin.PATCH("/role", handler.UpdateMemberRole)
		engin.DELETE("/remove", handler.RemoveUserFromServer)
		engin.DELETE("/left", handler.LeftFromServer)
		engin.DELETE("/destroy", handler.DeleteServer)

		engin.GET("/search", handler.SearchServer)

		engin.GET("/", func(ctx echo.Context) error {
			handler.InitSoketio(ctx)
			return echo.WrapHandler(soketioServer)(ctx)
		})

		categoryManagement := engin.Group("/category")
		{
			categoryManagement.POST("", handler.CreateCategory)
			categoryManagement.GET("", handler.GetCategoryOfServer)
		}

		channelManagement := engin.Group("/channel")
		{
			channelManagement.POST("", handler.CreateChannel)
			channelManagement.GET("", handler.GetChannelsOfServer)

			forumManagement := channelManagement.Group("/forum")
			{
				forumManagement.GET("", handler.GetForumPost)
				forumManagement.GET("/:postid", handler.GetSinglePost)
				forumManagement.GET("/command", handler.GetPostCommand)
			}
		}
	}

	engin.GET("/redis", handler.CheckReddis)
}
