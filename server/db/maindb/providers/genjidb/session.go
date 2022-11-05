package genjidb

import (
	"fmt"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/db/maindb/models"
	"github.com/google/uuid"
)

func (p *provider) AddSession(session models.Session) error {
	if session.ID == "" {
		session.ID = uuid.New().String()
	}

	session.CreatedAt = time.Now().Unix()
	session.UpdatedAt = time.Now().Unix()

	insertSessionQuery := fmt.Sprintf("INSERT INTO %s (id, user_id, user_agent, ip, created_at, updated_at) VALUES ('%s', '%s', '%s', '%s', %d, %d)", models.Collections.Session, session.ID, session.UserID, session.UserAgent, session.IP, session.CreatedAt, session.UpdatedAt)
	err := p.db.Exec(insertSessionQuery)
	if err != nil {
		return err
	}
	return nil
}
