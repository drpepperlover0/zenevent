package structs

import "gorm.io/gorm"

type User struct {
	UserID   uint   `gorm:"column:user_id;primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"text"`
}

type Organizer struct {
	gorm.Model
	IndividEmail string `gorm:"column:individual_email;unique"`
	Name         string `gorm:"unique"`
	SID          string `gorm:"column:org_id;not null;unique"`
}

type Event struct {
	EventID uint `gorm:"column:event_id;primaryKey;autoIncrement"`

	OrgName     string `gorm:"column:org_name"`
	EventName   string `gorm:"column:event_name"`
	EventTheme  string `gorm:"column:event_theme"`
	Description string `gorm:"column:event_desc"`
	EventDate   string `gorm:"column:event_date"`
}

type EventMember struct {
	EventID string `gorm:"column:event_id"`
	UserID  uint   `gorm:"column:user_id;not null"`
}

const (
	Role1 = "participant"
	Role2 = "organizer"
)
