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
	newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Silent,
			IgnoreRecordNotFoundError: true,
		},
	)
)

func CreateDB() error {

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Table(structs.Role1+"s").AutoMigrate(&structs.User{})
	db.Table(structs.Role2+"s").AutoMigrate(&structs.Organizer{})

	return nil
}

func AddPart(user structs.User) error {

	hashPwd, err := generateHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPwd

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		return err
	}

	return db.Table(structs.Role1+"s").Create(&user).Error
}

func AddOrg(org structs.Organizer) error {

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		return err
	}

	return db.Table(structs.Role2+"s").Create(&org).Error
}

func IsValidUser(user structs.User) error {

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: newLogger})
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

	db, err := gorm.Open(sqlite.Open("storage/storage.db"), &gorm.Config{Logger: newLogger})
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
