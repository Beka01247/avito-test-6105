package routes

import (
	"zadanie-6105/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.GET("/api/ping", controllers.Ping)

    // tender routes
    router.GET("/api/tenders", controllers.GetTenders)
    router.POST("/api/tenders/new", controllers.CreateTender)
    router.GET("/api/tenders/my", controllers.GetUserTenders)
    router.GET("/api/tenders/:tenderId/status", controllers.GetTenderStatus)
    router.PUT("/api/tenders/:tenderId/status", controllers.UpdateTenderStatus)
    router.PATCH("/api/tenders/:tenderId/edit", controllers.EditTender)    
    router.PUT("/api/tenders/:tenderId/rollback/:version", controllers.RollbackTender)

    // bid routes
    // router.GET("/api/bids/:tenderId/list", controllers.GetBidsForTender)
    router.GET("/api/bids/my", controllers.GetUserBids)
    router.POST("/api/bids/new", controllers.CreateBid)
    router.GET("/api/bids/:bidId/status", controllers.GetBidStatus)
    router.PUT("/api/bids/:bidId/status", controllers.UpdateBidStatus)
    router.PATCH("/api/bids/:bidId/edit", controllers.EditBid)
    router.PUT("/api/bids/:bidId/submit_decision", controllers.SubmitBidDecision)
    router.PUT("/api/bids/:bidId/feedback", controllers.SubmitBidFeedback)
    // router.GET("/api/bids/:tenderId/reviews", controllers.GetBidReviews)
    router.PUT("/api/bids/:bidId/rollback/:version", controllers.RollbackBid)
}
