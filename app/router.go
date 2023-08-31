package app

import (
	controlController "controlUniversity/internal/controller/http/v1/control"
	"controlUniversity/internal/controller/http/v1/student"
	"github.com/gin-gonic/gin"
)

type Router struct {
	student student.Controller
	control controlController.Controller
}

func CreateRouter(student student.Controller, control controlController.Controller) Router {
	return Router{student: student, control: control}
}

func (r Router) StudentRouter(engine *gin.Engine) {
	students := engine.Group("/students")
	{
		students.POST("/", r.student.AddStudent)
		students.GET("/:id", r.student.GetDetail)
	}
}

func (r Router) ControlRouter(engine *gin.Engine) {
	controls := engine.Group("/controls")
	{
		controls.POST("/", r.control.CreateControl)
	}
}
