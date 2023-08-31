package control

import (
	"context"
	control2 "controlUniversity/internal/service/control"
	"controlUniversity/internal/usecase/control"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Controller struct {
	control control.UseCase
}

func ControllerControl(control control.UseCase) Controller {
	return Controller{control: control}
}

func (ct Controller) CreateControl(c *gin.Context) {
	ctx := context.Background()
	fmt.Println(os.Getenv("DBNAME"))
	var createControl control2.Create

	errBind := c.ShouldBind(&createControl)
	if errBind != nil {
		c.JSON(200, gin.H{
			"message": errBind.Error(),
		})
		return
	}

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
