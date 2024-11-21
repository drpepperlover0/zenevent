package structs

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"text"`
	Role     string `gorm:"text"`
}

type Organizer struct {
	gorm.Model
	IndividEmail string `gorm:"column:individual_email;unique"`
	Name         string `gorm:"unique"`
	SID          string `gorm:"column:org_id;not null;unique"`
}

const (
	Role1 = "participant"
	Role2 = "organizer"
)
