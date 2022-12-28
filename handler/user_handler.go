package handler

import (
	"app/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserSave(c *gin.Context) {
	userName := c.Param("name")
	c.String(http.StatusOK, "save data for user: "+userName)
}

func UserSaveBYQuery(c *gin.Context) {
	userName := c.Query("name")
	phone := c.DefaultQuery("phone", "95279527")
	c.String(http.StatusOK, "save data for user: "+userName+", phone: "+phone)
}

func UserLogin(c *gin.Context) {
	var user model.UserModel

	if e := c.Bind(&user); e != nil {
		log.Panicln("login failure", e.Error())
		return
	}

	u := user.QueryByAccount()
	if u.Password == user.Password {
		log.Println("login success", u.Account)
		c.JSON(http.StatusOK, gin.H{"login success": u.Account})
		return
	} else {
		log.Println("login failure", u.Account)
		c.JSON(http.StatusUnauthorized, gin.H{"login failure": u.Account})
		return
	}
}
