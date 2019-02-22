package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strings"
	"time"
)

type Device struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	ProviderID    uint       `json:"provider_id" binding:"required" gorm:"not null"`
	Imei          string     `json:"imei" binding:"required" gorm:"not null;unique"`
	Mpn           string     `json:"mpn,omitempty"`
	Name          string     `json:"name,omitempty"`
	IsAssigned    bool       `json:"is_assigned,omitempty"`
	AssigmentID   uint       `json:"assigment_id,omitempty"`
	AdmissionDate string     `json:"admission_date,omitempty"`
	Ubication     string     `json:"ubication,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (d *Device) Create() error {
	d.IsAssigned = false
	d.Ubication = "in stock"
	return db.Create(d).Error
}

func (d *Device) Get() error {
	return db.First(d, "id = ?", d.ID).Error
}

func (d *Device) Delete() error {
	return db.Delete(d, "id = ?", d.ID).Error
}

func SaveDevicesFromFile(file *multipart.FileHeader, p *Provider) error {
	bucket := map[string]int{}
	tx := db.Begin()

	reader, err := file.Open()
	if err != nil {
		return err
	}

	rd := csv.NewReader(reader)

	countrows := 0
	fmt.Println("Provider id", p.ID)

	for {
		row, err := rd.Read()
		if err == io.EOF {
			break
		}

		fmt.Println("ROWW", row)

		if countrows == 0 {
			for i := 0; i < len(row); i++ {
				bucket[strings.ToLower(row[i])] = i
			}
			countrows++
			continue
		}

		if row[bucket["imei"]] == "" {
			continue
		}

		if err := tx.Create(&Device{
			ProviderID:    p.ID,
			Imei:          row[bucket["imei"]],
			Mpn:           row[bucket["modelo"]],
			Name:          row[bucket["nombre"]],
			IsAssigned:    false,
			AdmissionDate: row[bucket["fecha de ingreso"]],
			Ubication:     row[bucket["ubicacion"]],
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
