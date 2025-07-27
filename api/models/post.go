package models

import (
	"time"
)

type Post struct {
	ID        uint32         `gorm:"primary_key;auto_increment" json:"id"`
	Title     string         `gorm:"size:100;not null;" json:"title"`
	Content   string         `gorm:"type:text;not null;" json:"content"`
	AuthorID  uint32         `gorm:"not null;" json:"author_id"`
	CreatedAt time.Time 	 `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time 	 `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Author    User           `gorm:"foreignkey:AuthorID" json:"author"`

}