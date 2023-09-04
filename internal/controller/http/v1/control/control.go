package control

import (
	"context"
	control2 "controlUniversity/internal/service/control"
	"controlUniversity/internal/usecase/control"
	"controlUniversity/internal/utils/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	fmt.Println(id, "ishladikkkkkkkkkkkkkkkiuuuuuuuuuuuuuuuuuu")
	if errTk != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errTk.Error(),
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

func (ct Controller) DeleteControl(c *gin.Context) {
	pk := c.Param("id")
	id, errConv := strconv.Atoi(pk)
	if errConv != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	_, err := ct.control.DeleteControl(c, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    false,
		"is_deleted": true,
	})
}

func (ct Controller) UpdateControl(c *gin.Context) {
	pk := c.Param("id")
	id, errConv := strconv.Atoi(pk)
	if errConv != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	var updateControl control2.Update

	errBind := c.ShouldBind(&updateControl)
	if errBind != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errBind.Error(),
		})
		return
	}

	control, err := ct.control.UpadateControl(c, updateControl, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    false,
		"is_updated": true,
		"control":    control,
	})
}
