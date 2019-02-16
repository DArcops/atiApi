package models

import (
	"time"

	"github.com/darcops/atiApi/modules/encrypt"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"username"`
	Email     string     `gorm:"not null;unique" json:"email" binding:"required"`
	Password  string     `json:"pass,omitempty" binding:"required"`
	CanWrite  bool       `json:"administrator"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	//Permissions []*Permission  `gorm:"many2many:accesses;"`
}

type NewAdmin struct {
	Email string `json:"email" binding:"required"`
}

func GenerateToken(message []byte) ([]byte, error) {
	return encrypt.Encrypt(message)
}

func (u User) AddNewAdmin(email string) error {
	return db.Model(&User{}).Where("email = ?", email).Update("can_write", true).Error
}
