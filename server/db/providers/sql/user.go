package sql

import (
	"context"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Adds a user into the database
func (p *provider) AddUser(ctx context.Context, user models.User) (models.User, error) {

	if user.Id == "" {
		user.Id = uuid.New().String()
	}

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	result := p.db.Clauses(
		clause.OnConflict{
			UpdateAll: true,
			Columns:   []clause.Column{{Name: "email"}},
		}).Create(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// Updates a user by id and before the update omits the null value in the model and selectively updates only the felids that are not null
func (p *provider) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	user.UpdatedAt = time.Now().Unix()

	result := p.db.Save(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// Deletes a user by id
func (p *provider) DeleteUser(ctx context.Context, user models.User) error {
	result := p.db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Gets a list of users
func (p *provider) ListUsers(ctx context.Context, pagination models.Pagination) ([]models.User, error) {
	var users []models.User

	result := p.db.Limit(int(pagination.Limit)).Offset(int(pagination.Offset)).Order("created_at DESC").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// Gets a single the user by email
func (p *provider) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	result := p.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// Gets a single user by id
func (p *provider) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User

	result := p.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// Updates a list of users by ids or if ids not given updates all the users
func (p *provider) UpdateUsers(ctx context.Context, data map[string]interface{}, ids []string) error {
	data["updated_at"] = time.Now().Unix()

	var res *gorm.DB
	if ids != nil && len(ids) > 0 {
		res = p.db.Model(&models.User{}).Where("id in ?", ids).Updates(data)
	} else {
		res = p.db.Model(&models.User{}).Updates(data)
	}

	if res.Error != nil {
		return res.Error
	}

	return nil
}
