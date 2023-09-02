package main

import (
	"controlUniversity/app"
	controlController "controlUniversity/internal/controller/http/v1/control"
	studentController "controlUniversity/internal/controller/http/v1/student"
	userController "controlUniversity/internal/controller/http/v1/user"
	"controlUniversity/internal/pkg/postgresql"
	repoControl "controlUniversity/internal/repository/postgresql/control"
	repoStudent "controlUniversity/internal/repository/postgresql/student"
	repoUser "controlUniversity/internal/repository/postgresql/user"
	serviceControl "controlUniversity/internal/service/control"
	serviceStudent "controlUniversity/internal/service/student"
	serviceUser "controlUniversity/internal/service/user"
	useCaseControl "controlUniversity/internal/usecase/control"
	useCaseStudent "controlUniversity/internal/usecase/student"
	useCaseUser "controlUniversity/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	db := postgresql.ConnectPostgresql()

	repoUser := repoUser.RepositoryUser(db)
	repoStudent := repoStudent.RepositoryStudent(db)
	repoControl := repoControl.RepositoryControl(db)

	serviceUser := serviceUser.ServiceUser(repoUser)
	serviceStudent := serviceStudent.ServiceStudent(repoStudent)
	serviceControl := serviceControl.ServiceControl(repoControl)

	useCaseUser := useCaseUser.UserUseCase(serviceUser)
	useCaseStudent := useCaseStudent.StudentUseCase(serviceStudent, serviceControl, serviceUser)
	useCaseControl := useCaseControl.ControlUseCase(serviceControl)

	uController := userController.ControllerUser(useCaseUser)
	sController := studentController.ControllerStudent(useCaseStudent)
	cController := controlController.ControllerControl(useCaseControl)

	router := app.CreateRouter(uController, sController, cController)
	router.StudentRouter(server)
	router.ControlRouter(server)
	router.UserRouter(server)

	err := server.Run(":8000")

	if err != nil {
		log.Fatal(err)
	}
}
