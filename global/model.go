package global

import (
	"time"
)

type PvaModel struct {
	ID        uint      `gorm:"primary" json:"id"` // 主键ID
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdAt" structs:"-"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt" structs:"-"`
}
