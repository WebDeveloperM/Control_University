package student

import (
	"context"
	student2 "controlUniversity/internal/service/student"
	"controlUniversity/internal/usecase/student"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	student student.UseCase
}

func ControllerStudent(student student.UseCase) Controller {
	return Controller{student: student}
}

func (ct Controller) AddStudent(c *gin.Context) {
	ctx := context.Background()

	var addStudent student2.Create

	errBind := c.ShouldBind(&addStudent)
	if errBind != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errBind.Error(),
		})
		return
	}

	_, err := ct.student.AddStudent(ctx, addStudent)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": false,
		"student": addStudent,
	})

}

func (ct Controller) GetDetail(c *gin.Context) {
	pk := c.Param("id")
	id, errConv := strconv.Atoi(pk)
	if errConv != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": errConv.Error(),
		})
		return
	}
	ctx := context.Background()
	detail, err := ct.student.GetDetail(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": false,
		"data":    detail,
	})
}
