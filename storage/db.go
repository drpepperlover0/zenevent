package storage

import (
	"log"
	"os"

	"github.com/drpepperlover0/internal/structs"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	errLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
		},
	)
)

func CreateDB() error {

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Table(structs.Role1 + "s").AutoMigrate(&structs.User{})
	db.Table(structs.Role2 + "s").AutoMigrate(&structs.Organizer{})

	db.AutoMigrate(&structs.Event{})
	db.AutoMigrate(&structs.EventMember{})

	return nil
}

func AddEvent(event structs.Event) error {

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: errLogger})
	if err != nil {
		return err
	}

	return db.Table("events").Create(&event).Error
}

func FindEvents(event string) ([]structs.Event, error) {

	foundEvents := []structs.Event{}

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: errLogger})
	if err != nil {
		return foundEvents, err
	}

	if err := db.Table("events").Where("event_theme = ?", event).Find(&foundEvents).Error; err != nil {
		return foundEvents, err
	}

	return foundEvents, nil
}

func AddPart(user structs.User) error {

	hashPwd, err := generateHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPwd

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: errLogger})
	if err != nil {
		return err
	}

	return db.Table(structs.Role1 + "s").Create(&user).Error
}

func AddOrg(org structs.Organizer) error {

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: errLogger})
	if err != nil {
		return err
	}

	return db.Table(structs.Role2 + "s").Create(&org).Error
}

func AddToEvent(name, eventId string) error {

	var user structs.User

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: errLogger})
	if err != nil {
		return err
	}

	if err := db.Table(structs.Role1+"s").Find(&user, "username = ?", name).Error; err != nil {
		return err
	}

	return db.Table("event_members").Create(&structs.EventMember{
		EventID: eventId,
		UserID:  user.UserID,
	}).Error
}

func IsValidUser(user structs.User) error {

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: errLogger})
	if err != nil {
		return err
	}
	pw := user.Password

	db.Table(structs.Role1+"s").First(&user, "username = ?", user.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw)); err != nil {
		return err
	}

	return nil
}

func IsValidOrg(org structs.Organizer) error {

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: errLogger})
	if err != nil {
		return err
	}

	return db.Table(structs.Role2+"s").First(&org, "name = ? AND org_id = ?", org.Name, org.SID).Error
}

func generateHash(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
