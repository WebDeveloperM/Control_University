package control

import (
	"context"
	control2 "controlUniversity/internal/service/control"
	"controlUniversity/internal/usecase/control"
	"controlUniversity/internal/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	control control.UseCase
}

func ControllerControl(control control.UseCase) Controller {
	return Controller{control: control}
}

func (ct Controller) CreateControl(c *gin.Context) {
	ctx := context.Background()

	var createControl control2.Create

	errBind := c.ShouldBind(&createControl)
	if errBind != nil {
		c.JSON(200, gin.H{
			"message": errBind.Error(),
		})
		return
	}

	id, errTk := token.ExtractTokenID(c)

	if errTk != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errTk.Error(),
			"status":  false,
			"data":    "",
		})
		return
	}
	createControl.UserId = int(id)

	_, err := ct.control.AddControl(ctx, createControl)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": false,
		"control": createControl,
	})

}
