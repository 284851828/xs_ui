package xstock

import (
	v1 "xs/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (e *BaseRouter) InitRouter(priRouter *gin.RouterGroup, pubRouter *gin.RouterGroup) {

	pub := pubRouter.Group("xstock")
	api := v1.ApiGroupApp.XstockApiGroup

	pub.POST("getFutureClsList", api.GetFutureClsList) // 获取期货列表
	pub.POST("getPredList", api.GetPredList)           // 获取预测
	pub.POST("getPred0000List", api.GetPred0000List)   // 获取预测0000

	// customerRouter := Router.Group("xstock").Use(middleware.OperationRecord())
	// customerRouterWithoutRecord := Router.Group("xstock")
	// exaCustomerApi := v1.ApiGroupApp.ExampleApiGroup.CustomerApi
	// {
	// 	customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)   // 创建客户
	// 	customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)    // 更新客户
	// 	customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) // 删除客户
	// }
	// {
	// 	customerRouterWithoutRecord.GET("customer", exaCustomerApi.GetExaCustomer)         // 获取单一客户信息
	// 	customerRouterWithoutRecord.GET("customerList", exaCustomerApi.GetExaCustomerList) // 获取客户列表
	// }
}
