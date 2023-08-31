package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/control"
)

type Control interface {
	AddControl(ctx context.Context, data control.Create) (entity.Control, error)
}
