package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tucows-challenge/model"
)

func GetMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Menu)
}
