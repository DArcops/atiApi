package models

import "time"

type Provider struct {
	ID        uint `gorm:"primary_key" json:"id"`
	Name      string
	Email     string
	Phone     string
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (p *Provider) Create() error {
	return db.Create(p).Error
}

func (p *Provider) Get() error {
	return db.First(p, "id = ?", p.ID).Error
}

func GetProviders() ([]Provider, error) {
	providers := []Provider{}
	return providers, db.Find(&providers).Error
}

func (p *Provider) GetDevices(from, to int64) ([]Device, error) {
	devices := []Device{}

	if err := db.Find(&devices, "provider_id = ?", p.ID).Error; err != nil {
		return nil, err
	}

	return devices[from:to], nil
}

//refact this function
func (p *Provider) AddDevices(devices []*Device) error {
	for _, device := range devices {
		if err := device.Create(); err != nil {
			return err
		}
	}
	return nil
}
