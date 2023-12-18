package order

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Dbx *sqlx.DB
}

//func RegisterServiceOrder(router fiber.Router, db DB) {
//
//	var productRouter = router.Group("/v1/orders")
//	{
//		//productRouter.Post("/", middleware.AuthMiddleware(), handler.CheckoutOrder)
//		//productRouter.Get("/user", middleware.AuthMiddleware(), handler.GetHistoryOrderByUser)
//		//productRouter.Get("/merchant", middleware.AuthMiddleware(), handler.GetHistoryOrderByMerchant)
//		//productRouter.Post("/webhook",  handler.WebhookPaymentGateway)
//	}
//}
