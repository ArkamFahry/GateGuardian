package sql

import (
	"context"

	"github.com/ArkamFahry/GateGuardian/server/db/models"
)

func (p *provider) AddUser(ctx context.Context, user models.User) (models.User, error) {

	return user, nil
}

func (p *provider) UpdateUser(ctx context.Context, user models.User) (models.User, error) {

	return user, nil
}

func (p *provider) DeleteUser(ctx context.Context, user models.User) error {

	return nil
}

func (p *provider) ListUsers(ctx context.Context) (models.User, error) {

	return models.User{}, nil
}

func (p *provider) GetUserByEmail(ctx context.Context, email string) (models.User, error) {

	return models.User{}, nil
}

func (p *provider) GetUserByID(ctx context.Context, id string) (models.User, error) {

	return models.User{}, nil
}

func (p *provider) UpdateUsers(ctx context.Context, data map[string]interface{}, ids []string) error {

	return nil
}
