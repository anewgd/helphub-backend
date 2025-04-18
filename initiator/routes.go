package initiator

import (
	"helphub-backend/docs"
	"helphub-backend/internal/glue/routing/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(group *gin.RouterGroup, handler Handler) {
	docs.SwaggerInfo.BasePath = "/v1"
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.InitRoute(group, handler.user)
}
