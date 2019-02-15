package models

import "time"

type Assigment struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	DeviceID  uint       `json:"device_id" binding:"required" gorm:"not null"`
	UserID    uint       `json:"user_id" binding:"required" gorm:"not null"`
	EndDate   string     `json:"end_date" binding:"required" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (a *Assigment) Create() error {
	return db.Create(a).Error
}

func (a *Assigment) Get() error {
	return db.First(a, "id = ?", a.ID).Error
}

func (a *Assigment) Delete() error {
	return db.Delete(a, "id = ?", a.ID).Error
}
