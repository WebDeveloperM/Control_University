package app

import (
	controlController "controlUniversity/internal/controller/http/v1/control"
	"controlUniversity/internal/controller/http/v1/student"
	"controlUniversity/internal/controller/http/v1/user"
	token2 "controlUniversity/internal/utils/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type Router struct {
	user    user.Controller
	student student.Controller
	control controlController.Controller
}

func CreateRouter(user user.Controller, student student.Controller, control controlController.Controller) Router {
	return Router{user: user, student: student, control: control}
}

//func Middleware(role string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//		t := time.Now()
//		roleIncome, errRole := token2.ExtractTokenRole(c)
//		if role != roleIncome || errRole != nil {
//			log.Println(errRole, "err")
//			c.String(http.StatusUnauthorized, "Siz bu ishni qilolmaysiz to`g`ojon.")
//			c.Abort()
//			return
//		}
//		c.Next()
//		log.Println(time.Since(t))
//
//	}
//}

func Middleware(roleArr []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		roleOutcome, errRole := token2.ExtractTokenRole(c)
		var data bool
		for _, role := range roleArr {

			if strings.ToLower(role) == roleOutcome {
				data = true
				break
			} else {
				continue
			}
		}

		if !data || errRole != nil {
			log.Println(errRole, "err")
			c.String(http.StatusUnauthorized, "Siz bu ishni qilolmaysiz tog'ojon.")
			c.Abort()
			return

			c.Next()
			log.Println(time.Since(t))
		}

	}
}

func (r Router) UserRouter(engine *gin.Engine) {
	users := engine.Group("/users")
	{
		users.POST("/register", r.user.Register)
	}
}

func (r Router) StudentRouter(engine *gin.Engine) {
	students := engine.Group("/students")
	{
		students.POST("/", Middleware([]string{"Admin"}), r.student.AddStudent)      // Admin
		students.GET("/:id", Middleware([]string{"Superuser"}), r.student.GetDetail) // Superuser
	}
}

func (r Router) ControlRouter(engine *gin.Engine) {
	controls := engine.Group("/controls")
	{
		controls.POST("/", Middleware([]string{"Teacher"}), r.control.CreateControl) // Teacher
	}
}
