package handler

import (
	"app/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var user model.UserModel

	if e := c.ShouldBind(&user); e != nil {
		log.Panicln("register failure", e.Error())
		return
	}

	checkPassword := c.PostForm("password_check")
	if checkPassword != user.Password {
		c.String(http.StatusBadRequest, "password is not equal to check password")
		log.Panicln("password is not to check password")
		return
	}

	id := user.Store()
	log.Println("memberId is: ", id)
	c.String(http.StatusCreated, "memberId is: ", id)
	return
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
