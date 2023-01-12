package handler

import (
	"app/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func UserRegister(c *gin.Context) {
	var user model.UserModel

	if e := c.ShouldBind(&user); e != nil {
		log.Panicln("register failure", e.Error())
		return
	}

	//checkPassword := c.PostForm("password_check")
	//if checkPassword != user.Password {
	//	c.String(http.StatusBadRequest, "password is not equal to check password")
	//	log.Panicln("password is not to check password")
	//	return
	//}

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
		log.Println("login success", u.Email)
		c.JSON(http.StatusOK, gin.H{"login success": u.Email})
		return
	} else {
		log.Println("login failure", u.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"login failure": u.Email})
		return
	}
}

func UserGetAll(ctx *gin.Context) {
	user := model.UserModel{}
	users := user.FindAll()
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	intI, err := strconv.Atoi(id)
	if err != nil {
		log.Println("id is not int type", err.Error())
	}
	u := model.UserModel{
		Id: intI,
	}
	user := u.FindById()
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func DeleteOne(ctx *gin.Context) {
	id := ctx.Param("id")
	intI, err := strconv.Atoi(id)
	if err != nil {
		log.Println("id is not int type", err.Error())
	}
	u := model.UserModel{
		Id: intI,
	}
	u.DeleteById()

}
