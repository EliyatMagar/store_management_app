package routes

import (
	"store-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterReviewRoutes(r *gin.Engine) {
	reviewRoutes := r.Group("/api/reviews")
	{
		reviewRoutes.POST("/", controllers.CreateReview)
		reviewRoutes.GET("/", controllers.GetAllReviews)
		reviewRoutes.GET("/:id", controllers.GetReviewByID)
		reviewRoutes.PUT("/:id", controllers.UpdateReview)
		reviewRoutes.DELETE("/:id", controllers.DeleteReview)
	}
}
