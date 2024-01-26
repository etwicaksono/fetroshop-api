package cities

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	ProvinceID int64          `gorm:"column:province_id"`
	Name       string         `gorm:"column:name"`
	CreatedAt  time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}
