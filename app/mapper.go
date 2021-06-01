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
	jwtService := service.NewJwtService()
	serviceObject := service.NewService(*daoObject, *jwtService)
	contollerObject := controller.NewController(serviceObject)

	ginRouter := gin.Default()

	ginRouter.POST("/signin", contollerObject.Signin)
	ginRouter.POST("/register", contollerObject.Register)

	ginRouter.Run()
}
