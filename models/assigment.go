package models

import "time"

type Assigment struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	Devices      []Device   `json:"devices" gorm:"foreignkey:AssigmentID"`
	UserAssigned User       `gorm:"-" json:"user_assigned"`
	UserID       uint       `json:"user_id" binding:"required" gorm:"not null"`
	Description  string     `json:"description" binding:"required"`
	Ubication    string     `json:"ubication" binding:"required"`
	ProviderID   string     `json:"provider_id"`
	EndDate      string     `json:"end_date" binding:"required" gorm:"not null"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (a *Assigment) Create() error {
	if err := db.Create(a).Error; err != nil {
		return err
	}
	for i := 0; i < len(a.Devices); i++ {
		db.Model(&Device{}).Where("imei = ?", a.Devices[i].Imei).Update("assigment_id", a.ID)
	}
	return nil
}

func GetAssigments(provider *Provider) ([]Assigment, error) {
	assigments := []Assigment{}
	return assigments, db.Find(&assigments, "provider_id = ?", provider.ID).Unscoped().Error
}

func (a *Assigment) Get() error {
	devices := []Device{}
	db.Find(&devices, "assigment_id = ?", a.ID)
	if err := db.First(a, "id = ?", a.ID).Error; err != nil {
		return err
	}
	a.Devices = devices
	return nil
}

func (a *Assigment) Delete(device *Device) error {
	db.Model(device).Where("id = ?", device.ID).Update("ubication", "in stock")
	db.Model(device).Where("id = ?", device.ID).Update("is_assigned", false)
	return db.Delete(a, "id = ?", a.ID).Error
}
