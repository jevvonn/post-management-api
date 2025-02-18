package domain

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	ID                     uint      `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Title                  string    `gorm:"type:varchar(255);not null" json:"title"`
	ContentDescription     string    `gorm:"type:text;not null;column:content_description" json:"content"`
	Caption                string    `gorm:"type:text;not null" json:"caption"`
	Platforms              string    `gorm:"type:text;not null" json:"platforms"`
	DesignLink             string    `gorm:"type:text;not null" json:"design_link"`
	DeadlineBeforeRevision time.Time `gorm:"not null" json:"deadline_before_revision"`
	DeadlineAfterRevision  time.Time `json:"deadline_after_revision,omitempty"`
	UploadDate             time.Time `json:"upload_date,omitempty"`
	Status                 string    `gorm:"type:enum('DRAFT','PUBLISHED');not null;default:'DRAFT'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
