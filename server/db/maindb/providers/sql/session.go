package sql

import (
	"context"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/db/maindb/models"
	"github.com/google/uuid"
)

func (p *provider) AddSession(ctx context.Context, session models.Session) error {
	if session.ID == "" {
		session.ID = uuid.New().String()
	}

	session.CreatedAt = time.Now().Unix()
	session.UpdatedAt = time.Now().Unix()

	err := p.db.Create(&session)
	if err != nil {
		return err.Error
	}
	return nil
}
