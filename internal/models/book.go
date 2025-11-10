package models

import "time"

type Book struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	Name          string    `gorm:"type:varchar(100);not null;comment:书名" json:"name"`
	Author        string    `gorm:"type:varchar(100);not null;comment:作者" json:"author"`
	Category      string    `gorm:"type:varchar(50);comment:类别" json:"category"`
	PublishedDate time.Time `gorm:"type:date;comment:出版日期" json:"published_date"`
	Description   string    `gorm:"type:text;comment:内容简介" json:"description"`
	CreatedAt     time.Time `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt     time.Time `gorm:"comment:更新时间" json:"updated_at"`
	Removed       bool      `gorm:"default:false;comment:逻辑删除" json:"removed"`
}
