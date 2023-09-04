package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/control"
	"controlUniversity/internal/service/student"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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
			return q.Where("student_id = ?", *filter.StudentId).Where("time LIKE ?", "%"+today+"%").Where("is_delete = ?", false)
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

func (r Repository) DeleteControl(ctx *gin.Context, id int) (bool, error) {
	var contorl entity.Control
	var deleteControl entity.DeleteControl
	errFind := r.NewSelect().Model(&contorl).Where("id = ?", id).Scan(ctx)
	if errFind != nil {
		return false, errFind
	}

	query := r.NewUpdate().Model(&contorl)
	if contorl.IsDelete {
		return false, errors.New("This object already been deleted")
	}
	contorl.IsDelete = true

	//Delete-Controls
	userId, errID := ctx.Get("userID")
	if !errID {
		return false, errors.New("Deleting an object is only for Superuser or Admin")
	}
	deleteTime := time.Now().Format("01-02-2006 15:04")
	deleteControl.ContorlId = id
	deleteControl.UserId = userId.(int64)
	deleteControl.DeleteAt = deleteTime

	_, errDel := r.NewInsert().Model(&deleteControl).Exec(ctx)
	if errDel != nil {
		return false, errDel
	}

	_, err := query.Where("id = ?", id).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r Repository) UpadateControl(ctx *gin.Context, data control.Update, id int) (entity.Control, error) {
	var control entity.Control
	var updateControl entity.UpdateControl
	errFind := r.NewSelect().Model(&control).Where("id = ?", id).Scan(ctx)
	if errFind != nil {
		return entity.Control{}, errors.New("Id not found")
	}
	if control.IsDelete {
		return entity.Control{}, errors.New("This object already been deleted")
	}
	query := r.NewUpdate().Model(&control)

	if data.Status != "" {
		control.Status = data.Status
	}
	if data.Time != "" {
		control.Time = data.Time
	}

	//Update-Controls
	userId, errID := ctx.Get("userID")
	if !errID {
		return entity.Control{}, errors.New("Updating an object is only for Superuser or Admin")
	}

	chanded := fmt.Sprintf("Status: \"%v\", Time: \"%v\"", data.Status, data.Time)

	updateTime := time.Now().Format("01-02-2006 15:04")
	updateControl.ContorlId = id
	updateControl.UserId = userId.(int64)
	updateControl.UpdateAt = updateTime
	updateControl.Changes = chanded

	_, errDel := r.NewInsert().Model(&updateControl).Exec(ctx)
	if errDel != nil {
		return entity.Control{}, errDel
	}

	_, err := query.Where("id = ?", id).Exec(ctx)
	if err != nil {
		return entity.Control{}, err
	}
	return control, nil
}
