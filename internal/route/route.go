package route

import (
	"tinder-match/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRESTfulV1(
	api *gin.RouterGroup,
	ctrlV1 *controller.ControllerV1,
) {
	apiV1 := api.Group("/v1")
	{
		apiV1.POST("/person", ctrlV1.AddSinglePersonAndMatchHandler)
		apiV1.GET("/person", ctrlV1.QuerySinglePeopleHandler)
		apiV1.DELETE("/person/:name", ctrlV1.RemoveSinglePersonHandler)
	}
}
