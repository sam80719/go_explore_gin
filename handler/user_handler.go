package handler

import (
	"github.com/gin-gonic/gin"
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
