package model

import (
	"time"
)

// User 用户账号模型
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:50;uniqueIndex" json:"username"` // 登录用户名
	Password  string    `gorm:"size:100" json:"-"`                   // 加密后的密码（不返回给前端）
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
