package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/control"
	"controlUniversity/internal/service/student"
	"errors"
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

	addControl.StudentId = data.StudentId
	addControl.Status = data.Status
	addControl.Time = time.Now()

	errNew := r.NewSelect().Model(&existStudent).Where("student_id = ?", data.StudentId).Scan(ctx)
	if errNew != nil {
		return entity.Control{}, errNew
	}
	if len(existStudent) > 0 {
		if existStudent[len(existStudent)-1].Status == data.Status && data.Status == 1 {
			return entity.Control{}, errors.New("Siz allaqachon ichkaridasiz")
		} else if existStudent[len(existStudent)-1].Status == data.Status && data.Status == 0 {
			return entity.Control{}, errors.New("Siz allaqachon chiqib bo`lgansiz")
		}
	}
	_, err := r.NewInsert().Model(&addControl).Exec(ctx)

	if err != nil {
		return entity.Control{}, errors.New("Bunday id bo`yicha bizda talaba mavjud emas")
	}
	return addControl, nil

}

func (r Repository) GetAllControls(ctx context.Context, filter student.Filter) ([]entity.Control, int, error) {
	var controls []entity.Control
	query := r.NewSelect().Model(&controls).Column("*")

	if filter.StudentId != nil {
		query.Where("student_id = ?", *filter.StudentId)
	}

	count, err := query.ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return controls, count, nil
}
