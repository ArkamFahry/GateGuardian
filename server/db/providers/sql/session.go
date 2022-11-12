package sql

import (
	"context"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

// Adds session to the database
func (p *provider) AddSession(ctx context.Context, session models.Session) error {
	if session.Id == "" {
		session.Id = uuid.New().String()
	}

	session.CreatedAt = time.Now().Unix()
	session.UpdatedAt = time.Now().Unix()

	res := p.db.Clauses(
		clause.OnConflict{
			DoNothing: true,
		}).Create(&session)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
