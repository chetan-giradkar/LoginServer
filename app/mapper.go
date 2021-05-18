package mapper

import (
	controller "LoginServer/controllers"
	"LoginServer/dao"
	services "LoginServer/serviceLayer"
	"LoginServer/store"

	"github.com/gin-gonic/gin"
)

func Map(ds *store.DataStore) {
	daoObject := dao.NewDaoStore(ds)
	serviceObject := services.NewService(*daoObject)
	contollerObject := controller.NewController(*serviceObject)

	ginRouter := gin.Default()

	ginRouter.POST("/signin", contollerObject.Signin)

	ginRouter.Run()
}
