package http

import (
	v1 "github.com/London57/todo-app/internal/controller/http/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, c *v1.V1) { // TODO: import usecases
	apiV1Group := app.Group("api/v1")

	auth := apiV1Group.Group("/auth")
	{

		auth.POST("sign-up", c.Auth.SignUp)
		auth.POST("sign-in", c.Auth.SignIn)
	}

	lists := apiV1Group.Group("/lists")
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
