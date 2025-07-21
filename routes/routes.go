package routes

import (
	"segment_service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	segment := r.Group("/segments")
	{
		segment.POST("", handlers.CreateSegment)
		segment.DELETE(":name", handlers.DeleteSegment)
		segment.POST(":name/assign", handlers.AssignSegment)
	}

	user := r.Group("/users")
	{
		user.GET(":id/segments", handlers.GetUserSegments)
	}
}
