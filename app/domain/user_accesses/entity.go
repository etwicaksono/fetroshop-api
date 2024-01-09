package user_accesses

import (
	"time"

	"gorm.io/gorm"
)

type UserAccess struct {
	ID        string         `gorm:"column:id;" json:"id"`
	UserID    int64          `gorm:"column:user_id" json:"userId"`
	Platform  string         `gorm:"column:platform" json:"platform"`
	UserAgent string         `gorm:"column:user_agent" json:"userAgent"`
	ExpiredAt time.Time      `gorm:"column:expired_at;" json:"expiredAt"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}
