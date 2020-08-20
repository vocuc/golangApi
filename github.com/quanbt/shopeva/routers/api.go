package routers

import (
	"shopeva/controllers/authencontroller"
	"shopeva/controllers/couponcontroller"
	"shopeva/controllers/productcontroler"

	"github.com/gin-gonic/gin"
)

//InitRoutes ...
func InitRoutes() {
	router := gin.Default()
	group := router.Group("/api/v1")

	//Product
	group.GET("/products", productcontroler.Products)
	group.POST("/products", productcontroler.Store)
	group.GET("/products/:id", productcontroler.FindProduct)
	group.PUT("/products/:id", productcontroler.Update)

	//Coupon
	group.GET("/coupons/:id", couponcontroller.ShowCoupon)
	group.POST("/coupons", couponcontroller.CreateaCoupon)

	//Login
	group.POST("/login", authencontroller.Login)

	router.Run()
}
