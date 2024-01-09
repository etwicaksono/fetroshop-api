package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Update(data *user_accesses.UserAccess, condition *user_accesses.UserAccess) *gorm.DB {
	return p.DB.Where(condition).Updates(data)
}
