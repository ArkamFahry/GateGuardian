package genjidb

import (
	"context"
	"fmt"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/google/uuid"
)

func (p *provider) AddUser(ctx context.Context, user models.User) (models.User, error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	if user.Roles == "" {
		defaultRole, _ := env.GetEnvByKey(constants.DefaultRoles)
		user.Roles = defaultRole
	}

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	query := fmt.Sprintf(`INSERT INTO %s VALUES ?`, models.Collections.User)
	err := p.db.Exec(query, &user)

	return user, err
}

func (p *provider) UpdateUser() {

}

func (p *provider) DeleteUser() {

}

func (p *provider) ListUser() {

}

func (p *provider) GetUserByEmail() {

}

func (p *provider) GetUserByID() {

}

func (p *provider) UpdateUsers() {

}
