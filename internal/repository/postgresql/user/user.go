package user

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/user"
	token2 "controlUniversity/internal/utils/token"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	*bun.DB
}

func RepositoryUser(DB *bun.DB) Repository {
	return Repository{DB}
}

func (r Repository) CreateUser(ctx context.Context, data user.Create) (string, error) {
	var user entity.User
	hashed, _ := hashPassword(data.Password)
	user.Username = data.Username
	user.Password = hashed
	user.Role = data.Role

	exist, _ := r.NewSelect().Model((*entity.User)(nil)).Where("username = ?", data.Username).Exists(ctx)
	if exist {
		return "", errors.New(data.Username + " - allaqachon ro`yhatdan o`tgansiz")
	}

	_, err := r.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		return "", err
	}
	token, errToken := token2.Generate(user.Id, user.Role)
	if errToken != nil {
		return "", errToken
	}
	return token, nil
}

func (r Repository) GetOneUser(ctx context.Context, id int) (entity.User, error) {
	var oneUser entity.User
	err := r.NewSelect().Model(&oneUser).Where("id = ?", id).Scan(ctx)
	fmt.Println(err, oneUser, "000000000000000000000000000000")
	if err != nil {
		return entity.User{}, err
	}
	return oneUser, err
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}
