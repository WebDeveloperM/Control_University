package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/control"
	"controlUniversity/internal/service/student"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

type Repository struct {
	*bun.DB
}

func RepositoryControl(DB *bun.DB) Repository {
	return Repository{DB}
}

func (r Repository) AddControl(ctx context.Context, data control.Create) (entity.Control, error) {
	var addControl entity.Control
	var existStudent []entity.Control

	date := time.Now().Format("2006-01-02 15:04")
	addControl.StudentId = data.StudentId
	addControl.Status = data.Status
	addControl.Time = date
	addControl.UserId = data.UserId

	errNew := r.NewSelect().Model(&existStudent).Where("student_id = ?", data.StudentId).Scan(ctx)
	if errNew != nil {
		return entity.Control{}, errNew
	}
	if len(existStudent) > 0 {
		if existStudent[len(existStudent)-1].Status == data.Status && data.Status == "input" {
			return entity.Control{}, errors.New("Siz allaqachon ichkaridasiz")
		} else if existStudent[len(existStudent)-1].Status == data.Status && data.Status == "output" {
			return entity.Control{}, errors.New("Siz allaqachon chiqib bo`lgansiz")
		}
	}
	_, err := r.NewInsert().Model(&addControl).Exec(ctx)

	if err != nil {
		fmt.Println(err)
		return entity.Control{}, err
	}
	return addControl, nil

}

func (r Repository) GetAllControls(ctx context.Context, filter student.Filter) ([]entity.Control, int, error) {
	var controls []entity.Control
	query := r.NewSelect().Model(&controls).Column("*")

	if filter.StudentId != nil {
		today := time.Now().Format("2006-01-02")

		query.WhereGroup("AND", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("student_id = ?", *filter.StudentId).Where("time LIKE ?", "%"+today+"%")
		})
	}
	count, err := query.ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	for i := 0; i < len(controls); i++ {
		format := "2006-01-02 15:04"
		t, err := time.Parse(format, controls[i].Time)
		if err != nil {
			fmt.Println("Error parsing time:", err)
		}
		controls[i].Time = t.Format("01-02-2006 15:04")
	}

	return controls, count, nil
}
