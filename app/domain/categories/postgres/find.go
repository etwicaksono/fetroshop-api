package postgres

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Find(destination *categories.Category, condition *categories.Category) *gorm.DB {
	return p.DB.Where(condition).First(destination)
}