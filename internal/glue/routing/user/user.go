package user

import (
	"helphub-backend/internal/glue/routing"
	"helphub-backend/internal/handler/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute(grp *gin.RouterGroup, user rest.User) {
	users := grp.Group("users")
	usersRoutes := []routing.Router{
		{
			Method:  http.MethodPost,
			Handler: user.CreateUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/:id",
			Handler: user.UpdateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/:id",
			Handler: user.GetUser,
		},
		{
			Method:  http.MethodGet,
			Handler: user.GetUsers,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/:id",
			Handler: user.DeleteUser,
		},
	}
	routing.RegisterRoutes(users, usersRoutes)
}
