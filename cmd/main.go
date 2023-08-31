package main

import (
	"controlUniversity/app"
	controlController "controlUniversity/internal/controller/http/v1/control"
	studentController "controlUniversity/internal/controller/http/v1/student"
	"controlUniversity/internal/pkg/postgresql"
	repoControl "controlUniversity/internal/repository/postgresql/control"
	repoStudent "controlUniversity/internal/repository/postgresql/student"
	serviceControl "controlUniversity/internal/service/control"
	serviceStudent "controlUniversity/internal/service/student"
	useCaseControl "controlUniversity/internal/usecase/control"
	useCaseStudent "controlUniversity/internal/usecase/student"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	db := postgresql.ConnectPostgresql()

	repoStudent := repoStudent.RepositoryStudent(db)
	repoControl := repoControl.RepositoryControl(db)

	serviceStudent := serviceStudent.ServiceStudent(repoStudent)
	serviceControl := serviceControl.ServiceControl(repoControl)

	useCaseStudent := useCaseStudent.StudentUseCase(serviceStudent, serviceControl)
	useCaseControl := useCaseControl.ControlUseCase(serviceControl)

	sController := studentController.ControllerStudent(useCaseStudent)
	cController := controlController.ControllerControl(useCaseControl)

	router := app.CreateRouter(sController, cController)
	router.StudentRouter(server)
	router.ControlRouter(server)

	err := server.Run(":8000")

	if err != nil {
		log.Fatal(err)
	}
}
