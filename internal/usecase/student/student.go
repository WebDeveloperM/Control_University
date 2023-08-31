package student

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/student"
)

type UseCase struct {
	student  Student
	controls Controls
}

func StudentUseCase(student Student, controls Controls) UseCase {
	return UseCase{student, controls}
}

func (c UseCase) AddStudent(ctx context.Context, data student.Create) (entity.Student, error) {
	return c.student.AddStudent(ctx, data)
}

func (c UseCase) GetDetail(ctx context.Context, id int) (student.Detail, error) {
	var detail student.Detail
	studentDetail, err := c.student.GetDetailStudent(ctx, id)
	if err != nil {
		return student.Detail{}, err
	}
	detail.Id = studentDetail.Id
	detail.FirstName = studentDetail.FirstName
	detail.Lastname = studentDetail.LastName

	controls, _, errControls := c.controls.GetAll(ctx, student.Filter{StudentId: &id})
	if errControls != nil {
		return student.Detail{}, errControls
	}
	detail.Controls = controls
	return detail, nil
}
