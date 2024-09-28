package domain

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255);unique;not null" json:"email"`
	Phone     string    `gorm:"uniqueIndex;type:varchar(255);unique;not null" json:"phone"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
