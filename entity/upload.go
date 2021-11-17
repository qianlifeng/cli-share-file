package entity

import "time"

type Upload struct {
	Id        string `gorm:"primaryKey"`
	Location  string `gorm:"size:1024"`
	FileName  string `gorm:"size:1024"`
	ExpiredAt time.Time
	CreatedAt time.Time
}
