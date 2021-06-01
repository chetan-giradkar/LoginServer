package controller

import (
	service "LoginServer/ServiceLayer"
	"LoginServer/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Controller struct {
	Service service.LoginService
}

func NewController(svc *service.LoginService) *Controller {
	return &Controller{Service: *svc}
}

func (con *Controller) Signin(c *gin.Context) {
	var creds models.Credentials
	bindingError := c.ShouldBindBodyWith(&creds, binding.JSON)
	if bindingError != nil {
		log.Println("Error: Bad Request: ", bindingError)

		c.JSON(http.StatusBadRequest, bindingError.Error())
		c.Abort()
	}
	token, loginError := con.Service.Login(c, creds)
	if loginError != nil {
		c.JSON(http.StatusUnauthorized, loginError.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (con *Controller) Register(c *gin.Context) {
	var creds models.Credentials
	bindingError := c.ShouldBindBodyWith(&creds, binding.JSON)
	if bindingError != nil {
		log.Println("Error: Bad Request: ", bindingError)

		c.JSON(http.StatusBadRequest, bindingError.Error())
		c.Abort()
	}
	loginError := con.Service.Register(c, creds)
	if loginError != nil {
		c.JSON(http.StatusUnauthorized, loginError.Error())
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}
