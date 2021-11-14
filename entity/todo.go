package entity

import "time"

type Todo struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"<-:create"` //只創建不更新
	UpdatedAt   time.Time
	UserID      uint64 `gorm:"not null" json:"-"`
}
