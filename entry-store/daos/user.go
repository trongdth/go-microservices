package daos

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/trongdth/go_microservices/entry-store/models"
)

// User : struct
type User struct {
}

// NewUser :
func NewUser() *User {
	return &User{}
}

// Create : tx, user
func (u *User) Create(tx *gorm.DB, user *models.User) error {
	return errors.Wrap(tx.Create(user).Error, "tx.Create")
}

// Update : tx, user
func (u *User) Update(tx *gorm.DB, user *models.User) error {
	return errors.Wrap(tx.Save(user).Error, "tx.Save")
}

// FindByEmail : email
func (u *User) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "db.Where.First")
	}
	return &user, nil
}

// FindByID : id
func (u *User) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, errors.Wrap(err, "db.Where.First")
	}
	return &user, nil
}

// DeleteUser : userID
func (u *User) DeleteUser(tx *gorm.DB, userID uint) error {
	user := models.User{}
	err := tx.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return err
	}

	err = tx.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
