package mapper

import (
	service "LoginServer/ServiceLayer"
	controller "LoginServer/controllers"
	"LoginServer/dao"
	"LoginServer/store"

	"github.com/gin-gonic/gin"
)

func Map(ds *store.DataStore) {
	daoObject := dao.NewDaoStore(ds)
	serviceObject := service.NewService(*daoObject)
	contollerObject := controller.NewController(serviceObject)

	ginRouter := gin.Default()

	ginRouter.POST("/signin", contollerObject.Signin)

	ginRouter.Run()
}
