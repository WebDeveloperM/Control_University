package student

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/student"
	"errors"
	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func RepositoryStudent(DB *bun.DB) Repository {
	return Repository{DB}
}

func (r Repository) AddStudent(ctx context.Context, data student.Create) (entity.Student, error) {
	var newStudent entity.Student

	newStudent.FirstName = data.FirstName
	newStudent.LastName = data.LastName
	newStudent.UserId = data.UserId

	_, err := r.NewInsert().Model(&newStudent).Exec(ctx)
	if err != nil {
		return entity.Student{FirstName: "fN is null", LastName: "LN is null"}, err
	}
	return newStudent, nil

}

func (r Repository) GetDetailStudent(ctx context.Context, id int) (entity.Student, error) {
	var student entity.Student
	err := r.NewSelect().Model(&student).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return entity.Student{}, errors.New("Bunday talaba mavjud emas!!!")
	}
	return student, nil

}
