package http

import (
	"github.com/London57/todo-app/internal/controller/http/middleware"
	v1 "github.com/London57/todo-app/internal/controller/http/v1"
	"github.com/London57/todo-app/internal/infra/alias"
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, c *v1.V1, extra_d alias.Extra_data) {
	apiV1Group := app.Group("api/v1")

	auth := apiV1Group.Group("/auth")
	{
		auth.POST("/sign-up", c.Auth.SignUp)
		auth.POST("/sign-in", c.Auth.SignIn)
		auth.GET("/:provider/callback")
	}

	lists := apiV1Group.Group("/lists")
	lists.Use(middleware.JwtAuthMiddleware(extra_d["secret"]))
	{
		lists.POST("/", c.List.DeleteList)
		lists.GET("/", c.List.GetAllLists)
		lists.GET("/:id", c.List.GetListsById)
		lists.PUT("/:id", c.List.UpdateList)
		lists.DELETE("/:id", c.List.DeleteList)

		items := lists.Group(":id/items")
		{
			items.POST("/", c.Item.CreateItem)
			items.GET("/", c.Item.GetAllItems)
			items.GET("/:item_id", c.Item.GetAllItems)
			items.PUT("/:item_id", c.Item.UpdateItem)
			items.DELETE("/:item_id", c.Item.DeleteItem)
		}
	}
}
