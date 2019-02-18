package models

import "time"

type Assigment struct {
	ID          uint `gorm:"primary_key" json:"id"`
	DeviceID    uint `json:"device_id" binding:"required" gorm:"not null"`
	UserID      uint `json:"user_id" binding:"required" gorm:"not null"`
	Description string
	Ubication   string
	EndDate     string     `json:"end_date" binding:"required" gorm:"not null"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (a *Assigment) Create(device *Device) error {
	db.Model(device).Where("id = ?", device.ID).Update("ubication", a.Ubication)
	db.Model(device).Where("id = ?", device.ID).Update("is_assigned", true)
	return db.Create(a).Error
}

func GetAssigments(device *Device) ([]Assigment, error) {
	assigments := []Assigment{}
	return assigments, db.Find(&assigments, "device_id = ?", device.ID).Unscoped().Error
}

func (a *Assigment) Get() error {
	return db.First(a, "id = ?", a.ID).Error
}

func (a *Assigment) Delete(device *Device) error {
	db.Model(device).Where("id = ?", device.ID).Update("ubication", "in stock")
	db.Model(device).Where("id = ?", device.ID).Update("is_assigned", false)
	return db.Delete(a, "id = ?", a.ID).Error
}
